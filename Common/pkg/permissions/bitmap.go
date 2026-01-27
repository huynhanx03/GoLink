package permissions

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/bits"
	"sync"

	"github.com/RoaringBitmap/roaring"
)

const (
	ScopeShift = 4 // 16 scopes per resource
	ScopeMask  = (1 << ScopeShift) - 1
)

var (
	bitmapPool = sync.Pool{
		New: func() any {
			return roaring.New()
		},
	}
)

func GetBitmap() *roaring.Bitmap {
	return bitmapPool.Get().(*roaring.Bitmap)
}

func PutBitmap(rb *roaring.Bitmap) {
	rb.Clear()
	bitmapPool.Put(rb)
}

func Flatten(resourceID int, scopeMask int) []uint32 {
	var ids []uint32
	for i := 0; i < (1 << ScopeShift); i++ {
		if (scopeMask & (1 << i)) != 0 {
			flatID := uint32((resourceID << ScopeShift) | i)
			ids = append(ids, flatID)
		}
	}
	return ids
}

func Compress(resourceScopes map[int]int) (string, error) {
	rb := GetBitmap()
	defer PutBitmap(rb)

	for resID, scope := range resourceScopes {
		flatIDs := Flatten(resID, scope)
		for _, id := range flatIDs {
			rb.Add(id)
		}
	}
	rb.RunOptimize()

	var buf bytes.Buffer
	if _, err := rb.WriteTo(&buf); err != nil {
		return "", fmt.Errorf("failed to write bitmap: %w", err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func Decompress(b64 string) (*roaring.Bitmap, error) {
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("invalid base64: %w", err)
	}

	rb := GetBitmap()
	if _, err := rb.ReadFrom(bytes.NewReader(data)); err != nil {
		PutBitmap(rb)
		return nil, fmt.Errorf("failed to read bitmap: %w", err)
	}
	return rb, nil
}

// CheckPermission checks if the bitmap contains ALL permissions specified in the scopeMask.
func CheckPermission(rb *roaring.Bitmap, resourceID int, scopeMask int) bool {
	tempMask := scopeMask
	for tempMask != 0 {
		zeros := bits.TrailingZeros(uint(tempMask))

		flatID := uint32((resourceID << ScopeShift) | zeros)
		if !rb.Contains(flatID) {
			return false
		}

		tempMask &= (tempMask - 1)
	}

	return true
}

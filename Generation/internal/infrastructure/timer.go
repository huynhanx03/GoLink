package infrastructure

import (
	t "go-link/common/pkg/timer"
	"go-link/generation/global"
	"time"
)

func SetupTimer() {
	global.Time10ms = t.NewCachedTimer(10 * time.Millisecond)
	global.Time1s = t.NewCachedTimer(1 * time.Second)
}

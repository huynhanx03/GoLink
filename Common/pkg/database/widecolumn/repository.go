package widecolumn

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/gocql/gocql"

	"go-link/common/pkg/database"
	"go-link/common/pkg/dto"
)

// Repository defines the interface for Wide Column DB repositories
type Repository[T any] database.Repository[T, any]

// BaseRepository provides common database operations using generics
type BaseRepository[T Model] struct {
	session   *gocql.Session
	tableName string
}

// NewBaseRepository creates a new generic repository
func NewBaseRepository[T Model](session *gocql.Session, dummy T) *BaseRepository[T] {
	return &BaseRepository[T]{
		session:   session,
		tableName: dummy.TableName(),
	}
}

// Create inserts a new model
func (r *BaseRepository[T]) Create(ctx context.Context, model *T) error {
	val := any(model).(T)

	cols := val.ColumnNames()
	placeholders := make([]string, len(cols))
	for i := range cols {
		placeholders[i] = "?"
	}

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		r.tableName,
		strings.Join(cols, ", "),
		strings.Join(placeholders, ", "),
	)

	return r.session.Query(stmt, val.ColumnValues()...).WithContext(ctx).Exec()
}

// Update updates a model
func (r *BaseRepository[T]) Update(ctx context.Context, model *T) error {
	return r.Create(ctx, model)
}

// Delete removes a model by ID
func (r *BaseRepository[T]) Delete(ctx context.Context, id any) error {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.tableName)
	return r.session.Query(stmt, id).WithContext(ctx).Exec()
}

// Get retrieves a model by ID
func (r *BaseRepository[T]) Get(ctx context.Context, id any) (*T, error) {
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", r.tableName)

	iter := r.session.Query(stmt, id).WithContext(ctx).Iter()
	row := make(map[string]any)
	if !iter.MapScan(row) {
		return nil, ErrNotFound
	}

	var result T
	if err := r.scanToStruct(row, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Find retrieves models with pagination
func (r *BaseRepository[T]) Find(ctx context.Context, opts *dto.QueryOptions) (*dto.Paginated[*T], error) {
	stmt := fmt.Sprintf("SELECT * FROM %s", r.tableName)
	// TODO: Add WHERE clause builder based on opts.Filter

	iter := r.session.Query(stmt).WithContext(ctx).Iter()

	var records []*T
	for {
		row := make(map[string]any)
		if !iter.MapScan(row) {
			break
		}

		var model T
		if err := r.scanToStruct(row, &model); err != nil {
			continue
		}
		records = append(records, &model)
	}

	return &dto.Paginated[*T]{
		Records: &records,
		// Pagination metadata is hard to calculate without Count()
	}, nil
}

// Exists checks if a model exists
func (r *BaseRepository[T]) Exists(ctx context.Context, id any) (bool, error) {
	stmt := fmt.Sprintf("SELECT count(*) FROM %s WHERE id = ?", r.tableName)
	var count int
	if err := r.session.Query(stmt, id).WithContext(ctx).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// BatchCreate inserts multiple models
func (r *BaseRepository[T]) BatchCreate(ctx context.Context, models []*T) error {
	if len(models) == 0 {
		return nil
	}

	batch := r.session.NewBatch(gocql.LoggedBatch)
	for _, model := range models {
		val := any(model).(T)
		cols := val.ColumnNames()
		placeholders := make([]string, len(cols))
		for i := range cols {
			placeholders[i] = "?"
		}

		stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
			r.tableName,
			strings.Join(cols, ", "),
			strings.Join(placeholders, ", "),
		)
		batch.Query(stmt, val.ColumnValues()...)
	}

	return r.session.ExecuteBatch(batch)
}

// BatchDelete removes multiple models by IDs
func (r *BaseRepository[T]) BatchDelete(ctx context.Context, ids []any) error {
	if len(ids) == 0 {
		return nil
	}

	batch := r.session.NewBatch(gocql.LoggedBatch)
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.tableName)
	for _, id := range ids {
		batch.Query(stmt, id)
	}

	return r.session.ExecuteBatch(batch)
}

func (r *BaseRepository[T]) scanToStruct(row map[string]any, target any) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("target must point to a struct")
	}

	typ := elem.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldName := strings.ToLower(field.Name)

		tag := field.Tag.Get("json")
		if tag != "" {
			parts := strings.Split(tag, ",")
			if parts[0] != "" {
				fieldName = parts[0]
			}
		}

		if val, ok := row[fieldName]; ok {
			f := elem.Field(i)
			if f.CanSet() {
				v := reflect.ValueOf(val)
				if v.IsValid() && v.Type().ConvertibleTo(f.Type()) {
					f.Set(v.Convert(f.Type()))
				} else if v.IsValid() && f.Kind() == reflect.Ptr && v.Type().ConvertibleTo(f.Type().Elem()) {
					ptrVal := reflect.New(f.Type().Elem())
					ptrVal.Elem().Set(v.Convert(f.Type().Elem()))
					f.Set(ptrVal)
				}
			}
		}
	}
	return nil
}

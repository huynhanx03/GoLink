package ent

import (
	"fmt"
	"reflect"
	"time"

	"go-link/common/pkg/constraints"
)

type fieldInfo struct {
	Index      int
	Name       string
	SetterName string
	IsTime     bool
}

type entityMetadata struct {
	Typ            reflect.Type
	EntityName     string
	ClientIndex    int
	MethodCreate   int
	MethodUpdate   int
	MethodDelete   int
	MethodGet      int
	MethodQuery    int
	MethodMapBulk  int
	MethodTx       int
	MethodCommit   int
	MethodRollback int
	Fields         []fieldInfo
}

func newEntityMetadata[T any, ID constraints.ID](client any) (*entityMetadata, error) {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("T must be a struct or pointer to struct, got %s", typ.Kind())
	}

	entityName := typ.Name()
	meta := &entityMetadata{
		Typ:        typ,
		EntityName: entityName,
	}

	clientVal := reflect.ValueOf(client).Elem()
	clientTyp := clientVal.Type()
	clientField, ok := clientTyp.FieldByName(entityName)
	if !ok {
		return nil, fmt.Errorf("entity client field not found for %s", entityName)
	}
	meta.ClientIndex = clientField.Index[0]

	specificClientVal := clientVal.Field(meta.ClientIndex)
	specificClientTyp := specificClientVal.Type()

	findMethodIndex := func(typ reflect.Type, methodName string) (int, error) {
		m, ok := typ.MethodByName(methodName)
		if !ok {
			return -1, fmt.Errorf("method %s not found for type %s", methodName, typ.Name())
		}
		return m.Index, nil
	}

	var err error
	if meta.MethodCreate, err = findMethodIndex(specificClientTyp, MethodCreate); err != nil {
		return nil, err
	}
	if meta.MethodUpdate, err = findMethodIndex(specificClientTyp, MethodUpdateOneID); err != nil {
		return nil, err
	}
	if meta.MethodDelete, err = findMethodIndex(specificClientTyp, MethodDeleteOneID); err != nil {
		return nil, err
	}
	if meta.MethodGet, err = findMethodIndex(specificClientTyp, MethodGet); err != nil {
		return nil, err
	}
	if meta.MethodQuery, err = findMethodIndex(specificClientTyp, MethodQuery); err != nil {
		return nil, err
	}

	if m, ok := specificClientTyp.MethodByName(MethodMapCreateBulk); ok {
		meta.MethodMapBulk = m.Index
	} else {
		meta.MethodMapBulk = -1
	}

	// Cache Tx methods
	if meta.MethodTx, err = findMethodIndex(clientTyp, MethodTx); err != nil {
		meta.MethodTx = -1
	}

	if meta.MethodTx != -1 {
		txMethod, _ := clientTyp.MethodByName(MethodTx)
		txType := txMethod.Type.Out(0) // *Tx
		if txType.Kind() == reflect.Ptr {
			if meta.MethodCommit, err = findMethodIndex(txType, MethodCommit); err != nil {
				return nil, err
			}
			if meta.MethodRollback, err = findMethodIndex(txType, MethodRollback); err != nil {
				return nil, err
			}
		}
	} else {
		meta.MethodCommit = -1
		meta.MethodRollback = -1
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if field.PkgPath != "" ||
			field.Name == FieldID ||
			field.Name == FieldEdges ||
			field.Name == FieldConfig {
			continue
		}

		fInfo := fieldInfo{
			Index:      i,
			Name:       field.Name,
			SetterName: PrefixSet + field.Name,
			IsTime:     field.Type == reflect.TypeOf(time.Time{}),
		}
		meta.Fields = append(meta.Fields, fInfo)
	}

	return meta, nil
}

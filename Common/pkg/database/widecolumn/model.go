package widecolumn

type Model interface {
	TableName() string
	ColumnNames() []string
	ColumnValues() []interface{}
}

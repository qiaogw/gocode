package model

const indexPri = "PRIMARY"

var ModelApp Model

type Model interface {
	Init()
	GetDB() (data []Db, err error)
	GetTables(db string) ([]Table, error)
	GetColumn(db, table string) (*ColumnData, error)
}

func NewAutoCode(dt string) {
	switch dt {
	case "mysql":
		ModelApp = ModelMysqlApp
	case "postgres":
		ModelApp = ModelPostgresApp
	}
}

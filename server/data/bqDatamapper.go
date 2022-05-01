package data

type BqDatamapper[T any] struct {
	table string
	db    *BqDB
}

func NewDataMapper[T any](db *BqDB, tablename string) *BqDatamapper[T] {
	dm := new(BqDatamapper[T])
	dm.db = db
	dm.table = tablename

	return dm
}

func (dm *BqDatamapper[T]) FindAll() ([]T, error) {
	query := "SELECT * FROM " + dm.table
	return Query[T](dm.db, query)
}

package data

import (
	"log"
	"reflect"
)

type BqDatamapper[T any] struct {
	table   string
	db      *BqDB
	cFilter string // cFilter is the column filter, basically (for now) the list of columns in a sql table
}

func NewDataMapper[T any](db *BqDB, tablename string) *BqDatamapper[T] {
	dm := new(BqDatamapper[T])
	dm.db = db
	dm.table = tablename

	fieldNames := reflector(*new(T))
	n := len(fieldNames) - 1
	for i := 0; i < n; i++ {
		dm.cFilter = dm.cFilter + fieldNames[i] + ","
	}

	dm.cFilter = dm.cFilter + fieldNames[n] // no trailing comma
	log.Println("discovered the follwing fields:", dm.cFilter)

	return dm
}

func (dm *BqDatamapper[T]) FindAll() ([]T, error) {
	query := "SELECT * FROM " + dm.table
	return Query[T](dm.db, query)
}

func reflector(strct interface{}) []string {
	var items []interface{}
	var fieldNames []string
	items = append(items, strct) // init the analysis FIFO

	j := 0
	for j < len(items) {
		t := reflect.TypeOf(items[j])
		log.Println("(A) reflecting", t.Name(), "of type", t.Kind()) //, "with # fields", t.NumField())

		for i := 0; i < t.NumField(); i++ {
			log.Println("(A) field", t.Field(i).Name, "has type", t.Field(i).Type)

			if t.Field(i).Type.Kind() == reflect.Struct {
				// composite type, append for analysis in the FIFO
				log.Println("(+) appending field", t.Field(i).Name, "for analysis")
				items = append(items, reflect.ValueOf(items[j]).Field(i).Interface())
			} else {
				if t.Field(i).IsExported() {
					// store field name
					fieldNames = append(fieldNames, t.Field(i).Name)
					log.Println("(F) storing field", t.Field(i).Name)
				} else {
					log.Println("(P) skipping unexported field", t.Field(i).Name)
				}
			}
		}

		items = items[1:] // pop analysed field
	}

	return fieldNames
}

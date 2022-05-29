package data

import (
	"dsdr/models"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type BqDatamapper[T any] struct {
	table      string
	db         *BqDB
	fieldNames []string
	cFilter    string // cFilter is the column filter, basically (for now) the list of columns in a sql table
}

type BqQuery string

func NewDataMapper[T any](db *BqDB, tablename string) *BqDatamapper[T] {
	dm := new(BqDatamapper[T])
	dm.db = db
	dm.table = tablename

	dm.fieldNames = reflector(*new(T))
	for i := 0; i < len(fieldNames); i++ {
		dm.fieldNames[i] = strings.ToLower(dm.fieldNames[i])
	}

	n := len(dm.fieldNames) - 1
	for i := 0; i < n; i++ {
		dm.cFilter = dm.cFilter + dm.fieldNames[i] + ","
	}

	dm.cFilter = dm.cFilter + dm.fieldNames[n] // no trailing comma

	return dm
}

func (dm *BqDatamapper[T]) FindAll() BqQuery {
	query := BqQuery("SELECT " + dm.cFilter + " FROM " + dm.table)
	log.Println(query)
	return query
}

func (dm *BqDatamapper[T]) Run(query BqQuery) ([]T, error) {
	return Query[T](dm.db, query)
}

func (bq BqQuery) Where(filter models.BqIAMRole) BqQuery {
	filtersMap := bq.reflectFilter(filter)
	log.Println(filtersMap)

	return bq
}

func (bq *BqQuery) reflectFilter(filter models.BqIAMRole) map[string]string {
	var filterFields = make(map[string]string)

	t := reflect.TypeOf(filter)
	v := reflect.ValueOf(filter).FieldByName() //<-- ciò è figo!!!
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Int {
			filterFields[t.Field(i).Name] = strconv.Itoa(v.Field(i).Interface().(int))

		} else if t.Field(i).Type.Kind() == reflect.String {
			filterFields[t.Field(i).Name] = v.Field(i).Interface().(string)

		} else {
			log.Println("filtering for type", t.Field(i).Type.Kind(), "still missing sorry!")
		}
	}

	return filterFields
}

func reflector(strct interface{}) []string {
	var items []interface{}
	var fieldNames []string
	items = append(items, strct) // init the analysis FIFO

	j := 0
	for j < len(items) {
		t := reflect.TypeOf(items[j])
		log.Println("(A) reflecting", t.Name(), "of type", t.Kind())

		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Type.Kind() == reflect.Struct {
				// composite type, append for analysis in the FIFO
				log.Println("(+) appending field", t.Field(i).Name, "for analysis")
				items = append(items, reflect.ValueOf(items[j]).Field(i).Interface())
			} else {
				if t.Field(i).IsExported() {
					// store field name
					fieldNames = append(fieldNames, t.Field(i).Name)
					log.Println("(F) storing field", t.Field(i).Name, "of type", t.Field(i).Type.Kind())
				} else {
					log.Println("(P) skipping unexported field", t.Field(i).Name)
				}
			}
		}

		items = items[1:] // pop analysed field
	}

	return fieldNames
}

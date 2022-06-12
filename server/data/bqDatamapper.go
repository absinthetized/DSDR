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

type BqQuery struct {
	query      string
	fieldNames []string
}

func NewDataMapper[T any](db *BqDB, tablename string) *BqDatamapper[T] {
	dm := new(BqDatamapper[T])
	dm.db = db
	dm.table = tablename

	dm.fieldNames = reflector(*new(T))

	n := len(dm.fieldNames) - 1
	for i := 0; i < n; i++ {
		dm.cFilter = dm.cFilter + strings.ToLower(dm.fieldNames[i]) + ","
	}

	dm.cFilter = dm.cFilter + dm.fieldNames[n] // no trailing comma

	return dm
}

func (dm *BqDatamapper[T]) FindAll() BqQuery {
	query := BqQuery{"SELECT " + dm.cFilter + " FROM " + dm.table, dm.fieldNames}
	log.Println(query)
	return query
}

func (dm *BqDatamapper[T]) Run(query BqQuery) ([]T, error) {
	return Query[T](dm.db, query)
}

func (bq BqQuery) Where(filter models.BqIAMRoleFilter) BqQuery {
	filtersMap := bq.reflectFilter(filter)
	log.Println(filtersMap)

	return bq
}

func (bq *BqQuery) reflectFilter(filter models.BqIAMRoleFilter) map[string]string {
	var filterFields = make(map[string]string)

	for _, fieldName := range bq.fieldNames {
		v := reflect.ValueOf(&filter).Elem().FieldByName(fieldName)

		if v.Kind() == reflect.Pointer {
			e := v.Elem()
			if !e.IsValid() {
				continue

			} else if e.Kind() == reflect.Int {
				filterFields[fieldName] = strconv.Itoa(e.Interface().(int))

			} else if e.Kind() == reflect.String {
				filterFields[fieldName] = e.Interface().(string)

			} else {
				log.Println("unable to dereference pointer of type", e.Kind(), "sorry!")
			}

		} else {
			log.Println("filtering for type", v.Kind(), "still missing sorry!")
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

package migration

import (
	"reflect"
	"sort"
)

type Migration struct {
	Name   string
	UpFn   func() error
	DownFn func() error
}

var Migrations_Arr []Migration

func SortMigArr() {

	sort.Slice(Migrations_Arr, func(i, j int) bool {

		v1 := reflect.ValueOf(Migrations_Arr[i]).FieldByName("Name") // sort according to name
		v2 := reflect.ValueOf(Migrations_Arr[j]).FieldByName("Name")
		return v1.String() < v2.String()
	})

}

func init() {

}

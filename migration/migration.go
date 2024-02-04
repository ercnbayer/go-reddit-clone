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

func sortMigArr() {

	sort.Slice(Migrations_Arr, func(i, j int) bool {
		//sort according to name
		v1 := reflect.ValueOf(Migrations_Arr[i]).FieldByName("Name")
		v2 := reflect.ValueOf(Migrations_Arr[j]).FieldByName("Name")
		return v1.String() < v2.String()
	})

}

func init() {
	//for sorting mig's alphabetically
	sortMigArr()

}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var d *Data = &Data{f: "yiji"}

	var y = reflect.ValueOf(d)
	f := y.Elem().FieldByName("f")
	if f.CanSet() {
		f.SetString("reflect set")
	}

	fmt.Printf("%v", y)
}

type Data struct {
	f string
}

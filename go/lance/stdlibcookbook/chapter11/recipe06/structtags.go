// Accessing tags via reflection
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"p_name" bson:"pName"`
	Age  int    `json:"p_age" bson:"pAge"`
}

func main() {
	f := &Person{"Tom", 30}
	describe(f)
}

func describe(f interface{}) {
	val := reflect.TypeOf(f).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeF := val.Field(i)
		fieldName := typeF.Name
		jsonTag := typeF.Tag.Get("json")
		bsonTag := typeF.Tag.Get("bson")
		fmt.Printf("Field : %s jsonTag: %s bsonTag: %s\n", fieldName, jsonTag, bsonTag)
	}
}

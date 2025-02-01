package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greeting() string {
	return fmt.Sprintf("Hi, I am %s", p.Name)
}

func TestReflect(t *testing.T) {
	t.Run("reflect.TypeOf", func(t *testing.T) {
		var num int
		typ := reflect.TypeOf(num)

		if got, want := typ.Name(), "int"; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.ValueOf", func(t *testing.T) {
		var num int
		val := reflect.ValueOf(num)

		if v, ok := val.Interface().(int); !ok {
			t.Error("val should be int")
		} else if got, want := v, 0; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.Value.Kind", func(t *testing.T) {
		var num int64
		typ := reflect.ValueOf(num)

		if got, want := typ.Kind(), reflect.Int64; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.Value.FieldByName", func(t *testing.T) {
		var person any = Person{
			Age: 10,
		}
		val := reflect.ValueOf(person)
		ageVal := val.FieldByName("Age")

		if age, ok := ageVal.Interface().(int); !ok {
			t.Error("ageVal should be int")
		} else if got, want := age, 10; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.Value.MethodByName", func(t *testing.T) {
		var person any = Person{
			Name: "John",
			Age:  10,
		}
		val := reflect.ValueOf(person)
		greetingMethod := val.MethodByName("Greeting")

		if ret, ok := greetingMethod.Call([]reflect.Value{})[0].Interface().(string); !ok {
			t.Error("greeting should return string")
		} else if got, want := ret, "Hi, I am John"; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.New", func(t *testing.T) {
		intType := reflect.TypeOf(5)
		intVal := reflect.New(intType).Elem()
		intVal.SetInt(30)

		if got, want := intVal.Int(), int64(30); got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.Value.IsValid", func(t *testing.T) {
		val := reflect.ValueOf(5)

		if got, want := val.IsValid(), true; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.Value.IsNil", func(t *testing.T) {
		var ptr *int
		val := reflect.ValueOf(ptr)

		if got, want := val.IsNil(), true; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect.Value.NumField traverses struct", func(t *testing.T) {
		var person any = Person{Name: "John", Age: 10}
		personVal := reflect.ValueOf(person)

		for i := 0; i < personVal.NumField(); i++ {
			field := personVal.Field(i)

			switch fieldName := personVal.Type().Field(i).Name; fieldName {
			case "Name":
				if got, want := field.Interface(), "John"; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			case "Age":
				if got, want := field.Interface(), 10; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			default:
				t.Errorf("unknown field: %v", fieldName)
			}
		}
	})

	t.Run("reflect.Value.MethodByName checks if method exists", func(t *testing.T) {
		var person any = Person{
			Name: "John",
			Age:  10,
		}
		val := reflect.ValueOf(person)

		if got, want := val.MethodByName("Greeting").IsValid(), true; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
		if got, want := val.MethodByName("unknownMethod").IsValid(), false; got != want {
			t.Errorf("got:%v want:%v", got, want)
		}
	})

	t.Run("reflect and slice", func(t *testing.T) {
		var slice any = []int{2, 3, 5, 7, 11}
		sliceVal := reflect.ValueOf(slice)

		for i := 0; i < sliceVal.Len(); i++ {
			switch i {
			case 0:
				if got, want := sliceVal.Index(i).Interface(), 2; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			case 1:
				if got, want := sliceVal.Index(i).Interface(), 3; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			case 2:
				if got, want := sliceVal.Index(i).Interface(), 5; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			case 3:
				if got, want := sliceVal.Index(i).Interface(), 7; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			case 4:
				if got, want := sliceVal.Index(i).Interface(), 11; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			default:
				t.Errorf("unknown index: %v", i)
			}
		}
	})

	t.Run("reflect and map", func(t *testing.T) {
		var ma any = map[string]int{"a": 1, "b": 2}
		mapVal := reflect.ValueOf(ma)

		for _, keyVal := range mapVal.MapKeys() {
			v := mapVal.MapIndex(keyVal)
			switch key := keyVal.Interface(); key {
			case "a":
				if got, want := v.Interface(), 1; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			case "b":
				if got, want := v.Interface(), 2; got != want {
					t.Errorf("got:%v want:%v", got, want)
				}
			default:
				t.Errorf("unknown key: %v", key)
			}
		}
	})
}

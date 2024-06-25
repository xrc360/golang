package gstructs_test

import (
	"reflect"
	"testing"

	"github.com/xrc360/golang/os/gstructs"
)

type User struct {
	Id   int
	Name string `params:"name"`
	Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
}

var (
	user           = new(User)
	userNilPointer *User
)

func Benchmark_ReflectTypeOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.TypeOf(user).String()
	}
}

func Benchmark_TagFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gstructs.TagFields(user, []string{"params", "my-tag1"})
	}
}

func Benchmark_TagFields_NilPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gstructs.TagFields(&userNilPointer, []string{"params", "my-tag1"})
	}
}

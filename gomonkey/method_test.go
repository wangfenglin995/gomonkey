package gomonkey

import (
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMyStruct_StructAdd(t *testing.T) {
	var s *MyStruct
	patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "StructAdd", func(_ *MyStruct, a, b int) int {
		return 10
	})
	defer patches.Reset()

	res := s.StructAdd(2, 3)
	assert.Equal(t, 10, res)
}

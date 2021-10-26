package gomonkey

import (
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiAdd(t *testing.T) {
	//对函数AddOne打桩
	patches := gomonkey.ApplyFunc(Add10, func(t1 int32) int32 {
		return 0
	})
	defer patches.Reset()

	//对函数MinusOne打桩
	patches.ApplyFunc(Minus10, func(t1 int32) int32 {
		return -1
	})

	// MultiAddOne内部调用了Add10和Minus10方法
	result := MultiAdd(10)
	assert.Equal(t, int32(0), result)
}


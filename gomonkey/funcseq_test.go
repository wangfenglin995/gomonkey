package gomonkey

import (
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试给函数打一个桩序列
func TestAdd10Seq(t *testing.T) {
	// 调用3次分别返回10，20，30
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{int32(10)}},
		{Values: gomonkey.Params{int32(20)}},
		{Values: gomonkey.Params{int32(30)}},
	}
	//对函数AddOne打桩
	patches := gomonkey.ApplyFuncSeq(Add10, outputs)
	defer patches.Reset()

	// MultiAddOne内部调用了Add10和Minus10方法
	result := Add10(10)
	assert.Equal(t, int32(10), result)
	result = Add10(5)
	assert.Equal(t, int32(20), result)
	result = Add10(5)
	assert.Equal(t, int32(30), result)
}

// 测试mock掉函数内部调用的函数打一个桩序列
func TestMockInnerFunc(t *testing.T) {
	// 调用3次分别返回1，2，3
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{int32(1)}},
		{Values: gomonkey.Params{int32(2)}},
		{Values: gomonkey.Params{int32(3)}},
	}
	//对函数Add10打桩
	patches := gomonkey.ApplyFuncSeq(Add10, outputs)
	defer patches.Reset()

	result := DoubleAdd10(10)
	assert.Equal(t, int32(11), result)
	result = DoubleAdd10(10)
	assert.Equal(t, int32(12), result)
	result = DoubleAdd10(10)
	assert.Equal(t, int32(13), result)
}
package gomonkey

import (
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	num = 10
)

// 测试mock全局变量
func TestMockGlobalVariable(t *testing.T) {
	patches := gomonkey.ApplyGlobalVar(&num, 20)
	defer patches.Reset()

	assert.Equal(t, 20, 20)
}
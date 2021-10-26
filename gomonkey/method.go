package gomonkey

type MyStruct struct {
}

func (s *MyStruct) StructAdd(a, b int) int {
	return a + b
}

package gomonkey

func Add10(t int32) int32 {
	return t + 10
}

func Minus10(t int32) int32 {
	return t - 10
}

func MultiAdd(t int32) int32 {
	t = Minus10(t)
	t = Add10(t)
	t = Add10(t)
	return t
}

func DoubleAdd10(t int32) int32 {
	return Add10(10) + 10
}
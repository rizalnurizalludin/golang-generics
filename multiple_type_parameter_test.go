package golang_generic

import "testing"

func MultipleParameter[T1, T2 any](param T1, param2 T2) (T1, T2) {
	return param, param2
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[int, string](1, "Rizal")
	MultipleParameter[string, int]("Indomie", 4)
}

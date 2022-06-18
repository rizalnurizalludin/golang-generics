package golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	assert.True(t,true)
}

func TestLength(t *testing.T) {
	result := Length[string]("Rizal")
	assert.Equal(t,"Rizal", result)

	resultNumber := Length[int](100)
	assert.Equal(t,100,resultNumber)
}

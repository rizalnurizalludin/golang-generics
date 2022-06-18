package golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T){
	assert.True(t,IsSame[string]("Rizal","Rizal"))
	assert.False(t, IsSame[int](12,21))
}

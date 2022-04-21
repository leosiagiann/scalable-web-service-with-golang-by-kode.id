package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 30 {
// 		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 20)
// 	}
// }

// func TestSum2(t *testing.T) {
// 	result := Sum(10, 10)

// 	// require.Equal(t, 30, result, "Sum was incorrect, got: %d, want: %d.", result, 20)
// 	assert.Equal(t, 30, result, "Sum was incorrect, got: %d, want: %d.", result, 20)
// }

func TestTableSum(t *testing.T) {
	var tests = []struct {
		request  int
		expected int
		errMsg   string
	}{
		{Sum(10, 10), 20, "result should be 20"},
		{Sum(10, 10), 30, "result should be 30"},
		{Sum(10, 10), 40, "result should be 40"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, "result should be %d", test.expected)
		})
	}
}

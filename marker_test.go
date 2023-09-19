package marker

import (
	"math/rand"
	"testing"
	"time"
)

func TestMarker(t *testing.T) {
	marker := New()
	var err error

	marker.Mark("Event 1")
	generateSlice(1000000, 10000000)
	err = marker.Done()
	handleError(t, err)

	marker.Mark("")
	generateSlice(1000000, 50000000)
	err = marker.Done()
	handleError(t, err)

	marker.Mark("Event 3")
	generateSlice(1000000, 100000000)
	err = marker.Done()
	handleError(t, err)

	result, err := marker.String()
	handleError(t, err)
	t.Logf("Moments:\n%s", result)
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}
}

// generates a slice of n integers with random values ranging from 0 to m-1.
func generateSlice(m, n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(m))
	}
	return result
}

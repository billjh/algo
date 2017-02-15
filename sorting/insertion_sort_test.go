package insertion

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestInsertionSort10(t *testing.T) {
	arr := utils.GetIntArray(10)
	sort(arr)
	if !utils.IsSorted(arr) {
		t.Error("array is not sorted:", arr)
	}
}

func TestInsertionSort100(t *testing.T) {
	arr := utils.GetIntArray(100)
	sort(arr)
	if !utils.IsSorted(arr) {
		t.Error("array is not sorted:", arr)
	}
}

func TestInsertionSort1000(t *testing.T) {
	arr := utils.GetIntArray(1000)
	sort(arr)
	if !utils.IsSorted(arr) {
		t.Error("array is not sorted:", arr)
	}
}
func TestInsertionSort10000(t *testing.T) {
	arr := utils.GetIntArray(10000)
	sort(arr)
	if !utils.IsSorted(arr) {
		t.Error("array is not sorted:", arr)
	}
}

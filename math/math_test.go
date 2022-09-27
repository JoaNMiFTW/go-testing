package math

import "testing"

type AddData struct {
	x, y, result int
}

func TestAdd(t *testing.T) {
	// result := Add(1, 3)

	// if result != 4 {
	// 	t.Errorf("Add (1,3) FAILED. Expected %d, got %d\n", 4, result)
	// }

	// t.Logf("Add(1,3) PASSED. Expected %d, got %d\n", 4, result)

	testData := []AddData{
		{1, 2, 3},
		{3, 5, 8},
		{7, 5, 12},
		{3, -4, -1},
		{2, 2, 4},
		{-1, 4, 3},
	}

	for _, data := range testData {
		result := Add(data.x, data.y)

		if result != data.result {
			t.Errorf("Add(%d, %d) FAILED. Expected %d got %d\n", data.x, data.y, data.result, result)
		}

		t.Logf("Add(%d, %d) PASSED. Expected %d got %d\n", data.x, data.y, data.result, result)
	}
}

func TestSubtract(t *testing.T) {

	testData := []AddData{
		{1, 2, -1},
		{3, 5, -2},
		{7, 5, 2},
		{3, -4, 7},
		{2, 2, 0},
		{-1, 4, -5},
	}

	for _, data := range testData {
		result := Subtract(data.x, data.y)

		if result != data.result {
			t.Errorf("Subtract(%d, %d) FAILED. Expected %d got %d\n", data.x, data.y, data.result, result)
		}

		t.Logf("Subtract(%d, %d) PASSED. Expected %d got %d\n", data.x, data.y, data.result, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("Divide (5,0) FAILED. Expected %f, got %f\n", 0., result)
	}

	t.Logf("Divide(5,0) PASSED. Expected %f, got %f\n", 0., result)
}

func TestMultiply(t *testing.T) {
	testData := []AddData{
		{1, 2, 2},
		{3, 5, 15},
		{7, 5, 35},
		{3, -4, -12},
		{2, 2, 4},
		{-1, 4, -4},
	}

	for _, data := range testData {
		result := Multiply(data.x, data.y)

		if result != data.result {
			t.Errorf("Multiply(%d, %d) FAILED. Expected %d got %d\n", data.x, data.y, data.result, result)
		}

		t.Logf("Multiply(%d, %d) PASSED. Expected %d got %d\n", data.x, data.y, data.result, result)
	}
}

package functions

import "testing"

func TestMinElement(t *testing.T) {
	x := MinElement(1, -2, 5)
	res := -2
	if x != res {
		t.Errorf("Тест не пройдено! Результат %d, а повинен бути %d", x, res)
	}
}

func TestAverage(t *testing.T) {
	x := Average(15, -4, 13)
	res := 8
	if x != res {
		t.Errorf("Тест не пройдено! Результат %d, а повинен бути %d", x, res)
	}
}

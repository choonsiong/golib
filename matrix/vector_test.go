package matrix

import "testing"

func TestRowVector_GetShape(t *testing.T) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3, 4.4})
	t.Log(rv1.GetShape())
}

func TestRowVector_Mul(t *testing.T) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3})
	rv1.Mul(2.0)
	t.Log("rv1=", rv1)
}

func TestNewRowVector(t *testing.T) {
	t.Log(NewRowVector(5))
}

func TestRowVector_Add(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 4})
	rv2 := RowVector([]float64{14, 17, 16})
	rv3, _ := rv1.Add(rv2)
	t.Log("rv3=", rv3)
}

func TestRowVector_Minus(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 4})
	rv2 := RowVector([]float64{14, 17, 16})
	rv3, _ := rv1.Minus(rv2)
	t.Log("rv3=", rv3)
}

func TestRowVector_Dot(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 4})
	rv2 := RowVector([]float64{14, 17, 16})
	dot, _ := rv1.Dot(rv2)
	t.Log("dot=", dot)
}

func TestRowVector_Cross(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 4})
	rv2 := RowVector([]float64{14, 17, 16})
	rv3, _ := rv1.Cross(rv2)
	t.Log("rv3=", rv3)
}

func TestRowVector_Length(t *testing.T) {
	rv := RowVector([]float64{1, 2, 3})
	l := rv.Length()
	t.Log("l=", l)
}

func TestRowVector_Transpose(t *testing.T) {
	rv := RowVector([]float64{1.1, 2.2, 3.3})
	cv := rv.Transpose()
	t.Log("cv=", cv)
}

func BenchmarkRowVector_GetShape(b *testing.B) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3, 4.4})
	for i := 0; i < b.N; i++ {
		rv1.GetShape()
	}
}

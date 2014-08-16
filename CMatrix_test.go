package cmatrix

import "testing"

func TestSliceCMatrixInit(t *testing.T) {
	r, c := 2, 3
	M := InitSliceCMatrix(r, c)
	rD, cD := M.Dims()
	if r != rD || c != cD {
		t.Fatalf("Matrix created with InitSliceCMatrix reports incorrect dimensions.")
	}
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if M.At(i, j) != complex(0.0, 0.0) {
				t.Fatalf("Failed to initialize SliceCMatrix to 0 at r=%d, c=%d.", r, c)
			}
		}
	}
}

func TestSliceCMatrixSet(t *testing.T) {
	r, c := 2, 3
	M := InitSliceCMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			val := complex(float64(i + j), float64(2.0 * (i + j)))
			M[i][j] = val
			if M.At(i, j) != val {
				t.Fatalf("Failed to correctly set M at r=%d, c=%d.", r, c)
			}
		}
	}
}

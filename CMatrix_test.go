package cmatrix

import (
	"testing"
)

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
			val := complex(float64(i+j), float64(2.0*(i+j)))
			M[i][j] = val
			if M.At(i, j) != val {
				t.Fatalf("Failed to correctly set M at r=%d, c=%d.", r, c)
			}
		}
	}
}

func TestSliceCMatrixAddTo(t *testing.T) {
	r, c := 2, 3
	M, A := InitSliceCMatrix(r, c), InitSliceCMatrix(r, c)
	A[1][1] = complex(3.0, 2.0)
	A.AddTo(&M)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 1 && j == 1 {
				if M[i][j] != complex(3.0, 2.0) {
					t.Fatalf("Incorrect M value")
				}
			} else {
				if M[i][j] != complex(0.0, 0.0) {
					t.Fatalf("Incorrect M value")
				}
			}
		}
	}
}

func TestSliceCMatrixAddMulTo(t *testing.T) {
	r, c := 2, 3
	M, A := InitSliceCMatrix(r, c), InitSliceCMatrix(r, c)
	A[1][1] = complex(3.0, 2.0)
	A.AddMulTo(&M, 2.0)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 1 && j == 1 {
				if M[i][j] != complex(6.0, 4.0) {
					t.Fatalf("Incorrect M value")
				}
			} else {
				if M[i][j] != complex(0.0, 0.0) {
					t.Fatalf("Incorrect M value")
				}
			}
		}
	}
}

func TestSliceCMatrixMulBy(t *testing.T) {
	r, c := 2, 3
	M := InitSliceCMatrix(r, c)
	M[1][1] = complex(3.0, 4.0)
	(&M).MulBy(complex(3.0, -4.0))
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 1 && j == 1 {
				if M[i][j] != complex(25.0, 0.0) {
					t.Fatalf("Incorrect M value")
				}
			} else {
				if M[i][j] != complex(0.0, 0.0) {
					t.Fatalf("Incorrect M value")
				}
			}
		}
	}
}

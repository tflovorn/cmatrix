package cmatrix

import (
	"bytes"
	"fmt"
)

// CMatrix represents a matrix storing complex128 values.
// Follows the Matrix interface of gonum, which is found at:
// 	https://github.com/gonum/matrix/blob/master/mat64/matrix.go
type CMatrix interface {
	// Dims returns the dimensions of a CMatrix.
	Dims() (r, c int)

	// At returns the value of the matrix element at (r, c). It will
	// panic if r or c are out of bounds for the matrix.
	At(r, c int) complex128

	String() string
}

type SliceCMatrix [][]complex128

// Create a SliceCMatrix of dimensions r x c with all values initialized to 0.
func InitSliceCMatrix(r, c int) SliceCMatrix {
	M := make([][]complex128, r)
	for i := 0; i < r; i++ {
		M[i] = make([]complex128, c)
	}
	return M
}

// Assume all inner slices have the same length.
func (M SliceCMatrix) Dims() (r, c int) {
	r = len(M)
	c = len(M[0])
	return r, c
}

func (M SliceCMatrix) At(r, c int) complex128 {
	return M[r][c]
}

// String representation:
// "row1_col1 row1_col2 ... row1_colN \n row2_col1 ..."
func (M SliceCMatrix) String() string {
	r, c := M.Dims()
	var buffer bytes.Buffer

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			buffer.WriteString(fmt.Sprint(M.At(i, j)))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// Add M to A in-place (i.e. A --> A + M).
func (M SliceCMatrix) AddTo(A *SliceCMatrix) {
	mr, mc := M.Dims()
	ar, ac := A.Dims()
	if mr != ar || mc != ac {
		panic("Adding incompatible matrices")
	}
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			(*A)[i][j] += M[i][j]
		}
	}
}

// Add M*val to A in-place (i.e. A --> A + val*M).
func (M SliceCMatrix) AddMulTo(A *SliceCMatrix, val complex128) {
	mr, mc := M.Dims()
	ar, ac := A.Dims()
	if mr != ar || mc != ac {
		panic("Adding incompatible matrices")
	}
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			(*A)[i][j] += val * M[i][j]
		}
	}
}

// Multiply M by the given scalar in-place (i.e. M --> val*M).
func (M *SliceCMatrix) MulBy(val complex128) {
	mr, mc := M.Dims()
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			(*M)[i][j] *= val
		}
	}
}

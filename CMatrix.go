package cmatrix

// CMatrix represents a matrix storing complex128 values.
// Follows the Matrix interface of gonum, which is found at:
// 	https://github.com/gonum/matrix/blob/master/mat64/matrix.go
type CMatrix interface {
	// Dims returns the dimensions of a CMatrix.
	Dims() (r, c int)

	// At returns the value of the matrix element at (r, c). It will
	// panic if r or c are out of bounds for the matrix.
	At(r, c int) complex128
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

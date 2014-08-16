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

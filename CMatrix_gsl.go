package cmatrix

/*
#cgo LDFLAGS: -lgsl -lgslcblas
#include <gsl/gsl_complex.h>
#include <gsl/gsl_complex_math.h>
#include <gsl/gsl_vector.h>
#include <gsl/gsl_matrix.h>
#include <gsl/gsl_eigen.h>
*/
import "C"

// Perform eigendecomposition of the Hermitian matrix M.
// Returns a list of eigenvalues and a list of eigenvectors.
// Eigenvectors are ordered corresponding to the eigenvalues.
func Eigensystem(M CMatrix) (evals []float64, evecs [][]complex128) {
	// Convert M to GSL representation.
	Mgsl := HermToGSL(M)
	// Allocate space for eigendecomposition (needed for gsl_eigen_hermv call).
	r, _ := M.Dims()
	rC := C.size_t(r)
	work := C.gsl_eigen_hermv_alloc(rC)
	evalsgsl := C.gsl_vector_alloc(rC)
	evecsgsl := C.gsl_matrix_complex_alloc(rC, rC)
	// Do eigendecomposition.
	C.gsl_eigen_hermv(Mgsl, evalsgsl, evecsgsl, work)
	// Convert data back to Go format.
	evals = GSLvectorToSlice(evalsgsl)
	evecs = GSLmatrixColumnsToSlices(evecsgsl)
	// Clean up.
	C.gsl_eigen_hermv_free(work)
	C.gsl_vector_free(evalsgsl)
	C.gsl_matrix_complex_free(Mgsl)
	C.gsl_matrix_complex_free(evecsgsl)
	return evals, evecs
}

// Construct a Hermitian gsl_matrix_complex from the diagonal and
// lower triangular parts of M.
// When done using the returned gsl_matrix_complex, it must be manually
// freed using gsl_matrix_complex_free.
func HermToGSL(M CMatrix) (Mgsl *C.gsl_matrix_complex) {
	r, _ := M.Dims()
	Mgsl = C.gsl_matrix_complex_alloc(C.size_t(r), C.size_t(r))
	for i := 0; i < r; i++ {
		for j := 0; j <= i; j++ {
			val := M.At(i, j)
			val_re, val_im := real(val), imag(val)
			// set value at (i, j)
			valgsl := C.gsl_complex_rect(C.double(val_re), C.double(val_im))
			C.gsl_matrix_complex_set(Mgsl, C.size_t(i), C.size_t(j), valgsl)
			// set value at (j, i) given by complex conjugate of val
			valgsl_t := C.gsl_complex_rect(C.double(val_re), C.double(-val_im))
			C.gsl_matrix_complex_set(Mgsl, C.size_t(j), C.size_t(i), valgsl_t)
		}
	}
	return Mgsl
}

// Convert the GSL vector v to a slice of float64 values.
func GSLvectorToSlice(v *C.gsl_vector) []float64 {
	xs := []float64{}
	var i C.size_t
	for i = 0; i < v.size; i++ {
		xs = append(xs, float64(C.gsl_vector_get(v, i)))
	}
	return xs
}

// Convert the complex GSL matrix m to a slice of slices representing the
// columns of the matrix m.
// The first index of the output corresponds to the column index of m, while the
// second index of the output corresponds to the row index of m.
func GSLmatrixColumnsToSlices(m *C.gsl_matrix_complex) [][]complex128 {
	vectors := [][]complex128{}
	var i, j C.size_t
	for i = 0; i < m.size1; i++ {
		xs := []complex128{}
		for j = 0; j < m.size2; j++ {
			valgsl := C.gsl_matrix_complex_get(m, j, i)
			val_re, val_im := float64(valgsl.dat[0]), float64(valgsl.dat[1])
			val := complex(val_re, val_im)
			xs = append(xs, val)
		}
		vectors = append(vectors, xs)
	}
	return vectors
}

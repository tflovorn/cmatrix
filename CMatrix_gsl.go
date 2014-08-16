package cmatrix

/*
#cgo LDFLAGS: -lgsl -lgslcblas
#include <gsl/gsl_complex.h>
#include <gsl/gsl_complex_math.h>
#include <gsl/gsl_matrix.h>
#include <gsl/gsl_eigen.h>
*/
import "C"

func Eigensystem(M CMatrix) (evals []float64, evecs [][]complex128) {
	return nil, nil
}

// Construct a Hermitian gsl_matrix_complex from the diagonal and
// lower triangular parts of M.
// When done using the returned gsl_matrix_complex, it must be manually
// freed using gsl_matrix_complex_free.
func HermToGSL(M CMatrix) (Mgsl *C.gsl_matrix_complex) {
	r, c := M.Dims()
	Mgsl = C.gsl_matrix_complex_alloc(C.size_t(r), C.size_t(c))
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

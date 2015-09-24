package cmatrix

import (
	"math"
	"math/cmplx"
	"testing"
)

func TestEigenPauli(t *testing.T) {
	eps := 1e-9
	Ps := MakePaulis()
	for i, P := range Ps {
		evals, evecs := Eigensystem(P)
		if math.Abs(evals[0]-1.0) > eps {
			t.Fatalf("Incorrect eigenvalue obtained for Pi, i=%d\n", i)
		}
		if math.Abs(evals[1]+1.0) > eps {
			t.Fatalf("Incorrect eigenvalue obtained for Pi, i=%d\n", i)
		}
		if i == 1 {
			CheckXEigenvecs(t, evecs, eps)
		} else if i == 2 {
			CheckYEigenvecs(t, evecs, eps)
		} else {
			CheckZEigenvecs(t, evecs, eps)
		}
	}
}

func MakePaulis() []SliceCMatrix {
	Px := InitSliceCMatrix(2, 2)
	Px[0] = []complex128{complex(0.0, 0.0), complex(1.0, 0.0)}
	Px[1] = []complex128{complex(1.0, 0.0), complex(0.0, 0.0)}
	Py := InitSliceCMatrix(2, 2)
	Py[0] = []complex128{complex(0.0, 0.0), complex(0.0, -1.0)}
	Py[1] = []complex128{complex(0.0, 1.0), complex(0.0, 0.0)}
	Pz := InitSliceCMatrix(2, 2)
	Pz[0] = []complex128{complex(1.0, 0.0), complex(0.0, 0.0)}
	Pz[1] = []complex128{complex(0.0, 0.0), complex(-1.0, 0.0)}
	return []SliceCMatrix{Px, Py, Pz}
}

func CheckXEigenvecs(t *testing.T, evecs [][]complex128, eps float64) {
	v := complex(1.0/math.Sqrt(2.0), 0.0)
	vec1_ok := cmplx.Abs(evecs[0][0]-v) > eps || cmplx.Abs(evecs[0][1]-v) > eps
	vec2_ok := cmplx.Abs(evecs[1][0]-v) > eps || cmplx.Abs(evecs[1][1]+v) > eps
	if !vec1_ok || !vec2_ok {
		t.Fatalf("Got incorrect eigenvectors %v for x component.", evecs)
	}
}

func CheckYEigenvecs(t *testing.T, evecs [][]complex128, eps float64) {
	v := complex(1.0/math.Sqrt(2.0), 0.0)
	v_i := complex(0.0, 1.0/math.Sqrt(2.0))
	vec1_ok := cmplx.Abs(evecs[0][0]-v) > eps || cmplx.Abs(evecs[0][1]-v_i) > eps
	vec2_ok := cmplx.Abs(evecs[1][0]-v) > eps || cmplx.Abs(evecs[1][1]+v_i) > eps
	if !vec1_ok || !vec2_ok {
		t.Fatalf("Got incorrect eigenvectors %v for y component.", evecs)
	}
}

func CheckZEigenvecs(t *testing.T, evecs [][]complex128, eps float64) {
	v := complex(1.0, 0.0)
	vec1_ok := cmplx.Abs(evecs[0][0]-v) > eps || cmplx.Abs(evecs[0][1]) > eps
	vec2_ok := cmplx.Abs(evecs[1][0]) > eps || cmplx.Abs(evecs[1][1]-v) > eps
	if !vec1_ok || !vec2_ok {
		t.Fatalf("Got incorrect eigenvectors %v for z component.", evecs)
	}
}

func TestEigenPauliGSL(t *testing.T) {
	eps := 1e-9
	Ps := MakePaulisGSL()
	for i, P := range Ps {
		work, evals, evecs_gsl := HermEigensystemSetup(P)
		HermEigensystem(P, work, evals, evecs_gsl)
		if math.Abs(evals.At(0)-1.0) > eps {
			t.Fatalf("Incorrect eigenvalue obtained for Pi, i=%d\n", i)
		}
		if math.Abs(evals.At(1)+1.0) > eps {
			t.Fatalf("Incorrect eigenvalue obtained for Pi, i=%d\n", i)
		}
		evecs := GSLmatrixColumnsToSlices(evecs_gsl.M)
		if i == 1 {
			CheckXEigenvecs(t, evecs, eps)
		} else if i == 2 {
			CheckYEigenvecs(t, evecs, eps)
		} else {
			CheckZEigenvecs(t, evecs, eps)
		}
		HermEigensystemCleanup(work, evals, evecs_gsl)
		P.Destroy()
	}
}

func MakePaulisGSL() []*CMatrixGSL {
	Px := NewCMatrixGSL(2, 2)
	Px.Set(0, 0, complex(0.0, 0.0))
	Px.Set(0, 1, complex(1.0, 0.0))
	Px.Set(1, 0, complex(1.0, 0.0))
	Px.Set(1, 1, complex(0.0, 0.0))
	Py := NewCMatrixGSL(2, 2)
	Py.Set(0, 0, complex(0.0, 0.0))
	Py.Set(0, 1, complex(0.0, -1.0))
	Py.Set(1, 0, complex(0.0, 1.0))
	Py.Set(1, 1, complex(0.0, 0.0))
	Pz := NewCMatrixGSL(2, 2)
	Pz.Set(0, 0, complex(1.0, 0.0))
	Pz.Set(0, 1, complex(0.0, 0.0))
	Pz.Set(1, 0, complex(0.0, 0.0))
	Pz.Set(1, 1, complex(-1.0, 0.0))
	return []*CMatrixGSL{Px, Py, Pz}
}

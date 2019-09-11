package main

import "math/big"

var currentPrecision uint = 128

// Pow represents the pow function
func Pow(a *big.Float, e uint) *big.Float {
	e = uint(e)
	result := ZeroBigFloat().Copy(a)
	for i := uint(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

// Pow3 represents pow(x, 3)
func Pow3(in *big.Float) *big.Float {
	return Pow(in, 3)
}

// Pow4 represents pow(x, 4)
func Pow4(in *big.Float) *big.Float {
	return Pow(in, 4)
}

// Pow5 represents pow(x, 5)
func Pow5(in *big.Float) *big.Float {
	return Pow(in, 5)
}

// Pow6 represents pow(x, 6)
func Pow6(in *big.Float) *big.Float {
	return Pow(in, 6)
}

// Pow7 represents pow(x, 7)
func Pow7(in *big.Float) *big.Float {
	return Pow(in, 7)
}

// Pow8 represents pow(x, 8)
func Pow8(in *big.Float) *big.Float {
	return Pow(in, 8)
}

// Pow9 represents pow(x, 9)
func Pow9(in *big.Float) *big.Float {
	return Pow(in, 9)
}

// Pow10 represents pow(x, 10)
func Pow10(in *big.Float) *big.Float {
	return Pow(in, 10)
}

// Pow11 represents pow(x, 11)
func Pow11(in *big.Float) *big.Float {
	return Pow(in, 11)
}

// Pow12 represents pow(x, 12)
func Pow12(in *big.Float) *big.Float {
	return Pow(in, 12)
}

// Root represets root(x, n)
func Root(a *big.Float, n uint64) *big.Float {
	limit := Pow(NewBigFloat(2), currentPrecision)
	n1 := n - 1
	n1f, rn := NewBigFloat(float64(n1)), Div(NewBigFloat(1.0), NewBigFloat(float64(n)))
	x, x0 := NewBigFloat(1.0), ZeroBigFloat()
	_ = x0
	for {
		potx, t2 := Div(NewBigFloat(1.0), x), a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 = Mul(t2, potx)
			}
			potx = Mul(potx, potx)
		}
		x0, x = x, Mul(rn, Add(Mul(n1f, x), t2))
		if Lesser(Mul(Abs(Sub(x, x0)), limit), x) {
			break
		}
	}
	return x
}

// Root3 represents root(x, 3)
func Root3(in *big.Float) *big.Float {
	return Root(in, 3)
}

// Root4 represents root(x, 4)
func Root4(in *big.Float) *big.Float {
	return Root(in, 4)
}

// Root5 represents root(x, 5)
func Root5(in *big.Float) *big.Float {
	return Root(in, 5)
}

// Root6 represents root(x, 6)
func Root6(in *big.Float) *big.Float {
	return Root(in, 6)
}

// Root7 represents root(x, 7)
func Root7(in *big.Float) *big.Float {
	return Root(in, 7)
}

// Root8 represents root(x, 8)
func Root8(in *big.Float) *big.Float {
	return Root(in, 8)
}

// Root9 represents root(x, 9)
func Root9(in *big.Float) *big.Float {
	return Root(in, 9)
}

// Root10 represents root(x, 10)
func Root10(in *big.Float) *big.Float {
	return Root(in, 10)
}

// Root11 represents root(x, 11)
func Root11(in *big.Float) *big.Float {
	return Root(in, 11)
}

// Root12 represents root(x, 12)
func Root12(in *big.Float) *big.Float {
	return Root(in, 12)
}

// Abs represents the abs of a big float
func Abs(a *big.Float) *big.Float {
	return ZeroBigFloat().Abs(a)
}

// NewBigFloat returns a new big float with current global precision
func NewBigFloat(f float64) *big.Float {
	r := big.NewFloat(f)
	r.SetPrec(currentPrecision)
	return r
}

// Div devides two big floats
func Div(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Quo(a, b)
}

// ZeroBigFloat returns a big float with zero value
func ZeroBigFloat() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(currentPrecision)
	return r
}

// Mul multiplies two big floats
func Mul(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Mul(a, b)
}

// Add adds two big floats
func Add(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Add(a, b)
}

// Sub subs two big floats
func Sub(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Sub(a, b)
}

// Lesser returns a bool
func Lesser(x, y *big.Float) bool {
	return x.Cmp(y) == -1
}

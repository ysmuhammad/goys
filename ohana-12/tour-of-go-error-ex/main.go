package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	t, z := 0., 1.
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-8 {
			break
		}
	}
	return z, nil
}

func main() {
	result1, err := Sqrt(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}

	result2, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
}

package triangolo

import (
	"fmt"
	"math"
)

type Triangolo struct {
	lato1, lato2, lato3 float64
}

func NuovoTriangolo(l1, l2, l3 float64) (t Triangolo, ok bool) {
	if (l1+l2 > l3) && (l1+l3 > l2) && (l2+l3 > l1) {
		return Triangolo{l1, l2, l3}, true
	}
	return Triangolo{}, false
}

func Perimetro(t Triangolo) float64 {
	return t.lato1 + t.lato2 + t.lato3
}

func Area(t Triangolo) float64 {
	p := (t.lato1 + t.lato2 + t.lato3) / 2
	return math.Sqrt(p * (p - t.lato1) * (p - t.lato2) * (p - t.lato3))
}

func String(t Triangolo) string {
	return fmt.Sprintf("Triangolo con lati %.2f, %.2f e %.2f.", t.lato1, t.lato2, t.lato3)
}

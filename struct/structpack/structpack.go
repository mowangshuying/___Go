package structpack

import (
	"fmt"
)

type ExpStruct struct {
	Mi int
	Mf float32
}

func (s *ExpStruct)logself() {
	fmt.Printf("Mi:%d, Mf:%d\n", s.Mi, s.Mf)
}
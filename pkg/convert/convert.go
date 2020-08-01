package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	return strconv.Atoi(s.String())
}

func (s StrTo) MustInt() int {
	v, _ := s.Int() // ignore error
	return v
}

func (s StrTo) Uint32() (uint32, error) {
	v, err := s.Int()
	return uint32(v), err
}

func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32() // ignore error
	return v
}

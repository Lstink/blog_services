package convert

import "strconv"

type Strto string

func (s Strto) String() string {
	return string(s)
}

func (s Strto) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s Strto) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s Strto) Uint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s Strto) MustUInt32() uint32 {
	v, _ := s.Uint32()
	return v
}

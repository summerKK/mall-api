package convert

import "strconv"

type StrTo string

// 转换为 string
func (s StrTo) String() string {
	return string(s)
}

// 转换为 int
func (s StrTo) Int() (int, error) {
	return strconv.Atoi(s.String())
}

// 强制转换为 int
func (s StrTo) MustInt() int {
	v, _ := s.Int()

	return v
}

// 转换为 uint32
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	if err != nil {
		return 0, err
	}

	return uint32(v), err
}

// 强制装换为 uint32
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()

	return v
}

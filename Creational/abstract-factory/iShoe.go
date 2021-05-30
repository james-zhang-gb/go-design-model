package main

type iShoe interface {
	setLogo(logo string)
	getLogo() string
	getSize() int
	setSize(size int)
	// String() string
}
type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}

// func (s *shoe) String() string {
// 	return fmt.Sprintf("kind:%s logo:%s size:%d", "shoe", s.logo, s.size)
// }

var _ iShoe = &shoe{}

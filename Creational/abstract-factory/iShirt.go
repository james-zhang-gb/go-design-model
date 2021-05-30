package main

import "fmt"

type iShirt interface {
	setLogo(logo string)
	getLogo() string
	getSize() int
	setSize(size int)
	String() string
}
type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getSize() int {
	return s.size
}

func (s *shirt) String() string {
	return fmt.Sprintf("kind:%s logo:%s size:%d", "shirt", s.logo, s.size)
}

var _ iShirt = &shirt{}

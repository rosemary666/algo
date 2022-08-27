package stack

import "testing"

func TestStackOnLinklistPush(t *testing.T) {
	s := NewStackOnLinklist()
	s.Print()

	for i := 0; i < 20; i++ {
		s.Push(i)
	}
	s.Print()
}

func TestStackOnLinklistPop(t *testing.T) {
	s := NewStackOnLinklist()
	s.Print()

	for i := 0; i < 20; i++ {
		s.Push(i)
	}

	for i := 0; i < 20; i++ {
		v := s.Pop()
		t.Logf("pop elem is:[%v]", v)
	}
	s.Print()
}

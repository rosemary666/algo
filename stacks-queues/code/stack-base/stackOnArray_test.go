package stack

import "testing"

func TestStackOnArrayPush(t *testing.T) {
	s := NewStackOnArray()
	s.Print()

	t.Log("start push...")
	for i := 0; i < 20; i++ {
		s.Push(i)
	}
	s.Print()
}

func TestStackOnArrayPop(t *testing.T) {
	s := NewStackOnArray()
	s.Print()

	t.Log("start push...")
	for i := 0; i < 20; i++ {
		s.Push(i)
	}
	s.Print()
	println()
	println()
	t.Log("start pop...")
	for i := 0; i < 20; i++ {
		v := s.Pop()
		t.Logf("pop elem is:%v", v)
	}
	s.Print()
}

func TestStackOnArrayTop(t *testing.T) {
	s := NewStackOnArray()
	s.Print()

	t.Log("start push...")
	for i := 0; i < 20; i++ {
		s.Push(i)
	}

	t.Log("start top...")
	if s.Top().(int) != 19 {
		t.Fatalf("top fail!")
	}
	s.Print()
}

func TestStackOnArrayEmpty(t *testing.T) {
	s := NewStackOnArray()

	if !s.IsEmpty() {
		t.Fatalf("empty error!")
	}

	for i := 0; i < 20; i++ {
		s.Push(i)
	}

	if s.IsEmpty() {
		t.Fatalf("empty error!")
	}

	for i := 0; i < 20; i++ {
		s.Pop()
	}
	if !s.IsEmpty() {
		t.Fatalf("empty error!")
	}
}

func TestStackOnArrayReset(t *testing.T) {
	s := NewStackOnArray()

	if !s.IsEmpty() {
		t.Fatalf("empty error!")
	}

	for i := 0; i < 20; i++ {
		s.Push(i)
	}

	s.Reset()

	if !s.IsEmpty() {
		t.Fatalf("empty error!")
	}
}

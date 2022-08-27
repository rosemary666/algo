package circlelist

import "testing"

func TestRPush(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.RPush(i)
	}
	t.Log(list.GetLen())
	list.Print()
}

func TestGetByIndex(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.RPush(i)
	}
	for i := 0; i < 20; i++ {
		n := list.GetByIndex(uint(i))
		t.Log(n.Data)
	}
}

func TestRPop(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.RPush(i)
	}
	for {
		n := list.RPop()
		if n == nil {
			return
		}
		t.Log(n.Data)
	}
}

package doublelist

import "testing"

func TestRPush(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.RPush(i)
	}
	t.Logf("size:%v", list.Len())
	list.Print()
}

func TestLPush(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.LPush(i)
	}
	t.Logf("size:%v", list.Len())
	list.Print()
}

func TestRPop(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.RPush(i)
	}
	for {
		v := list.RPop()
		if v == nil {
			break
		}
		t.Logf("pop node value is:%v", v.Data)
	}
	t.Logf("size:%v", list.Len())
	list.Print()
}

func TestLPop(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.RPush(i)
	}
	for {
		v := list.LPop()
		if v == nil {
			break
		}
		t.Logf("pop node value is:%v", v.Data)
	}
	t.Logf("size:%v", list.Len())
	list.Print()
}

func TestFindByIndex(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.RPush(i)
	}
	for i := 0; i < 20; i++ {
		node := list.FindByIndex(uint(i))
		t.Logf("node value:%v", node.Data)
	}
}

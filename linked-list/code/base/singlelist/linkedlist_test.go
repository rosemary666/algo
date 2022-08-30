package singlelist

import (
	"testing"
)

func TestListInsertAfter(t *testing.T) {
	list := NewLinkList()
	curNode := list.head
	for i := 0; i < 20; i++ {
		list.InsertAfter(curNode, i)
		curNode = curNode.Next()
	}
	t.Log(list.Len())
	list.Print()
}

func TestListInsertToTail(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	t.Log(list.Len())
	list.Print()
}

func TestListInsertToHead(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.InsertToHead(i)
	}
	t.Log(list.Len())
	list.Print()
}

func TestListFront(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	t.Log(list.Len())
	t.Log(list.Front().Data)
}

func TestListBacl(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	t.Log(list.Len())
	t.Log(list.Back().Data)
}

func TestListFindByIndex(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 20; i++ {
		list.InsertToTail(i)
	}
	t.Log(list.Len())
	for i := 0; i < 20; i++ {
		t.Logf("[index/value]:[%v/%v]", i, list.FindByIndex(uint(i)).Value())
	}
}

func TestListRemove(t *testing.T) {
	list := NewLinkList()
	for i := 0; i < 5; i++ {
		list.InsertToTail(i)
	}
	totalLen := list.Len()
	t.Logf("[curlen/totallen]:[%v/%v]", list.Len(), totalLen)
	for i := 4; i >= 0; i-- {
		list.Remove(list.FindByIndex(uint(i)))
		t.Logf("[curlen/totallen]:[%v/%v]", list.Len(), totalLen)
	}
	list.Print()
}

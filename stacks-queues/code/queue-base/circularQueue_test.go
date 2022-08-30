package queue

import "testing"

func TestCircularQueue_EnQueue(t *testing.T) {
	capacity := 20
	q := NewCircularQueue(capacity)
	q.Print()
	if !q.IsEmpty() {
		t.Fatalf("queue IsEmpty error")
	}
	for i := 0; i < capacity-1; i++ {
		if !q.EnQueue(i) {
			t.Fatal("EnQueue error")
		}
	}
	if q.EnQueue(capacity) {
		t.Fatal("EnQueue error")
	}
	if q.Size() != capacity-1 {
		t.Fatalf("get Queue size error")
	}
	if !q.IsFull() {
		t.Fatalf("queue isFull error")
	}
	if q.IsEmpty() {
		t.Fatalf("queue IsEmpty error")
	}
	q.Print()
}

func TestCircularQueue_DeQueue(t *testing.T) {
	capacity := 20
	q := NewCircularQueue(capacity)
	q.Print()
	t.Log("EnQueue......")
	for i := 0; i < capacity-1; i++ {
		if !q.EnQueue(i) {
			t.Fatal("EnQueue error")
		}
	}
	q.Print()
	t.Log("DeQueue......")
	for i := 0; i < capacity-1; i++ {
		v := q.DeQueue()
		t.Logf("[elem/size] is: [%v/%v]", v, q.Size())
	}
	if q.DeQueue() != nil {
		t.Fatalf("queue DeQueue error")
	}
	if !q.IsEmpty() {
		t.Fatalf("queue IsEmpty error")
	}
	if q.IsFull() {
		t.Fatalf("queue IsFull error")
	}
	q.Print()
	t.Log("EnQueue......")
	for i := 0; i < capacity-1; i++ {
		if !q.EnQueue(i) {
			t.Fatal("EnQueue error")
		}
	}
	q.Print()
}

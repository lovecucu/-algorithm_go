package linked_list

import "testing"

func TestNewNode(t *testing.T) {
	newNode := NewNode(10)
	if newNode.next != nil || newNode.value != 10 {
		t.Error(`TestNewNode failed`)
	}
}

func TestNewSingleLinkedList(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	if singleLinkedList.head != nil || singleLinkedList.length != 0 {
		t.Error(`TestNewSingleLinkedList failed`)
	}
}

func TestInsertToHead(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToHead(1)
	if singleLinkedList.length != 1 || singleLinkedList.head == nil || singleLinkedList.head.value != 1 {
		t.Error(`TestInsertToHead failed`)
	}
	singleLinkedList.InsertToHead(2)
	if singleLinkedList.length != 2 || singleLinkedList.head == nil || singleLinkedList.head.value != 2 {
		t.Error(`TestInsertToHead failed`)
	}
}

func TestInsertToTail(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToTail(1)
	if singleLinkedList.length != 1 || singleLinkedList.head == nil || singleLinkedList.head.value != 1 {
		t.Error(`TestInsertToTail failed`)
	}
	singleLinkedList.InsertToTail(2)
	if singleLinkedList.length != 2 || singleLinkedList.head == nil || singleLinkedList.head.value != 1 {
		t.Error(`TestInsertToTail failed`, singleLinkedList.length, singleLinkedList.head.value)
	}
}

func TestGetTailNode(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToHead(1)
	singleLinkedList.InsertToHead(2)
	tailNode := singleLinkedList.getTailNode()
	if tailNode == nil || tailNode.value != 1 {
		t.Error(`TestGetTailNode failed`)
	}
}

func TestFindByIndex(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToHead(1)
	singleLinkedList.InsertToHead(2)
	targetNode := singleLinkedList.FindByIndex(0)
	if targetNode == nil || targetNode.value != 2 {
		t.Error(`TestFindByIndex failed`)
	}
	targetNode = singleLinkedList.FindByIndex(1)
	if targetNode == nil || targetNode.value != 1 {
		t.Error(`TestFindByIndex failed`)
	}
}

func TestInsertBefore(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToTail(1)
	singleLinkedList.InsertToTail(2)
	preNode := singleLinkedList.FindByIndex(1)
	singleLinkedList.InsertBefore(preNode, 3)
	if singleLinkedList.FindByIndex(1).value != 3 {
		t.Error(`TestInsertBefore failed`)
	}
	noExistNode := NewNode(10)
	if singleLinkedList.InsertBefore(noExistNode, 3) {
		t.Error(`TestInsertBefore failed`)
	}
}

func TestInsertAfter(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToTail(1)
	singleLinkedList.InsertToTail(2)
	preNode := singleLinkedList.FindByIndex(1)
	singleLinkedList.InsertAfter(preNode, 3)
	if singleLinkedList.FindByIndex(2).value != 3 {
		t.Error(`TestInsertAfter failed`)
	}
}

func TestLen(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToTail(1)
	singleLinkedList.InsertToTail(2)
	if singleLinkedList.Len() != 2 {
		t.Error(`TestLen failed`)
	}
}

func TestDeleteNode(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToTail(1)
	singleLinkedList.InsertToTail(2)
	curNode := singleLinkedList.FindByIndex(1)
	singleLinkedList.DeleteNode(curNode)
	if singleLinkedList.length != 1 || singleLinkedList.head.value != 1 {
		t.Error(`TestDeleteNode failed`)
	}
	singleLinkedList.DeleteNode(singleLinkedList.head)
	if singleLinkedList.length != 0 || singleLinkedList.head != nil {
		t.Error(`TestDeleteNode failed`)
	}
}

func TestString(t *testing.T) {
	singleLinkedList := NewSingleLinkedList()
	singleLinkedList.InsertToTail(1)
	singleLinkedList.InsertToTail(2)

	want := "SingleLinkedList: size = 2\n[1,2]"
	if want != singleLinkedList.String() {
		t.Error(`TestString failed`)
	}
}

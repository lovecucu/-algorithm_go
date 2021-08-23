package linked_list

import (
	"fmt"
	"testing"
)

func TestNcListReverseList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	reverse := ReverseList(head)
	if reverse.Val != 3 || reverse.Next.Val != 2 || reverse.Next.Next.Val != 1 {
		t.Error(`TestNcListReverseList failed`)
	}
}

func TestNcListHasCycle(t *testing.T) {
	node2 := &ListNode{Val: 2}
	head := &ListNode{Val: 1}
	head.Next = node2
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = node2
	if !hasCycle(head) {
		t.Error(`TestNcListHasCycle failed`)
	}
}

func TestNcListMerge(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node1.Next = &ListNode{Val: 2}
	node1.Next.Next = &ListNode{Val: 4}

	node2 := &ListNode{Val: 3}
	node2.Next = &ListNode{Val: 5}

	merge := Merge(node1, node2)
	target := "12345"
	real := ""
	for merge != nil {
		real += fmt.Sprint(merge.Val)
		merge = merge.Next
	}

	if real != target {
		t.Error(`TestNcListMerge failed`, real)
	}
}

func TestNcListSortInList(t *testing.T) {
	node1 := &ListNode{Val: 3}
	node1.Next = &ListNode{Val: 2}
	node1.Next.Next = &ListNode{Val: 4}
	node1.Next.Next.Next = &ListNode{Val: 1}

	sort := sortInList(node1)
	target := "1234"
	real := ""
	for sort != nil {
		real += fmt.Sprint(sort.Val)
		sort = sort.Next
	}

	if real != target {
		t.Error(`TestNcListMerge failed`, real)
	}
}

func TestNcListIsPail(t *testing.T) {
	node := &ListNode{Val: 1}
	if !isPail(node) {
		t.Error(`TestNcListIsPail failed`)
	}
	node = &ListNode{Val: 2}
	node.Next = &ListNode{Val: 1}
	if isPail(node) {
		t.Error(`TestNcListIsPail failed`)
	}

	node = &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 2}
	node.Next.Next.Next = &ListNode{Val: 1}
	if !isPail(node) {
		t.Error(`TestNcListIsPail failed`)
	}
}

func TestNcListFindFirstCommonNode(t *testing.T) {
	node6 := &ListNode{Val: 6}
	node5 := &ListNode{Val: 5}
	node5.Next = node6
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = node5

	head2 := &ListNode{Val: 4}
	head2.Next = &ListNode{Val: 5}
	head2.Next.Next = node5

	target := "56"
	firstNode := FindFirstCommonNode(head1, head2)
	real := getString(firstNode)
	if target != real {
		t.Error(`TestNcListFindFirstCommonNode failed`)
	}

	head3 := &ListNode{Val: 1}
	head4 := &ListNode{Val: 2}
	head4.Next = &ListNode{Val: 3}
	if FindFirstCommonNode(head3, head4) != nil {
		t.Error(`TestNcListFindFirstCommonNode failed`)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 1}
	head1.Next.Next = &ListNode{Val: 2}
	delete := deleteDuplicates(head1)
	if getString(delete) != "12" {
		t.Error(`TestDeleteDuplicates failed`, getString(delete))
	}

	head1 = &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 1}
	head1.Next.Next = &ListNode{Val: 2}
	head1.Next.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next.Next = &ListNode{Val: 3}

	delete = deleteDuplicates(head1)
	if getString(delete) != "123" {
		t.Error(`TestDeleteDuplicates failed`, getString(delete))
	}
}

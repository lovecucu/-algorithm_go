package nowcoder

import (
	"fmt"
	"sort"
	"testing"
)

func TestLRU(t *testing.T) {
	operators := [][]int{
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 3},
		{2, 1},
		{1, 4, 4},
		{2, 2},
		{1, 3, 5},
		{2, 3},
	}
	result := LRU(operators, 3)
	target := "[1 -1 5]"
	if fmt.Sprint(result) != target {
		t.Error(`TestLRU failed`, fmt.Sprint(result))
	}

	operators = [][]int{{1, 1, 1}, {1, 2, 2}, {2, 1}, {1, 3, 3}, {2, 2}, {1, 4, 4}, {2, 1}, {2, 3}, {2, 4}}
	result = LRU(operators, 2)
	target = "[1 -1 -1 3 4]"
	if fmt.Sprint(result) != target {
		t.Error(`TestLRU failed`, fmt.Sprint(result))
	}
}

func TestThreeOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right

	target := "[[1 2 3] [2 1 3] [2 3 1]]"
	if fmt.Sprint(threeOrders(root)) != target {
		t.Error(`TestThreeOrder failed`)
	}

	left.Left = &TreeNode{Val: 4}
	left.Right = &TreeNode{Val: 5}
	left.Right.Left = &TreeNode{Val: 7}
	left.Right.Right = &TreeNode{Val: 8}

	right.Right = &TreeNode{Val: 6}
	target = "[[1 2 4 5 7 8 3 6] [4 2 7 5 8 1 3 6] [4 7 8 5 2 6 3 1]]"
	if fmt.Sprint(threeOrders(root)) != target {
		t.Error(`TestThreeOrder failed`)
	}
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right
	left.Left = &TreeNode{Val: 4}
	left.Right = &TreeNode{Val: 5}
	left.Right.Left = &TreeNode{Val: 7}
	left.Right.Right = &TreeNode{Val: 8}
	right.Right = &TreeNode{Val: 6}

	target := "[[1] [2 3] [4 5 6] [7 8]]"
	if fmt.Sprint(levelOrder(root)) != target {
		t.Error(`TestLevelOrder failed`, fmt.Sprint(levelOrder(root)))
	}
}

func TestGetLeastNumbersQuickSort(t *testing.T) {
	least4 := GetLeastNumbers_QuickSort([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersQuickSort failed`)
	}

	least3 := GetLeastNumbers_QuickSort([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersQuickSort failed`)
	}
}

func TestGetLeastNumbersSelfSort(t *testing.T) {
	least4 := GetLeastNumbers_SelfSort([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersSelfSort failed`)
	}

	least3 := GetLeastNumbers_SelfSort([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersSelfSort failed`)
	}
}

func TestGetLeastNumbersHeap(t *testing.T) {
	least4 := GetLeastNumbers_Heap([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersHeap failed`, fmt.Sprint(least4))
	}

	least3 := GetLeastNumbers_Heap([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersHeap failed`, fmt.Sprint(least3))
	}
}

func TestGetLeastNumbersQuickSearch(t *testing.T) {
	least4 := GetLeastNumbers_QuickSearch([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersQuickSearch failed`, fmt.Sprint(least4))
	}

	least3 := GetLeastNumbers_QuickSearch([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersQuickSearch failed`, fmt.Sprint(least3))
	}
}

func TestGetLeastNumbersCounterSort(t *testing.T) {
	least4 := GetLeastNumbers_CounterSort([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersCounterSort failed`, fmt.Sprint(least4))
	}

	least3 := GetLeastNumbers_CounterSort([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersCounterSort failed`, fmt.Sprint(least3))
	}
}

func TestFindKthQuickSearch(t *testing.T) {
	if findKth([]int{1, 3, 5, 2, 2}, 5, 3) != 2 {
		t.Error(`TestFindKthQuickSearch failed`, findKth([]int{1, 3, 5, 2, 2}, 5, 3))
	}

	if findKth([]int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2}, 12, 3) != 9 {
		t.Error(`TestFindKthQuickSearch failed`, findKth([]int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2}, 12, 3))
	}

	if findKth([]int{1, 3, 5, 2, 2}, 5, 3) != 2 {
		t.Error(`TestFindKthQuickSearch failed`, findKth([]int{1, 3, 5, 2, 2}, 5, 3))
	}

	res := findKth([]int{1332802, 1177178, 1514891, 871248, 753214, 123866, 1615405, 328656, 1540395, 968891, 1884022, 252932, 1034406, 1455178, 821713, 486232, 860175, 1896237, 852300, 566715, 1285209, 1845742, 883142, 259266, 520911, 1844960, 218188, 1528217, 332380, 261485, 1111670, 16920, 1249664, 1199799, 1959818, 1546744, 1904944, 51047, 1176397, 190970, 48715, 349690, 673887, 1648782, 1010556, 1165786, 937247, 986578, 798663}, 49, 24)
	if res != 986578 {
		t.Error(`TestFindKthQuickSearch failed`, res)
	}

	res = findKth([]int{1793111, 1704885, 1533399, 1841885, 1106030, 1075047, 683720, 1775812, 384614, 1748120, 465909, 1464784, 371144, 1169996, 1547285, 1815434, 371770, 1534437, 1361914, 1908006, 1642892, 940052, 1214020, 1293092, 1974448, 161119, 1323434, 556599, 1373519, 1464940, 279574, 262433, 335617, 109745, 288046, 1489090, 1117600, 1256093, 846346, 752014, 786901, 163280, 110109, 1518282, 1233229, 314395, 369349, 1417147, 1865047, 73156, 798853, 130502, 574144, 988607, 265340, 1552401, 1727426, 1346779, 530528, 281836, 311013, 1646911, 839746, 1411724, 1353713, 1885642, 1218958, 640981, 1371397, 1901432, 82962, 1432921, 203321, 1595713, 321526, 948973, 1236208, 1363959, 934899, 896793, 1508384, 548839, 1814294, 22183, 725125, 1952668, 759735, 1834610, 12072, 950119, 837758, 1318463, 581829, 776083, 1773795, 9111, 166708, 1983888, 436686, 992239, 1494229, 7269, 91218, 1582448, 548987, 1041088, 1557663, 1061803, 181357, 69709, 1990660, 614682, 9689, 1458544, 877325, 863806, 930818, 1818782, 1005295, 652631, 1908046, 1108124, 1820235, 396014, 921750, 194445, 329391, 271492, 1231544, 1713579, 1211384, 483379, 921719, 144907, 768520, 1923510, 1172025, 1142186, 140381, 1221299, 314247, 26366, 429496, 599445, 57461, 1150445, 1885079, 1508820, 767293, 968731, 1498538, 1484970, 696383, 1631789, 191440, 1378019, 1252811, 72983, 7956, 1743383, 1589067, 389357, 1987464, 805141, 1503832, 631207, 1369959, 1128676, 1107546, 784351, 1807031, 1626747, 1080788, 867275, 1226917, 1791668, 413854, 1211411, 712406, 1002491, 1561140, 42784, 816874, 382335, 369016, 1113067, 62854, 936132, 62921, 1509838}, 190, 134)
	if res != 556599 {
		t.Error(`TestFindKthQuickSearch failed`, res)
	}
}

func TestFindKthHeapSort(t *testing.T) {
	if findKth_HeapSort([]int{1, 3, 5, 2, 2}, 5, 3) != 2 {
		t.Error(`TestFindKthHeapSort failed`, findKth([]int{1, 3, 5, 2, 2}, 5, 3))
	}

	if findKth_HeapSort([]int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2}, 12, 3) != 9 {
		t.Error(`TestFindKthHeapSort failed`, findKth([]int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2}, 12, 3))
	}

	if findKth_HeapSort([]int{1, 3, 5, 2, 2}, 5, 3) != 2 {
		t.Error(`TestFindKthHeapSort failed`, findKth([]int{1, 3, 5, 2, 2}, 5, 3))
	}

	res := findKth_HeapSort([]int{1332802, 1177178, 1514891, 871248, 753214, 123866, 1615405, 328656, 1540395, 968891, 1884022, 252932, 1034406, 1455178, 821713, 486232, 860175, 1896237, 852300, 566715, 1285209, 1845742, 883142, 259266, 520911, 1844960, 218188, 1528217, 332380, 261485, 1111670, 16920, 1249664, 1199799, 1959818, 1546744, 1904944, 51047, 1176397, 190970, 48715, 349690, 673887, 1648782, 1010556, 1165786, 937247, 986578, 798663}, 49, 24)
	if res != 986578 {
		t.Error(`TestFindKthHeapSort failed`, res)
	}

	res = findKth_HeapSort([]int{1793111, 1704885, 1533399, 1841885, 1106030, 1075047, 683720, 1775812, 384614, 1748120, 465909, 1464784, 371144, 1169996, 1547285, 1815434, 371770, 1534437, 1361914, 1908006, 1642892, 940052, 1214020, 1293092, 1974448, 161119, 1323434, 556599, 1373519, 1464940, 279574, 262433, 335617, 109745, 288046, 1489090, 1117600, 1256093, 846346, 752014, 786901, 163280, 110109, 1518282, 1233229, 314395, 369349, 1417147, 1865047, 73156, 798853, 130502, 574144, 988607, 265340, 1552401, 1727426, 1346779, 530528, 281836, 311013, 1646911, 839746, 1411724, 1353713, 1885642, 1218958, 640981, 1371397, 1901432, 82962, 1432921, 203321, 1595713, 321526, 948973, 1236208, 1363959, 934899, 896793, 1508384, 548839, 1814294, 22183, 725125, 1952668, 759735, 1834610, 12072, 950119, 837758, 1318463, 581829, 776083, 1773795, 9111, 166708, 1983888, 436686, 992239, 1494229, 7269, 91218, 1582448, 548987, 1041088, 1557663, 1061803, 181357, 69709, 1990660, 614682, 9689, 1458544, 877325, 863806, 930818, 1818782, 1005295, 652631, 1908046, 1108124, 1820235, 396014, 921750, 194445, 329391, 271492, 1231544, 1713579, 1211384, 483379, 921719, 144907, 768520, 1923510, 1172025, 1142186, 140381, 1221299, 314247, 26366, 429496, 599445, 57461, 1150445, 1885079, 1508820, 767293, 968731, 1498538, 1484970, 696383, 1631789, 191440, 1378019, 1252811, 72983, 7956, 1743383, 1589067, 389357, 1987464, 805141, 1503832, 631207, 1369959, 1128676, 1107546, 784351, 1807031, 1626747, 1080788, 867275, 1226917, 1791668, 413854, 1211411, 712406, 1002491, 1561140, 42784, 816874, 382335, 369016, 1113067, 62854, 936132, 62921, 1509838}, 190, 134)
	if res != 556599 {
		t.Error(`TestFindKthHeapSort failed`, res)
	}
}

func TestTwoSum(t *testing.T) {
	target := "[1 2]"
	if fmt.Sprint(twoSum([]int{20, 70, 110, 150}, 90)) != target {
		t.Error(`testTwoSum failed`, fmt.Sprint(twoSum([]int{20, 70, 110, 150}, 90)))
	}

	target = "[2 3]"
	if fmt.Sprint(twoSum([]int{3, 2, 4}, 6)) != target {
		t.Error(`testTwoSum failed`, fmt.Sprint(twoSum([]int{3, 2, 4}, 6)))
	}
}

func TestMergeTwoListNode(t *testing.T) {
	phead1 := &ListNode{Val: 1}
	phead1.Next = &ListNode{Val: 3}
	phead1.Next.Next = &ListNode{Val: 5}

	phead2 := &ListNode{Val: 2}
	phead2.Next = &ListNode{Val: 4}
	phead2.Next.Next = &ListNode{Val: 6}

	phead := Merge(phead1, phead2)
	real := []int{}
	target := "[1 2 3 4 5 6]"
	for cur := phead; cur != nil; cur = cur.Next {
		real = append(real, cur.Val)
	}
	if fmt.Sprint(real) != target {
		t.Error(`TestMergeTwoListNode failed`, real)
	}
}

func TestTwoStackQueue(t *testing.T) {
	Push(1)
	Push(2)
	ret := []int{}
	ret = append(ret, Pop())
	ret = append(ret, Pop())

	target := "[1 2]"
	if fmt.Sprint(ret) != target {
		t.Error(`TestTwoStackQueue failed`, fmt.Sprint(ret))
	}
}

func TestJumpFloor(t *testing.T) {
	jumpCache = nil
	if jumpFloor(2) != 2 {
		t.Error(`TestJumpFloor failed`, jumpFloor(2))
	}

	jumpCache = nil
	if jumpFloor(7) != 21 {
		t.Error(`TestJumpFloor failed`, jumpFloor(7))
	}
}

func TestReverseKGroup(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	revHead := reverseKGroup(head, 2)
	real := []int{}
	for cur := revHead; cur != nil; cur = cur.Next {
		real = append(real, cur.Val)
	}
	target := "[2 1 4 3 5]"
	if fmt.Sprint(real) != target {
		t.Error(`TestReverseKGroup failed`, fmt.Sprint(real))
	}
}

func TestHasCycle(t *testing.T) {
	head := &ListNode{Val: 1}
	next := &ListNode{Val: 2}
	head.Next = next
	next.Next = &ListNode{Val: 3}
	next.Next.Next = &ListNode{Val: 4}
	next.Next.Next.Next = next

	if !hasCycle(head) {
		t.Error(`TestHasCycle failed`)
	}

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	third := &ListNode{Val: 3}
	head.Next.Next = third
	third.Next = &ListNode{Val: 4}
	third.Next.Next = &ListNode{Val: 5}
	third.Next.Next.Next = &ListNode{Val: 6}
	third.Next.Next.Next.Next = third

	if !hasCycle(head) {
		t.Error(`TestHasCycle failed`)
	}
}

func TestCycleEntryNode(t *testing.T) {
	head := &ListNode{Val: 1}
	next := &ListNode{Val: 2}
	head.Next = next
	next.Next = &ListNode{Val: 3}
	next.Next.Next = &ListNode{Val: 4}
	next.Next.Next.Next = next

	if EntryNodeOfLoop(head) != next {
		t.Error(`TestCycleEntryNode failed`)
	}

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	third := &ListNode{Val: 3}
	head.Next.Next = third
	third.Next = &ListNode{Val: 4}
	third.Next.Next = &ListNode{Val: 5}
	third.Next.Next.Next = &ListNode{Val: 6}
	third.Next.Next.Next.Next = third

	if EntryNodeOfLoop(head) != third {
		t.Error(`TestCycleEntryNode failed`)
	}

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}

	if EntryNodeOfLoop(head) != nil {
		t.Error(`TestCycleEntryNode failed`)
	}
}

func TestIsBracketsValid(t *testing.T) {
	if isBracketsValid("((") || isBracketsValid("[") || isBracketsValid("{][}") {
		t.Error(`TestIsBracketsValid failed`)
	}

	if !isBracketsValid("[]{}") || !isBracketsValid("{[]}") {
		t.Error(`TestIsBracketsValid failed`)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}

	if SprintNode(removeNthFromEnd(head, 1)) != "[1]" {
		t.Error(`TestRemoveNthFromEnd failed`)
	}

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	if SprintNode(removeNthFromEnd(head, 2)) != "[2]" {
		t.Error(`TestRemoveNthFromEnd failed`)
	}
}

func TestMaxSumOfSubArray(t *testing.T) {
	if maxsumofSubarray([]int{1, -2, 3, 5, -2, 6, -1}) != 12 {
		t.Error(`TestMaxSumOfSubArray failed`)
	}
}

func TestMaxLength(t *testing.T) {
	if maxLength([]int{2, 3, 4, 5}) != 4 {
		t.Error(`TestMaxLength failed`)
	}

	if maxLength([]int{2, 2, 3, 4, 3}) != 3 {
		t.Error(`TestMaxLength failed`)
	}

	if maxLength([]int{1, 2, 3, 1, 2, 3, 2, 2}) != 3 {
		t.Error(`TestMaxLength failed`)
	}

	if maxLength([]int{2, 2, 3, 4, 8, 99, 3}) != 5 {
		t.Error(`TestMaxLength failed`)
	}
}

func TestMergeSortedArray(t *testing.T) {
	real := merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	target := "[1 2 3 4 5 6]"
	if fmt.Sprint(real) != target {
		t.Error(`TestMergeSortedArray failed`)
	}

	real = merge([]int{1}, 1, []int{2}, 1)
	target = "[1 2]"
	if fmt.Sprint(real) != target {
		t.Error(`TestMergeSortedArray failed`)
	}
}

func TestBigAdd(t *testing.T) {
	if solve("1", "99") != "100" {
		t.Error(`TestBigAdd failed`)
	}
}

func TestZPrint(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	if fmt.Sprint(Print(root)) != "[[1] [3 2] [4 5]]" {
		t.Error(`TestZPrint failed`)
	}

	root = &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 11}
	if fmt.Sprint(Print(root)) != "[[8] [10 6] [5 7 9 11]]" {
		t.Error(`TestZPrint failed`)
	}
}

func TestLCS(t *testing.T) {
	if LCS("1AB2345CD", "12345EF") != "2345" {
		t.Error(`TestLCS failed`)
	}
}

func TestFindFirstCommonNode(t *testing.T) {
	if FindFirstCommonNode(nil, nil) != nil {
		t.Error(`TestFindFirstCommonNode failed`)
	}

	head1 := &ListNode{Val: 1}

	head2 := &ListNode{Val: 2}
	head2.Next = &ListNode{Val: 3}

	if FindFirstCommonNode(head1, head2) != nil {
		t.Error(`TestFindFirstCommonNode failed`)
	}

	headCommon := &ListNode{Val: 6}
	headCommon.Next = &ListNode{Val: 7}

	head1 = &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = headCommon

	head2 = &ListNode{Val: 4}
	head2.Next = &ListNode{Val: 5}
	head2.Next.Next = headCommon

	if SprintNode(FindFirstCommonNode(head1, head2)) != "[6 7]" {
		t.Error(`TestFindFirstCommonNode failed`)
	}
}

func TestAddInList(t *testing.T) {
	head1 := &ListNode{Val: 9}
	head1.Next = &ListNode{Val: 3}
	head1.Next.Next = &ListNode{Val: 7}

	head2 := &ListNode{Val: 6}
	head2.Next = &ListNode{Val: 3}

	if SprintNode(addInList(head1, head2)) != "[1 0 0 0]" {
		t.Error(`TestAddInList failed`)
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	if lowestCommonAncestor(root, 4, 5) != 2 {
		t.Error(`TestLowestCommonAncestor failed`)
	}
}

func TestReverseString(t *testing.T) {
	if reverseString("abcd") != "dcba" {
		t.Error(`TestReverseString failed`)
	}
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{}
	if fmt.Sprint(spiralOrder(matrix)) != "[]" {
		t.Error(`TestSpiralOrder failed`)
	}

	matrix = [][]int{
		{7},
		{9},
		{6},
	}
	if fmt.Sprint(spiralOrder(matrix)) != "[7 9 6]" {
		t.Error(`TestSpiralOrder failed`)
	}

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	if fmt.Sprint(spiralOrder(matrix)) != "[1 2 3 6 9 8 7 4 5]" {
		t.Error(`TestSpiralOrder failed`)
	}

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	if fmt.Sprint(spiralOrder(matrix)) != "[1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10]" {
		t.Error(`TestSpiralOrder failed`)
	}
}

func TestFibonacci(t *testing.T) {
	if Fibonacci(3) != 2 {
		t.Error(`TestFibonacci failed`)
	}

	if Fibonacci(6) != 8 {
		t.Error(`TestFibonacci failed`)
	}
}

func TestGetLongestPalindrome(t *testing.T) {
	if getLongestPalindrome("a", 1) != 1 {
		t.Error(`TestGetLongestPalindrome failed`)
	}

	if getLongestPalindrome("aba", 3) != 3 {
		t.Error(`TestGetLongestPalindrome failed`)
	}

	if getLongestPalindrome("abc1234321ab", 12) != 7 {
		t.Error(`TestGetLongestPalindrome failed`)
	}
}

func TestThreeSum(t *testing.T) {
	if fmt.Sprint(threeSum([]int{-2, 0, 1, 1, 2})) != "[[-2 0 2] [-2 1 1]]" {
		t.Error(`TestThreeSum failed`)
	}
	if fmt.Sprint(threeSum([]int{-10, 0, 10, 20, -10, -40})) != "[[-10 -10 20] [-10 0 10]]" {
		t.Error(`TestThreeSum failed`)
	}

	if fmt.Sprint(threeSum([]int{-2, 0, 0, 0, 2, 1})) != "[[-2 0 2] [0 0 0]]" {
		t.Error(`TestThreeSum failed`)
	}
}

func TestReConstructBinaryTree(t *testing.T) {
	if SprintTreeNode(reConstructBinaryTree([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 2, 4, 1, 6, 5, 7})) != "[1 2 3 4 5 6 7]" {
		t.Error(`TestReConstructBinaryTree failed`)
	}

	if SprintTreeNode(reConstructBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})) != "[1 2 4 7 3 5 6 8]" {
		t.Error(`TestReConstructBinaryTree failed`)
	}
}

func TestLIS(t *testing.T) {
	real := LIS([]int{2, 1, 5, 3, 6, 4, 8, 9, 7})
	if fmt.Sprint(real) != "[1 3 4 8 9]" {
		t.Error(`TestLIS failed`)
	}

	if fmt.Sprint(LIS([]int{1, 2, 8, 6, 4})) != "[1 2 4]" {
		t.Error(`TestLIS failed`)
	}

	if fmt.Sprint(LIS([]int{6, 7, 8, 5, 6, 8, 12})) != "[5 6 8 12]" {
		t.Error(`TestLIS failed`)
	}
}

func TestLISBinary(t *testing.T) {
	real := LISBinary([]int{2, 1, 5, 3, 6, 4, 8, 9, 7})
	if fmt.Sprint(real) != "[1 3 4 8 9]" {
		t.Error(`TestLISBinary failed`)
	}

	if fmt.Sprint(LISBinary([]int{1, 2, 8, 6, 4})) != "[1 2 4]" {
		t.Error(`TestLISBinary failed`)
	}

	if fmt.Sprint(LISBinary([]int{6, 7, 8, 5, 6, 8, 12})) != "[5 6 8 12]" {
		t.Error(`TestLISBinary failed`)
	}
}

func TestSqrt(t *testing.T) {
	if sqrt(0) != 0 || sqrt(2) != 1 || sqrt(1) != 1 || sqrt(10) != 3 || sqrt(9) != 3 {
		t.Error(`TestSqrt failed`)
	}
}

func TestSearch(t *testing.T) {
	if search([]int{6, 8, 10, 0, 2, 4}, 0) != 3 || search([]int{6, 8, 10, 0, 2, 4}, 2) != 4 || search([]int{6, 8, 10, 0, 2, 4}, 8) != 1 {
		t.Error(`TestSearch failed`)
	}
	if search([]int{6, 8, 10, 0, 2, 4}, 10) != 2 || search([]int{6, 8, 10, 0, 2, 4}, 3) != -1 || search([]int{2}, 1) != -1 {
		t.Error(`TestSearch failed`)
	}
}

func TestStackMin(t *testing.T) {
	// ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"]
	output := []int{}
	StackPush(-1)
	StackPush(2)
	output = append(output, StackMin())
	output = append(output, StackTop())
	StackPop()
	StackPush(1)
	output = append(output, StackTop())
	output = append(output, StackMin())
	if fmt.Sprint(output) != "[-1 2 1 -1]" {
		t.Error(`TestStackMin failed`)
	}
}

func TestMergeKListNode(t *testing.T) {
	phead1 := &ListNode{Val: 1}
	phead1.Next = &ListNode{Val: 3}
	phead1.Next.Next = &ListNode{Val: 5}

	phead2 := &ListNode{Val: 2}
	phead2.Next = &ListNode{Val: 4}
	phead2.Next.Next = &ListNode{Val: 6}

	phead := mergeKLists([]*ListNode{phead1, phead2})
	if fmt.Sprint(SprintNode(phead)) != "[1 2 3 4 5 6]" {
		t.Error(`TestMergeTwoListNode failed`, SprintNode(phead))
	}

	phead1 = &ListNode{Val: 1}
	phead1.Next = &ListNode{Val: 3}
	phead1.Next.Next = &ListNode{Val: 5}

	phead2 = &ListNode{Val: 2}
	phead2.Next = &ListNode{Val: 4}
	phead2.Next.Next = &ListNode{Val: 6}

	phead3 := &ListNode{Val: 7}
	phead3.Next = &ListNode{Val: 8}

	phead = mergeKLists([]*ListNode{phead1, phead2, phead3})
	if fmt.Sprint(SprintNode(phead)) != "[1 2 3 4 5 6 7 8]" {
		t.Error(`TestMergeTwoListNode failed`, SprintNode(phead))
	}
}

func TestPermutation(t *testing.T) {

	real := Permutation("ab")
	sort.Strings(real)
	if fmt.Sprint(real) != "[ab ba]" {
		t.Error(`TestPermutation failed`)
	}

	real = Permutation("abc")
	sort.Strings(real)
	if fmt.Sprint(real) != "[abc acb bac bca cab cba]" {
		t.Error(`TestPermutation failed`)
	}
}

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right

	if maxDepth(root) != 2 {
		t.Error(`TestMaxDepth failed`)
	}

	left.Left = &TreeNode{Val: 4}
	left.Right = &TreeNode{Val: 5}
	left.Right.Left = &TreeNode{Val: 7}
	left.Right.Right = &TreeNode{Val: 8}

	left.Right.Left.Left = &TreeNode{Val: 9}

	if maxDepth(root) != 5 {
		t.Error(`TestMaxDepth failed`)
	}
}

func TestJudge(t *testing.T) {
	if judge("") || judge("ab") || judge("abc") {
		t.Error(`TestJudge failed`)
	}

	if !judge("a") || !judge("aba") || !judge("abba") {
		t.Error(`TestJudge failed`)
	}
}

func TestIsBalanced(t *testing.T) {
	var root *TreeNode
	if !IsBalanced_Solution(root) {
		t.Error(`TestIsBalanced failed`)
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	if !IsBalanced_Solution(root) {
		t.Error(`TestIsBalanced failed`)
	}

	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 5}
	if IsBalanced_Solution(root) {
		t.Error(`TestIsBalanced failed`)
	}
}

func TestMaxProfit(t *testing.T) {
	if maxProfit([]int{1, 4, 2}) != 3 || maxProfit([]int{2, 4, 1}) != 2 || maxProfit([]int{3, 1, 5, 2}) != 4 {
		t.Error(`TestMaxProfit failed`)
	}

	if maxProfit([]int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9}) != 8 {
		t.Error(`TestMaxProfit failed`)
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

func TestMoreThanHalfNum_Solution(t *testing.T) {
	if MoreThanHalfNum_Solution([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}) != 2 {
		t.Error(`TestMoreThanHalfNum_Solution failed`)
	}

	if MoreThanHalfNum_Solution([]int{3, 3, 3, 3, 2, 2, 2}) != 3 {
		t.Error(`TestMoreThanHalfNum_Solution failed`)
	}

	if MoreThanHalfNum_Solution([]int{1}) != 1 {
		t.Error(`TestMoreThanHalfNum_Solution failed`)
	}
}

func TestSolveExpression(t *testing.T) {
	if solveExpression("(2*(3-4))*5") != -10 {
		t.Error(`TestSolveExpression failed`)
	}

	if solveExpression("3+2*3*4-1") != 26 {
		t.Error(`TestSolveExpression failed`)
	}

	if solveExpression("3+-2*3*4-1") != -22 {
		t.Error(`TestSolveExpression failed`)
	}

	if solveExpression("100-1") != 99 {
		t.Error(`TestSolveExpression failed`)
	}

	if solveExpression("((10+2)*10-(100-(10+20*10-(2*3)))*10*1*2)-2") != 2198 {
		t.Error(`TestSolveExpression failed`)
	}

	if solveExpression("1-2-3") != -4 {
		t.Error(`TestSolveExpression failed`)
	}

	if solveExpression("100+100") != 200 {
		t.Error(`TestSolveExpression failed`)
	}
}

func TestSolveRightView(t *testing.T) {
	if fmt.Sprint(solveRightView([]int{1, 2, 4, 5, 3}, []int{4, 2, 5, 1, 3})) != "[1 3 5]" {
		t.Error(`TestSolveRightView failed`, solveRightView([]int{1, 2, 4, 5, 3}, []int{4, 2, 5, 1, 3}))
	}
}

func TestMaxProfitTwice(t *testing.T) {
	if maxProfitTwice([]int{8, 9, 3, 5, 1, 3}) != 4 || maxProfitTwice([]int{}) != 0 {
		t.Error(`TestMaxProfitTwice failed`)
	}
}

func TestMaxProfitInfinite(t *testing.T) {
	if maxProfitInfinite([]int{5, 4, 3, 2, 1}) != 0 || maxProfitInfinite([]int{1, 2, 3, 4, 5}) != 4 {
		t.Error(`TestMaxProfitInfinite failed`)
	}
}

func TestTopKStrings(t *testing.T) {
	if fmt.Sprint(topKstrings([]string{"123", "123", "231", "32"}, 2)) != "[[123 2] [231 1]]" {
		t.Error(`TestTopKStrings failed`)
	}
	if fmt.Sprint(topKstrings([]string{"a", "b", "c", "b"}, 2)) != "[[b 2] [a 1]]" {
		t.Error(`TestTopKStrings failed`)
	}
	if fmt.Sprint(topKstrings([]string{"abcd", "abcd", "abcd", "pwb2", "abcd", "pwb2", "p12"}, 3)) != "[[abcd 4] [pwb2 2] [p12 1]]" {
		t.Error(`TestTopKStrings failed`)
	}
}

func TestHexConvert(t *testing.T) {
	if hexConvert(2500, 16) != "9C4" {
		t.Error(`TestHexConvert failed`)
	}

	if hexConvert(10, 2) != "1010" {
		t.Error(`TestHexConvert failed`)
	}

	if hexConvert(5, 2) != "101" {
		t.Error(`TestHexConvert failed`)
	}

	if hexConvert(7, 2) != "111" {
		t.Error(`TestHexConvert failed`)
	}

	if hexConvert(-4, 3) != "-11" {
		t.Error(`TestHexConvert failed`)
	}
}

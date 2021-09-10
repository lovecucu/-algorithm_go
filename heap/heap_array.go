package heap

func HeapSort(input []int) {
	inputLen := len(input)
	for i := 0; i < inputLen; i++ {
		minAjust(input[i:]) // 每次排好第i位的数据
	}
}

func minAjust(input []int) {
	inputLen := len(input)
	if inputLen <= 1 {
		return
	}
	for i := inputLen/2 - 1; i >= 0; i-- { // 小根堆，父结点如果大于子结点，则相反
		if (2*i+1 <= inputLen-1) && (input[i] > input[2*i+1]) { // 左子结点
			input[i], input[2*i+1] = input[2*i+1], input[i]
		}
		if (2*i+2 <= inputLen-1) && (input[i] > input[2*i+2]) { // 右子结点
			input[i], input[2*i+2] = input[2*i+2], input[i]
		}
	}
}

package heap

type IntHeap []int

func (h *IntHeap) Push(x int) {
	*h = append(*h, x)
	h.SiftUp(h.Size() - 1)
}

func (h *IntHeap) Pop() int {
	if h.IsEmpty() {
		return -1
	}

	h.Swap(0, h.Size()-1)
	val := (*h)[h.Size()-1]
	*h = (*h)[:h.Size()-1]
	h.SiftDown(0)
	return val
}

func (h *IntHeap) Size() int {
	return len(*h)
}

func (h *IntHeap) IsEmpty() bool {
	return h.Size() == 0
}

// func (h *IntHeap) Less(i, j int) bool {
// 	return (*h)[i] <= (*h)[j]
// }

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// func (h *IntHeap) Top() int {
// 	return (*h)[0]
// }

func (h *IntHeap) Left(i int) int {
	return 2*i + 1
}

func (h *IntHeap) Right(i int) int {
	return 2*i + 2
}

func (h *IntHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

func (h *IntHeap) SiftUp(i int) {
	for true {
		p := h.Parent(i)

		// 当前是根节点或者
		if p < 0 || (*h)[p] >= (*h)[i] {
			break
		}
		h.Swap(p, i)
		i = p
	}
}

func (h *IntHeap) SiftDown(i int) {
	for true {
		l, r, max := h.Left(i), h.Right(i), i
		if l < h.Size() && (*h)[l] > (*h)[max] {
			max = l
		}

		if r < h.Size() && (*h)[r] > (*h)[max] {
			max = r
		}

		if max == i {
			break
		}

		h.Swap(i, max)
		i = max
	}
}

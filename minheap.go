package gosort

type MinHeap struct {
	heap     []int
	heapSize int
}

func NewMinHeap(storage []int) *MinHeap {
	newHeap := new(MinHeap)
	newHeap.heap = storage
	newHeap.buildMinHeap()

	return newHeap
}

func MinHeapSort(input []int) {
	heap := NewMinHeap(input)
	heap.inplaceSort()
}

func (self *MinHeap) Insert(newItem int) {
}

func (self *MinHeap) Min() int {
	return self.heap[0]
}

func (self *MinHeap) ExtractMin() int {
	return 0
}

func (self *MinHeap) buildMinHeap() {
	self.heapSize = len(self.heap)
	for i := self.heapSize / 2; i >= 0; i-- {
		self.minHeapify(i)
	}
}

func (self *MinHeap) minHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i

	if l < self.heapSize && self.heap[l] < self.heap[i] {
		smallest = l
	}
	if r < self.heapSize && self.heap[r] < self.heap[smallest] {
		smallest = r
	}
	if i != smallest {
		self.heap[i], self.heap[smallest] = self.heap[smallest], self.heap[i]
		self.minHeapify(smallest)
	}
}

func (self *MinHeap) checkInvariant() bool {
	for i := 1; i < self.heapSize; i++ {

		// The current node should be greater than or equal to the parent
		invariant := self.heap[i] >= self.heap[parent(i)]

		if !invariant {
			return false
		}
	}

	return true
}

func (self *MinHeap) inplaceSort() {
	for i := self.heapSize - 1; i > 0; i-- {
		self.heap[0], self.heap[i] = self.heap[i], self.heap[0]
		self.heapSize--
		self.minHeapify(0)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

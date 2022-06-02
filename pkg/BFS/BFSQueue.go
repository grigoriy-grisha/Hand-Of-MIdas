package BFS

type Node struct {
	Visited bool
	Value   interface{}
}

type queue struct {
	queue []Node
}

func NewQueue() *queue {
	return &queue{queue: []Node{}}
}

func (BFSQueue *queue) IsNotEmpty() bool {
	if len(BFSQueue.queue) == 0 {
		return false
	}

	return true
}

func (BFSQueue *queue) Push(Node Node) {
	Node.Visited = true
	BFSQueue.queue = append(BFSQueue.queue, Node)
}

func (BFSQueue *queue) Shift() Node {
	firstElement := BFSQueue.queue[0]

	BFSQueue.queue = BFSQueue.queue[1:]

	return firstElement
}

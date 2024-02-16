package adaptive

type Node4[T any] struct {
	partialLen  uint32
	artNodeType uint8
	numChildren uint8
	partial     []byte
	keys        [4]byte
	children    [4]*Node[T]
}

func (n *Node4[T]) getPartialLen() uint32 {
	return n.partialLen
}

func (n *Node4[T]) setPartialLen(partialLen uint32) {
	n.partialLen = partialLen
}

func (n *Node4[T]) getArtNodeType() uint8 {
	return n.artNodeType
}

func (n *Node4[T]) setArtNodeType(ARTNodeType uint8) {
	n.artNodeType = ARTNodeType
}

func (n *Node4[T]) getNumChildren() uint8 {
	return n.numChildren
}

func (n *Node4[T]) setNumChildren(numChildren uint8) {
	n.numChildren = numChildren
}

func (n *Node4[T]) getPartial() []byte {
	return n.partial
}

func (n *Node4[T]) setPartial(partial []byte) {
	n.partial = partial
}

func (n *Node4[T]) isLeaf() bool {
	return false
}

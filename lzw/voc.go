package lzw

type WordTree struct {
	value    byte
	index    uint16
	children []*WordTree
}

type VocWord []byte

type VocEntry struct {
	base   uint16
	suffix byte
}

type Voc struct {
	wordTree  *WordTree
	nextIndex uint16
}

const vocNoIndex = uint16(0xFFFF)
const vocFullSize = 4096
const vocInitialSize = 256

func NewWordTree() (root *WordTree) {
	root = &WordTree{
		value:    byte(0),
		index:    vocNoIndex,
		children: make([]*WordTree, vocInitialSize),
	}

	for i := 0; i < 256; i++ {
		root.AddChild(byte(i), uint16(i))
	}

	return
}

func (tree *WordTree) AddChild(b byte, index uint16) {
	if tree.children == nil {
		tree.children = make([]*WordTree, vocInitialSize)
	}

	tree.children[b] = &WordTree{
		value:    b,
		index:    index,
		children: nil,
	}
}

func (tree *WordTree) Walk(b byte) (child *WordTree, ok bool) {
	child = tree
	if tree.children == nil {
		ok = false
		return
	}
	if tree.children[b] == nil {
		ok = false
		return
	}

	child = tree.children[b]
	ok = true
	return
}

func NewVoc() *Voc {
	return &Voc{
		wordTree:  NewWordTree(),
		nextIndex: 256,
	}
}

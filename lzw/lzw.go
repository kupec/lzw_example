package lzw

import (
	"encoding/binary"
	"io"
)

func Compress(r io.Reader, w io.Writer) (err error) {
	buffer := make([]byte, 1)
	voc := NewVoc()

	tree := voc.wordTree
	for {
		if _, err = r.Read(buffer); err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			err = nil
			if tree.index != vocNoIndex {
				err = binary.Write(w, binary.LittleEndian, tree.index)
			}
			return
		}

		nextByte := buffer[0]

		var ok bool
		tree, ok = tree.Walk(nextByte)
		if ok {
			continue
		}

		tree.AddChild(nextByte, voc.nextIndex)
		voc.nextIndex++

		err = binary.Write(w, binary.LittleEndian, tree.index)

		tree = voc.wordTree
		tree, _ = tree.Walk(nextByte)
	}
}

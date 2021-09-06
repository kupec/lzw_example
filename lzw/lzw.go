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

func Decompress(r io.Reader, w io.Writer) (err error) {
	voc := NewVoc()
	vocMap := make(map[uint16][]byte)
	for i := 0; i < 256; i++ {
		vocMap[uint16(i)] = []byte{byte(i)}
	}

	var prevSequence []byte
	for {
		var index uint16
		if err = binary.Read(r, binary.LittleEndian, &index); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				err = nil
			}
			return
		}

		sequence, ok := vocMap[index]
		if !ok {
			sequence = append(prevSequence, prevSequence[0])
		}

		if _, err = w.Write(sequence); err != nil {
			return
		}

		if prevSequence != nil {
			tree := voc.wordTree
			for _, b := range prevSequence {
				tree, _ = tree.Walk(b)
			}

			nextIndex := voc.nextIndex
			voc.nextIndex++

			tree.AddChild(sequence[0], nextIndex)
			vocMap[nextIndex] = append(prevSequence, sequence[0])
		}
		prevSequence = sequence
	}
}

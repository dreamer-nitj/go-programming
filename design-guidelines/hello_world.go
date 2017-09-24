package design_guidelines

import (
	"fmt"
	"unsafe"
)

func PrintHello() {
	fmt.Println("Hello, playground")
}

// the sizeof returns 24 for this struct, the
// reason is due to added padding. an alignment
// can be 1, 2, 4, 8
/**
Knowing this rule and remembering that:

bool, int8/uint8 take 1 chunk
int16, uint16 - 2 chunks
int32, uint32, float32 - 4 chunks
int64, uint64, float64, pointer - 8 chunks
string - 16 chunks (2 alignments of 8 chunks)
any slice takes 24 chunks (3 alignments of 8 chunks).
So []bool, [][][]string are the same (do not forget to reread the citation I added in the beginning)
array of length n takes n * type it takes of chunks.
 */
type example struct {
	flag bool // size of bool is 1 byte
	// there will be a 7 byte padding here as counter needs to start at 8 byte alignment
	counter int64 // size here is 8 byte
	pi float32 // size here is 4 bytes
}

func PrintSize() {
	var e example
	fmt.Printf("The size is %d\n", unsafe.Sizeof(e))
	fmt.Printf("%+v\n", e)
}

type Vertex struct {
	X int
	Y int
}

func PrintVertex() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
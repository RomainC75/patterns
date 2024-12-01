package bt

type node struct {
	value int
	left  *node
	right *node
}

type btr struct {
	root *node
	len  int
}

func NewBt() *btr {
	return &btr{}
}

package compiler

type BlockType int
type Block struct {
	Type    BlockType
	Content string
	Props   any
}

type BlockHeadingProps struct {
	Level uint
}

const (
	BLK_HEADING = iota
	BLK_PARAGRAPH
	BLK_END

	BLK_UNKNOW
)

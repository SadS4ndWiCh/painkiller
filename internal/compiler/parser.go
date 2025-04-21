package compiler

type Parser struct {
	lexer *Lexer

	cur Token
	pek Token
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer, pek: lexer.Next()}
}

func (p *Parser) nextToken() {
	p.cur = p.pek
	p.pek = p.lexer.Next()
}

func (p *Parser) nextBlock() Block {
	p.nextToken()

	switch p.cur.Type {
	case TKN_HASH:
		level := uint(1)
		for p.pek.Type == TKN_HASH {
			level += 1
			p.nextToken()
		}

		p.nextToken()
		if p.cur.Type != TKN_STRING {
			return Block{BLK_UNKNOW, p.cur.Literal, nil}
		}

		return Block{BLK_HEADING, p.cur.Literal, BlockHeadingProps{Level: level}}

	case TKN_STRING:
		return Block{BLK_PARAGRAPH, p.cur.Literal, nil}

	case TKN_EOF:
		return Block{BLK_END, p.cur.Literal, nil}

	default:
		return Block{BLK_UNKNOW, p.cur.Literal, nil}
	}
}

func (p *Parser) Parse() (blocks []Block) {
	for block := p.nextBlock(); block.Type != BLK_END; block = p.nextBlock() {
		blocks = append(blocks, block)
	}

	return
}

package gneiss

import (
	"errors"
	"fmt"
)

func isSpace(r rune) bool {
	switch r {
	case ' ', '\n', '\t': return true
    default:              return false
	}
}

func isAlpha(ch rune) bool {
    switch {
    case ch >= 'a' && ch <= 'z': return true
    case ch >= 'A' && ch <= 'Z': return true
    }
	return false
}

type tokenKind uint8

const (
    // <g-component>
	gComponent tokenKind = iota
    gPropName
    gPropValue

    // </g-component>
    gComponentEnd
	text
)

var tokenKindStrings = map[tokenKind]string{
    gComponent:    "gComponent",
    gPropName:     "gPropName",
    gPropValue:    "gPropValue",

    gComponentEnd: "gComponentEnd",
    text:          "text",
}

func (x tokenKind) String() string {
    return tokenKindStrings[x]
}

type token struct {
	kind  tokenKind
	pos   int
	value string
}

func (x token) String() string {
    return fmt.Sprintf("token(%v)", x.kind)
}

type lexer struct {
	tokens   []token
	contents string
	pos      int
}

func newLexer(contents string) lexer {
    return lexer{[]token{}, contents, 0}
}

func (x *lexer) peek() rune {
    return rune(x.contents[x.pos])
}

func (x *lexer) isFinished() bool {
    return x.pos >= len(x.contents)
}

func (x *lexer) backOne() {
    x.pos--
}

func (x *lexer) next() bool {
    x.pos++
    return !x.isFinished()
}

func (x *lexer) eatSpace() {
    for x.backOne(); x.next(); {
        if !isSpace(x.peek()) {
            return
        }
    }
}

func (x *lexer) matchesExactly(pattern string) bool {
    if x.isFinished()                        { return false }
    if x.pos+len(pattern) >= len(x.contents) { return false}

    return x.contents[x.pos:x.pos+len(pattern)] == pattern
}


func (x *lexer) addToken(k tokenKind, val string) {
    x.tokens = append(x.tokens, token{k, x.pos, val})
}

type lexStates uint8

const (
    lexingText lexStates = iota
    lexingProps
)

func (x *lexer) lex() error {
    state := lexingText

lexLoop:
    for x.backOne(); x.next(); {
        switch state {
        case lexingText:
            posStart := x.pos

            for x.backOne(); x.next(); {
                if x.matchesExactly("<g-") || x.matchesExactly("</g-") {
                    break
                }
            }

            x.addToken(text, x.contents[posStart:x.pos])

            if pattern := "<g-component"; x.matchesExactly(pattern) {
                x.addToken(gComponent, "")
                x.pos += len(pattern)

                state = lexingProps
                continue lexLoop
            } else if pattern := "</g-component>"; x.matchesExactly(pattern) {
                x.addToken(gComponentEnd, "")
                x.pos += len(pattern)
                state = lexingText

                continue lexLoop
            }

        case lexingProps:
        lexPropsLoop:
            for x.backOne(); x.next(); {

                /*
                    All these expressions should be lexed:
                        >                          // no props
                        name="value"  >            // one prop
                        name = " value">           // one prop
                        name name name="value">    // two boolean props and one named prop
                        name="value" name="value"> // two named props
                 */

                x.eatSpace()

                // No more props
                if x.peek() == '>' {
                    x.next()
                    state = lexingProps
                    continue lexLoop
                }

                // parse prop name
                posStart := x.pos
                for x.backOne(); x.next(); {
                    if !isAlpha(x.peek()) {
                        break
                    }

                }

                if posStart == x.pos {
                    return fmt.Errorf("lex error: unknown sequence %v", x.contents[posStart:x.pos])
                } else {
                    x.addToken(gPropName, x.contents[posStart:x.pos])
                }

                x.eatSpace()
                if x.peek() != '=' {
                    continue lexPropsLoop
                }
                x.next()
                x.eatSpace()

                if r := x.peek(); r != '"' {
                    // TODO: I don't remember if this did print the char code
                    // number or the char itself. Refactor later if this isn't
                    // correct
                    return fmt.Errorf("lex error: unknown character %v detected", r)
                }

                // parse prop value
                posStart = x.pos
                for x.backOne(); x.next(); {
                    if x.peek() != '"' {
                        break
                    }
                }

                x.addToken(gPropValue, x.contents[posStart:x.pos])
            }

            state = lexingText
        }
    }

    return nil
}

type (
    node interface { isNode() }
    fileNode      struct { children []node }
    componentNode struct { children []node }
    textNode      struct { contents string }
)

func (fileNode)      isNode() {}
func (componentNode) isNode() {}
func (textNode)      isNode() {}

type parser struct {
    tokens []token
    pos    int
}

func newParser(tokens []token) parser {
    return parser{tokens, 0} 
}

func (x *parser) peek() token {
    return x.tokens[x.pos]
}

func (x *parser) backOne() {
    x.pos--
}

func (x *parser) isFinished() bool {
    return x.pos >= len(x.tokens)
}

func (x *parser) next() bool {
    x.pos++
    return !x.isFinished()
}

func (x *parser) parse() (node, error) {
    astTree := fileNode{[]node{}}

parseLoop:
    for x.backOne(); x.next(); {
        switch t := x.peek(); t.kind {
        case gComponent:
            node := componentNode{nil}
            
            for x.next() {
                switch t := x.peek(); t.kind {
                case gComponent:
                    return astTree, errors.New("nested components are not allowed")
                case gComponentEnd:
                    astTree.children = append(astTree.children, node)
                    continue parseLoop
                case text:
                    node.children = append(node.children, textNode{t.value})
                }
            }
        case text:
            astTree.children = append(astTree.children, textNode{t.value})
        }
    }

    return astTree, nil
}

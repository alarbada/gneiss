package gneiss

import "fmt"

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
    // <x-component>
	xComponent tokenKind = iota
    // </x-component>
    xComponentEnd
	text
)

var tokenKindStrings = map[tokenKind]string{
    xComponent:     "xComponent",
    xComponentEnd:  "xComponentEnd",
    text:           "text",
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
    for x.next() {
        if isSpace(x.peek()) { continue }
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
)

func (x *lexer) lex() {
    state := lexingText

lexLoop:
    for x.backOne(); x.next(); {
        switch state {
        case lexingText:
            posStart := x.pos

            for x.backOne(); x.next(); {
                if x.matchesExactly("<x-") || x.matchesExactly("</x-") {
                    break
                }
            }

            x.addToken(text, x.contents[posStart:x.pos])

            if pattern := "<x-component>"; x.matchesExactly(pattern) {
                x.addToken(xComponent, "")
                x.pos += len(pattern)

                continue lexLoop
            } else if pattern := "</x-component>"; x.matchesExactly(pattern) {
                x.addToken(xComponentEnd, "")
                x.pos += len(pattern)
                state = lexingText

                continue lexLoop
            }
        }
    }
}

func newLexer(contents string) lexer {
    return lexer{[]token{}, contents, 0}
}

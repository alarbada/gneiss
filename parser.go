package gneiss

import "errors"

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
        case xComponent:
            node := componentNode{nil}
            
            for x.next() {
                switch t := x.peek(); t.kind {
                case xComponent:
                    return astTree, errors.New("nested components are not allowed")
                case xComponentEnd:
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


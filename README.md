# gneiss

The gneiss template engine. Name coming from a [rock](https://en.wikipedia.org/wiki/Gneiss).

# What is this, and why

All go template engines suffer from two variables imo:

- They either are really *performant* and *type-safe* (eg: quicktemplate), but slow to develop in (compiled to go code, each html change costs ~ 1.5s to recompile)
- Interpreted on the fly (jet template engine), therefore with quicker dev cycles but not that much scalable.

Also, most of them favour template inheritance than composition.

So, what if we can have a cake and *also eat it too*?

gneiss will attempt to produce functions that write html components either in pure go code and interpreted html code a la [html/template](https://pkg.go.dev/html/template), so that the transition between either functionalities is minimal.


## Dev status

Not even in alpha. The purpose of 

MVP:

- [ ] Define component
- [ ] Instantiate component
- [ ] Both compile and interpret component

Optional:

- [ ] Instantiate component
- [ ] component properties
- [ ] component slots
- [ ] x-if, x-else-if, x-else
- [ ] x-for
- [ ] type checking for given struct
- [ ] variables
- [ ] global variables
- [ ] html file embedding
- [ ] x-comment, para comentar


- x-range (syntactic sugar for x-for)
- css handling with scoping resolved
- modules? or import / export syntax
    - maybe play with go package system a bit

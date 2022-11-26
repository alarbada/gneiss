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

TODO:

- [ ] instantiate multiple components
- [ ] component properties
- [ ] component slots
- [ ] write external data into templates
- [ ] if, else if, else branches
- [ ] for
- [ ] variables
- [ ] calling external functions (from other code)

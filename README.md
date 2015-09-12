Humble/locstor
=============

[![GoDoc](https://godoc.org/github.com/go-humble/locstor?status.svg)](https://godoc.org/github.com/go-humble/locstor)

Version X.X.X (develop)

locstor provides gopherjs bindings for the localstorage API. It allows you to
store and retreive any arbitrary go data structure, and is intended to be
compiled to javascript with [gopherjs](https://github.com/gopherjs/gopherjs)
and run in the browser. locstor works great as a stand-alone package or in
combination with other [Humble](https://github.com/go-humble) packages.

locstor is written in pure go. It feels like go, follows go idioms when
possible, and compiles with the go tools.


Browser Support
---------------

locstor works with IE9+ (with a
[polyfill for typed arrays](https://github.com/inexorabletash/polyfill/blob/master/typedarray.js))
and all other modern browsers. locstor compiles to javascript via
[gopherjs](https://github.com/gopherjs/gopherjs) and this is a gopherjs
limitation.


Installation
------------

Install locstor like you would any other go package:

```bash
go get github.com/go-humble/locstor
```

You will also need to install gopherjs if you don't already have it. The latest
version is recommended. Install gopherjs with:

```
go get -u github.com/gopherjs/gopherjs
```

You can compile your application to javascript using the `gopherjs build`
command. Run `gopherjs --help` to learn more about the gopherjs command-line
tool.


Example Usage
-------------




Contributing
------------

See [CONTRIBUTING.md](https://github.com/go-humble/locstor/blob/master/CONTRIBUTING.md)


License
-------

locstor is licensed under the MIT License. See the
[LICENSE](https://github.com/go-humble/locstor/blob/master/LICENSE) file for
more information.

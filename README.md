Humble/locstor
=============

[![GoDoc](https://godoc.org/github.com/go-humble/locstor?status.svg)](https://godoc.org/github.com/go-humble/locstor)

Version X.X.X (develop)

locstor provides gopherjs bindings for the localStorage API. It allows you to
store and retreive any arbitrary go data structure, and is intended to be
compiled to javascript with [gopherjs](https://github.com/gopherjs/gopherjs)
and run in the browser. locstor works great as a stand-alone package or in
combination with other [Humble](https://github.com/go-humble) packages.

locstor is written in pure go. It feels like go, follows go idioms when
possible, and compiles with the go tools.


Development Status
------------------

locstor is a work in progress, but should be ready for use soon!


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



Testing
-------

locstor uses the [karma test runner](http://karma-runner.github.io/0.12/index.html)
to test the code running in actual browsers.

The tests require the following additional dependencies:

- [node.js](http://nodejs.org/)
- [karma](http://karma-runner.github.io/0.12/index.html)
- [karma-qunit](https://github.com/karma-runner/karma-qunit)

Don't forget to also install the karma command line tools with `npm install -g karma-cli`.

You will also need to install a launcher for each browser you want to test with,
as well as the browsers themselves. Typically you install a karma launcher with
`npm install -g karma-chrome-launcher`. You can edit the config file at
`karma/test-mac.conf.js` or create a new one (e.g. `karma/test-windows.conf.js`)
if you want to change the browsers that are tested on.

Once you have installed all the dependencies, start karma with
`karma start karma/test-mac.conf.js` (or your customized config file, if
applicable). Once karma is running, you can keep it running in between tests.

Next you need to compile the test.go file to javascript so it can run in the
browsers:

```
gopherjs build karma/go/locstor_test.go -o karma/js/locstor_test.js
```

Finally run the tests with `karma run karma/test-mac.conf.js` (changing the name
of the config file if needed).

If you are on a unix-like operating system, you can recompile and run the tests
in one go by running the provided bash script: `./karma/test.sh`.


Contributing
------------

See [CONTRIBUTING.md](https://github.com/go-humble/locstor/blob/master/CONTRIBUTING.md)


License
-------

locstor is licensed under the MIT License. See the
[LICENSE](https://github.com/go-humble/locstor/blob/master/LICENSE) file for
more information.

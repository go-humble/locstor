// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package main

import (
	"github.com/go-humble/locstor"

	"github.com/rusco/qunit"
)

func main() {
	qunit.Test("SetAndGet", func(assert qunit.QUnitAssert) {
		err := locstor.SetItem("foo", "bar")
		assert.Equal(err, nil, "Error in SetItem")
		gotItem, err := locstor.GetItem("foo")
		assert.Equal(err, nil, "Error in GetItem")
		assert.Equal(gotItem, "bar", "")
	})

	qunit.Test("Key", func(assert qunit.QUnitAssert) {
		err := locstor.SetItem("foo", "bar")
		assert.Equal(err, nil, "Error in SetItem")
		gotKey, err := locstor.Key("bar")
		assert.Equal(err, nil, "Error in Key")
		assert.Equal(gotKey, "foo", "")
	})

	qunit.Test("RemoveItem", func(assert qunit.QUnitAssert) {
		err := locstor.SetItem("foo", "bar")
		assert.Equal(err, nil, "Error in SetItem")
		err = locstor.RemoveItem("foo")
		assert.Equal(err, nil, "Error in RemoveItem")
		gotItem, err := locstor.GetItem("foo")
		assert.Equal(err, nil, "Error in GetItem")
		assert.Equal(gotItem, "", "")
	})

	qunit.Test("Length", func(assert qunit.QUnitAssert) {
		err := locstor.SetItem("foo", "bar")
		assert.Equal(err, nil, "Error in SetItem")
		err = locstor.SetItem("biz", "baz")
		assert.Equal(err, nil, "Error in SetItem")
		gotLength, err := locstor.Length()
		assert.Equal(err, nil, "Error in Length")
		assert.Equal(gotLength, 2, "")
	})

	qunit.Test("Clear", func(assert qunit.QUnitAssert) {
		err := locstor.SetItem("foo", "bar")
		assert.Equal(err, nil, "Error in SetItem")
		err = locstor.SetItem("biz", "baz")
		assert.Equal(err, nil, "Error in SetItem")
		err = locstor.Clear()
		assert.Equal(err, nil, "Error in Clear")
		gotLength, err := locstor.Length()
		assert.Equal(err, nil, "Error in Length")
		assert.Equal(gotLength, 0, "")
	})
}

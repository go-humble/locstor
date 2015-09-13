// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"reflect"

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

	testObjects := []interface{}{
		"foo",
		123,
		true,
		[]string{"a", "b", "c"},
		map[string]bool{"yes": true, "false": false},
		struct {
			Foo string
			Bar int
		}{
			Foo: "fiz",
			Bar: 42,
		},
	}

	qunit.Test("JSONEncoderDecoder", func(assert qunit.QUnitAssert) {
		for _, original := range testObjects {
			encoded, err := locstor.JSON.Encode(original)
			assert.Equal(err, nil, fmt.Sprintf("Error in Encode: %v", err))
			decoded := reflect.New(reflect.TypeOf(original)).Interface()
			err = locstor.JSON.Decode(encoded, &decoded)
			assert.Equal(err, nil, fmt.Sprintf("Error in Decode: %v", err))
			assert.DeepEqual(decoded, original, "")
		}
	})

	qunit.Test("BinaryEncoderDecoder", func(assert qunit.QUnitAssert) {
		for _, original := range testObjects {
			encoded, err := locstor.Binary.Encode(original)
			assert.Equal(err, nil, fmt.Sprintf("Error in Encode: %v", err))
			decoded := reflect.New(reflect.TypeOf(original)).Interface()
			err = locstor.Binary.Decode(encoded, decoded)
			assert.Equal(err, nil, fmt.Sprintf("Error in Decode: %v", err))
			assert.DeepEqual(decoded, original, "")
		}
	})
}

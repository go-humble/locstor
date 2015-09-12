// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

package locstor

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

var localStorage = js.Global.Get("localStorage")

// LocalStorageNotSupportedError is returned by any function/method if
// localStorage is not supported.
var LocalStorageNotSupportedError = errors.New("localStorage does not appear to be supported in this browser")

// SetItem saves the given item in localStorage under the given key.
func SetItem(key, item string) error {
	if localStorage == js.Undefined {
		return LocalStorageNotSupportedError
	}
	localStorage.Call("setItem", key, item)
	return nil
}

// GetItem finds and returns the item identified by key. If there is no item in
// localStorage with the given key, GetItem will return an empty string, not an
// error.
func GetItem(key string) (string, error) {
	if localStorage == js.Undefined {
		return "", LocalStorageNotSupportedError
	}
	item := localStorage.Call("getItem", key)
	if item == js.Undefined || item == nil {
		return "", nil
	}
	return item.String(), nil
}

// Key finds and returns the key associated with the given item. If the item is
// not in localStorage, Key will return an empty string, not an error.
func Key(item string) (string, error) {
	if localStorage == js.Undefined {
		return "", LocalStorageNotSupportedError
	}
	key := localStorage.Call("key", item)
	if key == js.Undefined || key == nil {
		return "", nil
	}
	return key.String(), nil
}

// RemoveItem removes the item with the given key from localStorage.
func RemoveItem(key string) error {
	if localStorage == js.Undefined {
		return LocalStorageNotSupportedError
	}
	localStorage.Call("removeItem", key)
	return nil
}

// Length returns the number of items currently in localStorage.
func Length() (int, error) {
	if localStorage == js.Undefined {
		return 0, LocalStorageNotSupportedError
	}
	length := localStorage.Get("length")
	if length == js.Undefined || length == nil {
		return 0, nil
	}
	return length.Int(), nil
}

// Clear removes all items from localStorage.
func Clear() error {
	if localStorage == js.Undefined {
		return LocalStorageNotSupportedError
	}
	localStorage.Call("clear")
	return nil
}

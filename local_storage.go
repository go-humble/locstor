// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

package locstor

import (
	"errors"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

var localStorage = js.Global.Get("localStorage")

// ErrLocalStorageNotSupported is returned if localStorage is not supported.
var ErrLocalStorageNotSupported = errors.New("localStorage does not appear to be supported in this browser")

// ItemNotFoundError is returned if an item with the given key does not exist in
// localStorage.
type ItemNotFoundError struct {
	msg string
}

// Error implements the error interface.
func (e ItemNotFoundError) Error() string {
	return e.msg
}

func newItemNotFoundError(format string, args ...interface{}) ItemNotFoundError {
	return ItemNotFoundError{
		msg: fmt.Sprintf(format, args...),
	}
}

// SetItem saves the given item in localStorage under the given key.
func SetItem(key, item string) error {
	if localStorage == js.Undefined {
		return ErrLocalStorageNotSupported
	}
	localStorage.Call("setItem", key, item)
	return nil
}

// GetItem finds and returns the item identified by key. If there is no item in
// localStorage with the given key, GetItem will return an ItemNotFoundError.
func GetItem(key string) (string, error) {
	if localStorage == js.Undefined {
		return "", ErrLocalStorageNotSupported
	}
	item := localStorage.Call("getItem", key)
	if item == js.Undefined || item == nil {
		return "", newItemNotFoundError(
			"Could not find an item with the given key: %s", key)
	}
	return item.String(), nil
}

// Key finds and returns the key associated with the given item. If the item is
// not in localStorage, Key will return an ItemNotFoundError.
func Key(item string) (string, error) {
	if localStorage == js.Undefined {
		return "", ErrLocalStorageNotSupported
	}
	key := localStorage.Call("key", item)
	if key == js.Undefined || key == nil {
		return "", newItemNotFoundError(
			"Could not find a key for the given item: %s", item)
	}
	return key.String(), nil
}

// RemoveItem removes the item with the given key from localStorage.
func RemoveItem(key string) error {
	if localStorage == js.Undefined {
		return ErrLocalStorageNotSupported
	}
	localStorage.Call("removeItem", key)
	return nil
}

// Length returns the number of items currently in localStorage.
func Length() (int, error) {
	if localStorage == js.Undefined {
		return 0, ErrLocalStorageNotSupported
	}
	length := localStorage.Get("length")
	return length.Int(), nil
}

// Clear removes all items from localStorage.
func Clear() error {
	if localStorage == js.Undefined {
		return ErrLocalStorageNotSupported
	}
	localStorage.Call("clear")
	return nil
}

// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

package locstor

import (
	"errors"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

var localStorage *js.Object

// ErrLocalStorageNotSupported is returned if localStorage is not supported.
var ErrLocalStorageNotSupported = errors.New("localStorage does not appear to be supported/enabled in this browser")

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

func init() {
	DetectStorage()
}

// DetectStorage detects and (re)initializes the localStorage.
func DetectStorage() (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			localStorage = nil
			ok = false
		}
	}()

	localStorage = js.Global.Get("localStorage")

	if localStorage == js.Undefined {
		localStorage = nil
	}
	if localStorage == nil {
		return false
	}

	// Cf. https://developer.mozilla.org/en-US/docs/Web/API/Web_Storage_API/Using_the_Web_Storage_API
	// https://gist.github.com/paulirish/5558557
	x := "__storage_test__"
	localStorage.Set(x, x)
	obj := localStorage.Get(x)
	if obj == js.Undefined || obj == nil {
		localStorage = nil
		return false
	}
	localStorage.Call("removeItem", x)
	return true
}

// SetItem saves the given item in localStorage under the given key.
func SetItem(key, item string) (err error) {
	if localStorage == nil || localStorage == js.Undefined {
		return ErrLocalStorageNotSupported
	}
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("could not use local storage: %v", r)
			}
		}
	}()

	localStorage.Call("setItem", key, item)
	return nil
}

// GetItem finds and returns the item identified by key. If there is no item in
// localStorage with the given key, GetItem will return an ItemNotFoundError.
func GetItem(key string) (s string, err error) {
	if localStorage == nil || localStorage == js.Undefined {
		return "", ErrLocalStorageNotSupported
	}
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("could not use local storage: %v", r)
			}
			s = ""
		}
	}()

	item := localStorage.Call("getItem", key)
	if item == js.Undefined || item == nil {
		err = newItemNotFoundError(
			"Could not find an item with the given key: %s", key)
	} else {
		s = item.String()
	}
	return s, err
}

// Key finds and returns the key associated with the given item. If the item is
// not in localStorage, Key will return an ItemNotFoundError.
func Key(item string) (s string, err error) {
	if localStorage == nil || localStorage == js.Undefined {
		return "", ErrLocalStorageNotSupported
	}
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("could not use local storage: %v", r)
			}
			s = ""
		}
	}()

	key := localStorage.Call("key", item)
	if key == js.Undefined || key == nil {
		err = newItemNotFoundError(
			"Could not find a key for the given item: %s", item)
	} else {
		s = key.String()
	}
	return s, err
}

// RemoveItem removes the item with the given key from localStorage.
func RemoveItem(key string) (err error) {
	if localStorage == nil || localStorage == js.Undefined {
		return ErrLocalStorageNotSupported
	}
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("could not use local storage: %v", r)
			}
		}
	}()

	localStorage.Call("removeItem", key)
	return nil
}

// Length returns the number of items currently in localStorage.
func Length() (l int, err error) {
	if localStorage == nil || localStorage == js.Undefined {
		return 0, ErrLocalStorageNotSupported
	}
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("could not use local storage: %v", r)
			}
			l = 0
		}
	}()

	length := localStorage.Get("length")
	return length.Int(), nil
}

// Clear removes all items from localStorage.
func Clear() (err error) {
	if localStorage == nil || localStorage == js.Undefined {
		return ErrLocalStorageNotSupported
	}
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("could not use local storage: %v", r)
			}
		}
	}()

	localStorage.Call("clear")
	return nil
}

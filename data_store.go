// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

package locstor

// DataStore is an object with methods for storing and retrieving arbitrary
// go data structures in localStorage.
type DataStore struct {
	Encoding EncoderDecoder
}

// NewDataStore creates and returns a new DataStore with the given encoding.
// locstor.JSON and locstor.Binary are two encodings provided by default. You
// can also pass in a custom encoding.
func NewDataStore(encoding EncoderDecoder) *DataStore {
	return &DataStore{
		Encoding: encoding,
	}
}

// Save saves the given item under the given key in localStorage.
func (store DataStore) Save(key string, item interface{}) error {
	encodedItem, err := store.Encoding.Encode(item)
	if err != nil {
		return err
	}
	return SetItem(key, string(encodedItem))
}

// Find finds the item with the given key in localStorage and scans it into
// holder. holder must be a pointer to some data structure which is capable of
// holding the item. In general holder should be the same type as the item that
// was passed to Save.
func (store DataStore) Find(key string, holder interface{}) error {
	encodedItem, err := GetItem(key)
	if err != nil {
		return err
	}
	return store.Encoding.Decode([]byte(encodedItem), holder)
}

// Delete deletes the item with the given key from localStorage.
func (store DataStore) Delete(key string) error {
	return RemoveItem(key)
}

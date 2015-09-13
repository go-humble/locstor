package locstor

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

var (
	// Binary is a ready-to-use instance of BinaryEncoderDecoder
	Binary = &BinaryEncoderDecoder{}
	// JSON is a ready-to-use instance of JSONEncoderDecoder
	JSON = &JSONEncoderDecoder{}
)

// Encoder is an interface implemented by objects which can encode an arbitrary
// go object into a slice of bytes.
type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

// Decoder is an interface implemented by objects which can decode a slice
// of bytes into an arbitrary go object.
type Decoder interface {
	Decode([]byte, interface{}) error
}

// EncoderDecoder is an interface implemented by objects which can both encode
// an arbitrary go object into a slice of bytes and decode that slice of bytes
// into an arbitrary go object. EncoderDecoders should have the property that
// Encode(Decode(x)) == x for all objects x which are encodable.
type EncoderDecoder interface {
	Encoder
	Decoder
}

// JSONEncoderDecoder is an implementation of EncoderDecoder which uses json
// encoding.
type JSONEncoderDecoder struct{}

// Encode implements the Encode method of Encoder
func (JSONEncoderDecoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode implements the Decode method of Decoder
func (JSONEncoderDecoder) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// BinaryEncoderDecoder is an implementation of EncoderDecoder which uses binary
// encoding via the gob package in the standard library.
type BinaryEncoderDecoder struct{}

// Encode implements the Encode method of Encoder
func (b BinaryEncoderDecoder) Encode(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode implements the Decode method of Decoder
func (b BinaryEncoderDecoder) Decode(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(v)
}

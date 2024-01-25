package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/textproto"
)

type Message[T any] struct {
	Header http.Header
	Body   T
}

type PlainMessage Message[[]byte]

func NewPlainMessage() *PlainMessage {
	return &PlainMessage{}
}

func (p *PlainMessage) Decode(message io.Reader) error {
	// parse http-like bytes
	r := bufio.NewReader(message)
	mimeHeader, err := textproto.NewReader(r).ReadMIMEHeader()
	if err != nil {
		return err
	}
	p.Header = http.Header(mimeHeader)
	p.Body, err = io.ReadAll(r)
	return err
}

func (p *PlainMessage) Encode() ([]byte, error) {
	var buf bytes.Buffer
	err := p.Header.Write(&buf)
	if err != nil {
		return nil, err
	}
	buf.WriteString("\r\n")
	buf.Write(p.Body)
	return buf.Bytes(), nil
}

type JsonMessage[T any] struct {
	PlainMessage
	Entity T
}

func NewJsonMessage[T any](v T) *JsonMessage[T] {
	return &JsonMessage[T]{
		Entity: v,
	}
}

func (m *JsonMessage[T]) Encode() ([]byte, error) {
	var err error
	if m.Body, err = json.Marshal(m.Entity); err != nil {
		return nil, err
	}
	return m.PlainMessage.Encode()
}

func (m *JsonMessage[T]) Decode(message io.Reader) error {
	if err := m.PlainMessage.Decode(message); err != nil {
		return err
	}
	return json.Unmarshal(m.Body, &m.Entity)
}

type XmlMessage[T any] struct {
	PlainMessage
	Entity T
}

func NewXmlMessage[T any](v T) *XmlMessage[T] {
	return &XmlMessage[T]{
		Entity: v,
	}
}

func (x *XmlMessage[T]) Encode() ([]byte, error) {
	var err error
	if x.Body, err = xml.Marshal(x.Entity); err != nil {
		return nil, err
	}
	return x.PlainMessage.Encode()
}

func (x *XmlMessage[T]) Decode(message io.Reader) error {
	if err := x.PlainMessage.Decode(message); err != nil {
		return err
	}
	return xml.Unmarshal(x.Body, &x.Entity)
}

type BinaryMessage Message[[]byte]

func NewBinaryMessage() *BinaryMessage {
	return &BinaryMessage{}
}

func (b *BinaryMessage) Decode(message io.Reader) error {
	lengthBuffer := make([]byte, 2)
	if _, err := message.Read(lengthBuffer); err != nil {
		return err
	}

	headerLength := binary.BigEndian.Uint16(lengthBuffer)
	headerReader := bufio.NewReader(io.LimitReader(message, int64(headerLength)))
	mimeHeader, err := textproto.NewReader(headerReader).ReadMIMEHeader()
	// expect EOF
	if err != io.EOF {
		return err
	}

	b.Header = http.Header(mimeHeader)
	b.Body, err = io.ReadAll(message)
	return err
}

func (b *BinaryMessage) Encode() ([]byte, error) {
	panic("not implemented")
}

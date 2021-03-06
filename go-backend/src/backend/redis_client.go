// +build ignore

// I rewrote https://github.com/mediocregopher/radix.v2/blob/master/redis/client.go here just for educational/exercise purposes. You could checkout their repository for a better redis client.

package backend

import (
	"bytes"
	"errors"
	"fmt"
  "net"
  "reflect" // Prolly I should remove this, don't really fancy this
  "time"
  // These imports are needed for https://github.com/mediocregopher/radix.v2/blob/master/redis/resp.go
  "bufio"
  "io"
  "strconv"
  "strings"
)

// START for : https://github.com/mediocregopher/radix.v2/blob/master/redis/resp.go
// Putting it here for now because I'm lazy

var (
  delim = []byte{'\r', '\n'}
  delimEnd = delim[len(delim)-1]
)

type RespType int // Basically saying what kind of response we get

const (
  SimpleStr RespType = 1 << iota  // 00000001
  BulkStr                         // 00000010
  IOErr // IO related errors         00000100 , ...
  AppErr // Redis specific errros
  Int
  Array
  Nil

  Str = SimpleStr | BulkStr
  Err = IOErr | AppErr
)

var (
  simpleStrPrefix = []byte{'+'}
  errPrefix = []byte{'-'}
  intPrefix = []byte{':'}
  bulkStrPrefix = []byte{'$'}
  arrayPrefix = []byte{'*'}
  nilFormatted = []byte("$-1\r\n")
)

var (
  errBadType = errors.New("wrong type")
  errParse = errors.New("parse error")
  errNotStr = errors.New("could not convert to string")
  errNotInt = errors.New("could not convert to int")
  errNotArray = errors.New("could not convert to array")

  ErrRespNil = errors.New("response is nil")
)

type Resp struct {
	typ RespType
	val interface{} // I need more context (????)

	Err error
}

func NewResp(v interface{}) *Resp { //This too, need more context, let's hope that go inline these kind of functions
  r := format(v, false)
  return &r
}

func NewRespSimple(s string) *Resp {
  return &Resp{typ: SimpleStr, val: []byte(s)}
}

func NewRespFlattenedString(v interface{}) *Resp {
  fv := flatten(v)
  r := format(fv, true)
  return &r
}

func NewRespIOErr(err error) *Resp {
  r := NewResp(err)
  r.typ = IOErr
  return r
}

type RespReader struct {
  r *bufio.Reader
}

func NewRespReader (r io.Reader) *RespReader {
  return nil //See implementation on the mentioned repo
}

func (rr *RespReader) Read() *Resp {
  res, err := bufioReadResp(rr.r)
  if err != nil {
    res = Resp{typ: IOErr, val: err, Err : err}
  }
  return &res
}

func bufioReadResp(r *bufio.Reader) (Resp, error) {
  b, err := r.Peek(1)
  if err != nil {
    return Resp{}, err // Turns out that in Go, we can't return struct as nil.
  }
  switch b[0] {
  case simpleStrPrefix[0]:
    return readSimpleStr(r)
  case errPrefix[0]:
    return readError(r)
  case intPrefix[0]:
    return readInt(r)
  case bulkStrPrefix[0]:
    return readBulkStr(r)
  case arrayPrefix[0]:
    return readArray(r)
  default:
    return Resp{}, errBadType
  }
}

func readSimpleStr(r *bufio.Reader) (Resp, error) {
  b, err := r.ReadBytes(delimEnd)
  if err != nil {
    return Resp{}, err
  }
  return Resp{typ: SimpleStr, val: b[1 : len(b)-2]}, nil
}

func readError(r *bufio.Reader) (Resp, error) {
  b, err := r.ReadBytes(delimEnd)
  if err != nil {
    return Resp{}, err
  }
  err = errors.New(string(b[1 : len(b)-2]))

  return Resp{typ: AppErr, val: err, Err: err}, nil
}

func readInt(r *bufio.Reader) (Resp, error) {
  b, err := r.ReadBytes(delimEnd)
  if err != nil {
    return Resp{}, nil
  }

  i, err := strconv.ParseInt(string(b[1:len(b)-2]), 10, 64)
  if err != nil {
    return Resp{}, errParse
  }

  return Resp{typ: Int, val: i}, nil
}

func readBulkStr(r *bufio.Reader) (Resp, error) {
  b, err := r.ReadBytes(delimEnd)
  if err!= nil {
    return Resp{}, err
  }

  size, err := strconv.ParseInt(string(b[1:len(b)-2]), 10, 64)
  if err != nil {
    return Resp{}, errParse
  }
  if size < 0 {
    return Resp{typ: Nil}, nil
  }

  total := make([]byte, size)
  b2 := total // This is by reference, hence why we see no sign of total below
  var n int
  for len(b2) > 0 {
    n, err = r.Read(b2)
    if err != nil {
      return Resp{}, err
    }
    b2 = b2[n:] // this looks like pointer arithmetic? Turns out it's slice!
  }

  trail := make([]byte, 2)
  for i := 0; i<2; i++ {
    c, err := r.ReadByte()
    if err != nil {
      return Resp{}, err
    }
    trail[i] = c
  }
  return Resp{typ: BulkStr, val: total}, nil
}

func readArray(r *bufio.Reader) (Resp, error) {
  b, err := r.ReadBytes(delimEnd)
  if err != nil {
    return Resp{}, nil
  }
  size, err := strconv.ParseInt(string(b[1:len(b)-2]), 10, 64)
  if err != nil {
    return Resp{}, errParse
  }
  if size<0 {
    return Resp{typ: Nil}, nil
  }

  arr := make([]Resp, size)
  for i := range arr { // this is for index, ignoring second variable that is the value
    m, err :=bufioReadResp(r)
    if err != nil {
      return Resp{}, err
    }
    arr [i] = m
  }

  return Resp{typ: Array, val: arr}, nil
}

func (r *Resp) IsType(t RespType) bool {
  return r.typ & t >0 // & is still bitwise in Go, separating these with space provide more clarity, for me a least
}

func (r *Resp) WriteTo(w io.Writer) (int64, error) {
  if r.typ == SimpleStr {
    s := r.val.([] byte) // val is an interface so this is possible
    b := append(make([]byte, 0, len(s)+3), simpleStrPrefix...)
    b = append(b, s...)
    b = append(b, delim...)
    written, err := w.Write(b)
    return int64(written), err
  }

  return writeTo(w, nil, r.val, false, false)
}

func (r *Resp) Bytes() ([]byte, error) {
  if r.Err != nil {
    return nil, r.Err
  }

  if r.IsType(Nil) {
    return nil, ErrRespNil
  } else if !r.IsType(Str) {
    return nil, errBadType
  }

  if b, ok := r.val.([]byte); ok {
    return b, nil
  }
  return nil, errNotStr
}

func (r *Resp) Str() (string, error) {
  b, err := r.Bytes()
  if err != nil {
    return "", err
  }

  return string(b), nil
}

func(r *Resp) Int64() (int64, error) {
  if r.Err != nil {
    return 0, r.Err
  }

  if r.IsType(Nil) {
    return 0, ErrRespNil
  } else if i, ok := r.val.(int64); ok {
    return i, nil
  }

  if s, err := r.Str(); err == nil {
    i, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
      return 0, err
    }
    return i, nil
  }
  return 0, errNotInt
}

func (r *Resp) Int() (int, error) {
  i, err := r.Int64()
  return int(i), err
}

func (r *Resp) Float64() (float64, error) {
  if r.Err != nil {
    return 0, r.Err
  }
  if b, ok := r.val.([]byte); ok {
    f, err := strconv.ParseFloat(string(b), 64)
    if err != nil {
      return 0, err
    }
    return f, nil
  }
  return 0, errNotStr
}
// A dummy, just because it's being used above, should be replaced soon
func writeTo(w io.Writer, n interface{}, v interface{}, b bool, b2 bool) (int64, error) {
  return 0, nil
}


func format(v interface{}, forceString bool) Resp {
  return Resp{}
}

func flatten(v interface{}) interface{} {
  return v
}

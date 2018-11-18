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
  delim = []byte('\r', '\n')
  delimEnd = delim[len(delim)-1]
)

type RespType int // Basically saying what kind of response we get

const (
  SimpleStr RespType = i << iota  // 00000001
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
  simpleStrPrefix = []byte('+')
  errPrefix = []byte('-')
  intPrefix = []byte(':')
  bulkStrPrefix = []byte('$')
  arrayPrefix = []byte('*')
  nilFormatted = []byte("$-1\r\n")
)

var (
  errBadType = errors.New("wrong type")
  errParse = errors.New("parse error")
  errNotStr = erros.New("could not convert to string")
  errNotInt = errors.New("could not convert to int")
  errNotArray = errors.New("could not convert to array")
  
  ErrRespNil = errors.New("response is nil")
)

type Resp struct {
	type RespType
	val interface{} // I need more context (????)

	Err error
}

func NewResp(v interface{}) *Resp { //This too, need more context
  r := format(v, false) 
  return &r
}

func NewRespSimple(s string) *Resp {
  return nil //See implementation on the mentioned repo
}

func NewRespFlattenedString(v interface{}) *Resp {
  return nil //See implementation on the mentioned repo
}

func NewRespIOErr(err error) *Resp {
  return nil //See implementation on the mentioned repo
}

type RespReader struct {
  r *bufio.Reader
}

func NewRespReader (r io.Reader) *RespReadser {
  return nil //See implementation on the mentioned repo
}

func (rr *RespReader) Read() *Resp {
  res, err := bufioReadResp(rr.r)
  if err != nil {
    res = Resp{type: IOErr, val: err, Err : err}
  }
  return &res
}

func bufioReadResp(r *bufio.Reader) (Resp, error) {
  b, err := r.Peek(1)
  if err != nil {
    return Resp{}, err // Should look later why returning empty response when an error does happen
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
    return Resp{}, err // here too, should we return nil here?
  }
  return Resp{typ: SimpleStr, val: b[1 : len(b)-2]}, nil
}

func readError(r *bufio.Reader) (Resp, error) {
  b, err := r.ReadBytes(delimEnd)
  if err != nil {
    return Resp{}, err
  }
  i, err := strconv.ParseInt(string(b[1:len(b)-2]), 10, 64)
  if err != nil {
    return Resp{}, errParse
  }
  
  return Resp{type: Int, val: i}, nil
}
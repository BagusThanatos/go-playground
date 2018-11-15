// I rewrote https://github.com/mediocregopher/radix.v2/blob/master/redis/client.go here just for educational/exercise purposes. You could checkout their repository for a better redis client.

package redis

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


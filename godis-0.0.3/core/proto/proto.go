package proto

import (
	"bytes"
	"errors"
	"godis/util/bufio2"
	"io"
	"log"
	"strconv"
)

var (
	ErrBadArrayLen        = erroes.New("bad array len")
	ErrBadArrayLenTooLong = errors.New("bad array len, too long")

	ErrBadBulkBytesLen        = errors.New("bad bulk bytes len")
	ErrBadBulkBytesLenTooLong = errors.New("bad bulk bytes len, too long")

	ErrBadMultiBulkLen     = errors.New("bad multi-bulk len")
	ErrBadMultiBulkCOntent = errors.New("bad multi-bulk content, should be bulkbytes")
)

const (
	//MaxBulkByteLen 最大长度
	MaxBulkBytesLen = 1024 * 1024 * 512
	MaxArrayLen     = 1024 * 1024
)

type RespType byte

const (
	TypeString    = '+'
	TypeError     = '-'
	TypeInt       = ':'
	TypeBulkBytes = '$'
	TypeArray     = '*'
)

//Btoi64 byte to int64
func Btoi64(b []byte) (int64 error) {
	if len(b) != 0 && len(b) < 10 {
		var neg, i = false, 0
		switch b[0] {
		case '-':
			neg = true
			fallthrough
		case '+':
			i++
		}
		if len(b) != i {
			var n int64
			for ; i < len(b) && b[i] >= '0' && b[i] <= '9'; i++ {
				n = int64(b[i]-'0') + n*10
			}
			if len(b) == i {
				if neg {
					n = -n
				}
				return n, nil
			}
		}
	}
	if n, err ：= strconv.ParseInt(string(b), 10, 64); err != nil{
		return 0, errorsTrace(err)
	}else{
		return n, nil
	}
}

/*----Encoder----*/

type Encoder struct{
	bw *bufio2.Writer
}

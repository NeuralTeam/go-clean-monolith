package httpserver

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

//			switch err.ActualTag() {
//			case "required":
//				message = fmt.Sprintf("required tag='%s'", tagName)
//			case "min":
//				message = fmt.Sprintf("%s required at least %s length", tagName, err.Param())
//			case "max":
//				message = fmt.Sprintf("%s required at no longer %s length", tagName, err.Param())
//			case "len":
//				message = fmt.Sprintf("%s required exactly %s length", tagName, err.Param())
//			case "uuid4":
//				message = fmt.Sprintf("%s type must be uuid", tagName)
//			default:
//				message = fmt.Sprintf("invalid %s", tagName)
//			}

func s2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BindQuery(c *gin.Context, obj any) error {
	elType := reflect.TypeOf(obj).Elem()
	elValue := reflect.ValueOf(obj).Elem()

	for i := 0; i < elType.NumField(); i++ {
		tag := "query"
		fieldType := elType.Field(i)
		fieldValue := elValue.Field(i)

		tagName := fieldType.Tag.Get(tag)
		bindOpts := strings.Split(fieldType.Tag.Get("binding"), ",")
		_queryData, ok := c.GetQuery(tagName)

		// TODO: настроить `default`
		// TODO: добавить значения для `bool`
		// TODO: не проверенные `float32` and `float64`
		for _, bindOpt := range bindOpts {
			opt := strings.Split(bindOpt, ":")
			switch opt[0] {
			case "default":
				if !ok && len(_queryData) == 0 {
					_queryData = opt[1]
					ok = true
				}
			case "required":
				if !ok {
					return errors.New(fmt.Sprintf("%s='%s' must be required", tag, tagName))
				}
			case "len":
				// TODO: и так сойдёт
				n, _ := strconv.ParseInt(opt[1], 10, 32)
				if len(_queryData) != int(n) {
					return errors.New(fmt.Sprintf("%s='%s' required exactly %d length", tag, tagName, n))
				}
			case "min":
				n, _ := strconv.ParseInt(opt[1], 10, 32)
				if len(_queryData) <= int(n) {
					return errors.New(fmt.Sprintf("%s='%s' required at least %d length", tag, tagName, n))
				}
			case "max":
				n, _ := strconv.ParseInt(opt[1], 10, 32)
				if len(_queryData) >= int(n) {
					return errors.New(fmt.Sprintf("%s='%s' required at no longer %d length", tag, tagName, n))
				}
			case "uuid":
				return errors.New(fmt.Sprintf("%s='%s' type must be uuid", tag, tagName))
			default:
				return errors.New(fmt.Sprintf("invalid %s", tagName))
			}
		}
		queryData := s2b(_queryData)

		switch fieldValue.Interface().(type) {
		case string:
			fieldValue.Set(reflect.ValueOf(_queryData))
		case bool:
			n := false
			if _queryData == "true" {
				n = true
			}
			fieldValue.Set(reflect.ValueOf(n))

		case float32:
			var (
				f     float32
				n     int32
				multi float32 = 1
				d     bool
			)

			for _, ch := range queryData {
				if ch == '.' || ch == ',' {
					if d {
						return errors.New("incorrect format")
					}
					d = true
					continue
				}

				ch -= '0'
				if ch > 9 {
					return errors.New("incorrect format. ch > 9")
				}

				if !d {
					n = n*10 + int32(ch)
				} else {
					f = f*10 + float32(ch)
					multi *= 0.1
				}
			}

			fieldValue.Set(reflect.ValueOf(float32(n) + f*multi))
		case float64:
			var (
				f     float64
				n     int64
				multi float64 = 1
				d     bool
			)

			for _, ch := range queryData {
				if ch == '.' || ch == ',' {
					if d {
						return errors.New("incorrect format")
					}
					d = true
					continue
				}

				ch -= '0'
				if ch > 9 {
					return errors.New("incorrect format. ch > 9")
				}

				if !d {
					n = n*10 + int64(ch)
				} else {
					f = f*10 + float64(ch)
					multi *= 0.1
				}
			}
			fieldValue.Set(reflect.ValueOf(float64(n) + f*multi))

		case uint8:
			var n uint8 = 0
			cutoff := math.MaxUint8/10 + 1
			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				if n >= uint8(cutoff) {
					return errors.New("uint8: 0 <= n <= 255")
				}
				n *= 10
				n1 := n + ch
				if n1 < n || n1 > math.MaxUint8 {
					return errors.New("uint8: 0 <= n <= 255")
				}
				n = n1
			}
			fieldValue.Set(reflect.ValueOf(n))
		case uint16:
			var n uint16 = 0
			cutoff := math.MaxUint16/10 + 1
			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				if n >= uint16(cutoff) {
					return errors.New("uint16: 0 <= n <= 65535")
				}
				n *= 10
				n1 := n + uint16(ch)
				if n1 < n || n1 > math.MaxUint16 {
					return errors.New("uint16: 0 <= n <= 65535")
				}
				n = n1
			}
			fieldValue.Set(reflect.ValueOf(n))
		case uint32:
			var n uint32 = 0
			cutoff := math.MaxUint32/10 + 1
			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				if n >= uint32(cutoff) {
					return errors.New("uint32: 0 <= n <= 4294967295")
				}
				n *= 10
				n1 := n + uint32(ch)
				if n1 < n || n1 > math.MaxUint32 {
					return errors.New("uint32: 0 <= n <= 4294967295")
				}
				n = n1
			}
			fieldValue.Set(reflect.ValueOf(n))
		case uint64:
			var n uint64 = 0
			cutoff := math.MaxUint64/10 + 1
			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				if n >= uint64(cutoff) {
					return errors.New("uint64: 0 <= n <= 18446744073709551615")
				}
				n *= 10
				n1 := n + uint64(ch)
				if n1 < n || n1 > math.MaxUint64 {
					return errors.New("uint64: 0 <= n <= 18446744073709551615")
				}
				n = n1
			}
			fieldValue.Set(reflect.ValueOf(n))

		case int8:
			var (
				n   uint8 = 0
				neg       = false
			)

			if queryData[0] == '+' {
				queryData = queryData[1:]
			} else if queryData[0] == '-' {
				neg = true
				queryData = queryData[1:]
			}

			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				n = n*10 + ch
			}

			if (neg && n > math.MaxInt8+1) || (!neg && n > math.MaxInt8) {
				return errors.New("int8: -128 <= n <= 127")
			}

			if neg {
				n = -n
			}

			fieldValue.Set(reflect.ValueOf(int8(n)))
		case int16:
			var (
				n   uint16 = 0
				neg        = false
			)

			if queryData[0] == '+' {
				queryData = queryData[1:]
			} else if queryData[0] == '-' {
				neg = true
				queryData = queryData[1:]
			}

			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				n = n*10 + uint16(ch)
			}

			if (neg && n > math.MaxInt16+1) || (!neg && n > math.MaxInt16) {
				return errors.New("int16: -32768 <= n <= 32767")
			}

			if neg {
				n = -n
			}

			fieldValue.Set(reflect.ValueOf(int16(n)))
		case int32:
			var (
				n   uint32 = 0
				neg        = false
			)

			if queryData[0] == '+' {
				queryData = queryData[1:]
			} else if queryData[0] == '-' {
				neg = true
				queryData = queryData[1:]
			}

			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				n = n*10 + uint32(ch)
			}

			if (neg && n > math.MaxInt32+1) || (!neg && n > math.MaxInt32) {
				return errors.New("int32: -2147483648 <= n <= 2147483647")
			}

			if neg {
				n = -n
			}

			fieldValue.Set(reflect.ValueOf(int32(n)))
		case int64:
			var (
				n   uint64 = 0
				neg        = false
			)

			if queryData[0] == '+' {
				queryData = queryData[1:]
			} else if queryData[0] == '-' {
				neg = true
				queryData = queryData[1:]
			}

			for _, ch := range queryData {
				ch -= '0'
				if ch > 9 {
					return nil
				}
				n = n*10 + uint64(ch)
			}

			if (neg && n > math.MaxInt64+1) || (!neg && n > math.MaxInt64) {
				return errors.New("int64: -9223372036854775808 <= n <= 9223372036854775807")
			}

			if neg {
				n = -n
			}

			fieldValue.Set(reflect.ValueOf(int64(n)))

		default:
			return errors.New("unknown type in 'query' bind")
		}
	}

	return nil
}

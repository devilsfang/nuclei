// Taken from https://github.com/spf13/cast.

package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/devilsfang/nuclei/v3/pkg/model/types/severity"
)

// JSONScalarToString converts an interface coming from json to string
// Inspired from: https://github.com/cli/cli/blob/09b09810dd812e3ede54b59ad9d6912b946ac6c5/pkg/export/template.go#L72
func JSONScalarToString(input interface{}) (string, error) {
	switch tt := input.(type) {
	case string:
		return ToString(tt), nil
	case float64:
		return ToString(tt), nil
	case nil:
		return ToString(tt), nil
	case bool:
		return ToString(tt), nil
	default:
		return "", fmt.Errorf("cannot convert type to string: %v", tt)
	}
}

// ToString converts an interface to string in a quick way
func ToString(data interface{}) string {
	switch s := data.(type) {
	case nil:
		return ""
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(s, 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case []byte:
		return string(s)
	case severity.Holder:
		return s.Severity.String()
	case severity.Severity:
		return s.String()
	case fmt.Stringer:
		return s.String()
	case error:
		return s.Error()
	default:
		return fmt.Sprintf("%v", data)
	}
}

// ToStringNSlice converts an interface to string in a quick way or to a slice with strings
// if the input is a slice of interfaces.
func ToStringNSlice(data interface{}) interface{} {
	switch s := data.(type) {
	case []interface{}:
		var a []string
		for _, v := range s {
			a = append(a, ToString(v))
		}
		return a
	default:
		return ToString(data)
	}
}

func ToHexOrString(data interface{}) string {
	switch s := data.(type) {
	case string:
		if govalidator.IsASCII(s) {
			return s
		}
		return hex.Dump([]byte(s))
	case []byte:
		return hex.Dump(s)
	default:
		return fmt.Sprintf("%v", data)
	}
}

// ToStringSlice casts an interface to a []string type.
func ToStringSlice(i interface{}) []string {
	var a []string

	switch v := i.(type) {
	case []interface{}:
		for _, u := range v {
			a = append(a, ToString(u))
		}
		return a
	case []string:
		return v
	case string:
		return strings.Fields(v)
	default:
		return nil
	}
}

// ToByteSlice casts an interface to a []byte type.
func ToByteSlice(i interface{}) []byte {
	switch v := i.(type) {
	case []byte:
		return v
	case []string:
		return []byte(strings.Join(v, ""))
	case string:
		return []byte(v)
	case []interface{}:
		var buff bytes.Buffer
		for _, u := range v {
			buff.WriteString(ToString(u))
		}
		return buff.Bytes()
	default:
		strValue := ToString(i)
		return []byte(strValue)
	}
}

// ToStringMap casts an interface to a map[string]interface{} type.
func ToStringMap(i interface{}) map[string]interface{} {
	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = val
		}
		return m
	case map[string]interface{}:
		return v
	default:
		return nil
	}
}

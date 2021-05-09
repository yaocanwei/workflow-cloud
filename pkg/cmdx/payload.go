package cmdx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cast"
)

// Payload 数据存储
type Payload struct {
	payload map[string]interface{}
}

type errKeyNotFound struct {
	key string
}

func (e *errKeyNotFound) Error() string {
	return fmt.Sprintf("cmdx %q does not exist", e.key)
}

// GetString 获取字符串
func (p Payload) GetString(key string) (string, error) {
	v, ok := p.payload[key]
	if !ok {
		return "", &errKeyNotFound{key}
	}
	return cast.ToStringE(v)
}

func toInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case json.Number:
		val, err := v.Int64()
		if err != nil {
			return 0, err
		}
		return int(val), nil
	default:
		return cast.ToIntE(v)
	}
}

// GetInt 返回整型值
func (p Payload) GetInt(key string) (int, error) {
	v, ok := p.payload[key]
	if !ok {
		return 0, &errKeyNotFound{key}
	}
	return toInt(v)
}

// GetBool 返回布尔值
func (p Payload) GetBool(key string) (bool, error) {
	v, ok := p.payload[key]
	if !ok {
		return false, &errKeyNotFound{key}
	}
	return cast.ToBoolE(v)
}

// GetStringMap 返回 map string
func (p Payload) GetStringMap(key string) (map[string]interface{}, error) {
	v, ok := p.payload[key]
	if !ok {
		return nil, &errKeyNotFound{key}
	}
	return cast.ToStringMapE(v)
}

package lodash

import (
"encoding/json"
"errors"
"fmt"
"reflect"
)

type lodash struct {
	input interface{}
	err error
}

func Chain (input interface{}) *lodash {
	l :=  lodash{}
	l.input = input
	return &l
}

func (l *lodash) Value (output interface{}) error {
	if l.err != nil {
		return l.err
	}
	// 简单类型和struct可以直接反射set
	inputKind := reflect.ValueOf(l.input).Kind().String()
	switch inputKind {
	case`array`, `map`, `slice`:
		outputJson, err := json.Marshal(l.input)
		if err != nil {
			return errors.New(fmt.Sprintf(`output format error: %s`, err.Error()))
		}
		err = json.Unmarshal(outputJson, output)
		if err != nil {
			return errors.New(fmt.Sprintf(`output format error: %s`, err.Error()))
		}
	default:
		reflect.ValueOf(output).Elem().Set(reflect.ValueOf(l.input))
	}
	return nil
}

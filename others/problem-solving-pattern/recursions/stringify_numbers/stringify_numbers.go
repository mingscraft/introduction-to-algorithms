package stringify_numbers

import "strconv"

type Obj map[string]interface{}

func StringifyNumber(obj interface{}) interface{}  {
	switch v := obj.(type) {
	case Obj:
		for k, val := range v {
			v[k] = StringifyNumber(val)
		}
		return v
	default:
		if n, isInt := interface{}(v).(int); isInt {
			return strconv.FormatInt(int64(n), 10)
		}
		return v
	}
}
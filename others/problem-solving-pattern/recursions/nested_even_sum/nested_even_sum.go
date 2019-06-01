package nested_even_sum


type Obj map[string]interface{}

func NestedEvenSum(obj interface{}) int {
	switch v := obj.(type) {
	case Obj:
		sum := 0
		for _, val := range v {
			sum += NestedEvenSum(val)
		}
		return sum
	default:
		if n, isInt := interface{}(v).(int); isInt && n %2 == 0 {
			return n
		}
		return 0
	}
}
package collect_strings

type Obj map[string]interface{}

func CollectString(obj interface{}) []string {
	switch v := obj.(type) {
	case Obj:
		var res []string
		for _, val := range v {
			res = append(res, CollectString(val)...)
		}
		return res
	default:
		if val, ok := v.(string); ok {
			return []string{val}
		}

		return []string{}
	}
}

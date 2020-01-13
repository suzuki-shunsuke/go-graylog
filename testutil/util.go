package testutil

func ConvertIntToFloat64(data interface{}) interface{} {
	switch value := data.(type) {
	case int:
		return float64(value)
	case map[string]interface{}:
		for k, v := range value {
			value[k] = ConvertIntToFloat64(v)
		}
	case []interface{}:
		for i, v := range value {
			value[i] = ConvertIntToFloat64(v)
		}
	}
	return data
}

func ConvertIntToFloat64OfMap(data map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{}, len(data))
	for k, v := range data {
		m[k] = ConvertIntToFloat64(v)
	}
	return m
}

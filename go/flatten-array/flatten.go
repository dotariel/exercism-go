package flatten

// Flatten accepts an arbitrarily-deep nested list-like structure and returns a flattened
// structure without any nil values.
func Flatten(input interface{}) interface{} {
	return fill(make([]interface{}, 0), input)
}

func fill(list []interface{}, val interface{}) []interface{} {
	if _, ok := val.([]interface{}); !ok {
		if val != nil {
			list = append(list, val)
		}

		return list
	}

	for _, val := range val.([]interface{}) {
		list = fill(list, val)
	}

	return list
}

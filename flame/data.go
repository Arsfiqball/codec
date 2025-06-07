package flame

type Data map[string]interface{}

func (d Data) RemoveNil() Data {
	for k, v := range d {
		if v == nil {
			delete(d, k)
		}
	}

	return d
}

func (d Data) Merge(other Data) Data {
	for k, v := range other {
		d[k] = v
	}

	return d
}

func (d Data) IsEmpty() bool {
	return len(d) == 0
}

func (d Data) With(key string, value interface{}) Data {
	if _, ok := d[key]; ok && value == nil {
		delete(d, key) // Remove key if value is nil
		return d
	}

	if value == nil {
		return d // Do nothing if value is nil
	}

	d[key] = value // Set key to value

	return d
}

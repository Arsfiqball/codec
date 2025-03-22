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

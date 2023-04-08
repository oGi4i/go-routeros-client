package pointer

func To[T any](v T) *T {
	return &v
}

func ValueOf[T any](v *T) T {
	if v != nil {
		return *v
	}

	var tmp T
	return tmp
}

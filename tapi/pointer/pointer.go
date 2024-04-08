package pointer

// Ptr returns the Pointer
func Ptr[T any](v T) *T {
	return &v
}

// DeRef gets a Pointer and Dereference its value or if nil returns its null value
func DeRef[T any](p *T) T {
	if p == nil {
		return *new(T)
	}
	return *p
}

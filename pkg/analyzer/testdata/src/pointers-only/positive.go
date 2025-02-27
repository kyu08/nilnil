package examples

type User struct{}

func structPtr() (*User, error, error) {
	return nil, nil, nil // want "return both a `nil` error and an invalid value: use a sentinel error instead"
}

package test

type TestWantHandler[T any] struct {
	Code        int
	ContentType string
	Response    T
}

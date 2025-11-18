package test

type TestWantWanthandlerWithUrl[T any] struct {
	TestWantHandler[T]
	URL string
}

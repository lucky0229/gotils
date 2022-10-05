package optional

import "reflect"

type Optional[T any] struct {
	val T
}

func New[T any](v T) *Optional[T] {
	return &Optional[T]{
		val: v,
	}
}

func (o *Optional[T]) Get() T {
	return o.val
}

func (o *Optional[T]) OrElse(n T) T {
	if reflect.ValueOf(&o.val).Elem().IsZero() {
		return n
	}
	return o.val
}

func (o *Optional[T]) OrElseError(e error) (T, error) {
	if reflect.ValueOf(&o.val).Elem().IsZero() {
		return o.val, e
	}
	return o.val, nil
}

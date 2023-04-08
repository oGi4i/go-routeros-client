package query

import (
	"fmt"
)

type (
	set[T any] struct {
		param Param[T]
	}

	notSet[T any] struct {
		param Param[T]
	}

	eq[T comparable] struct {
		param Param[T]
		value T
	}

	lt[T comparable] struct {
		param Param[T]
		value T
	}

	gt[T comparable] struct {
		param Param[T]
		value T
	}

	in[T comparable] struct {
		param  Param[T]
		values []T
	}
)

func Set[T any](param Param[T]) Filter {
	return set[T]{param: param}
}

func (f set[T]) Prepare() []string {
	return []string{string(f.param)}
}

func NotSet[T any](param Param[T]) Filter {
	return notSet[T]{param: param}
}

func (f notSet[T]) Prepare() []string {
	return []string{fmt.Sprintf("-%s", f.param)}
}

func Eq[T comparable](param Param[T], value T) Filter {
	return eq[T]{param: param, value: value}
}

func (f eq[T]) Prepare() []string {
	return []string{fmt.Sprintf("%s=%v", f.param, f.value)}
}

func Lt[T comparable](param Param[T], value T) Filter {
	return lt[T]{param: param, value: value}
}

func (f lt[T]) Prepare() []string {
	return []string{fmt.Sprintf("<%s=%v", f.param, f.value)}
}

func Gt[T comparable](param Param[T], value T) Filter {
	return gt[T]{param: param, value: value}
}

func (f gt[T]) Prepare() []string {
	return []string{fmt.Sprintf(">%s=%v", f.param, f.value)}
}

func In[T comparable](param Param[T], values []T) Filter {
	return in[T]{param: param, values: values}
}

func (f in[T]) Prepare() []string {
	queries := make([]string, len(f.values))
	for i := range f.values {
		queries[i] = fmt.Sprintf("%s=%v", f.param, f.values[i])
	}

	return queries
}

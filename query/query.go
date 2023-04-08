package query

type (
	// Filter is an interface for a query filter
	Filter interface {
		Prepare() []string
	}

	// Param represent a typed query parameter name
	// T represents the parameter value type
	Param[T any] string
)

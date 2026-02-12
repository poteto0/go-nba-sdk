package parser

import "github.com/moznion/go-optional"

func toInt(v any) int {
	if v == nil {
		return 0
	}

	if i, ok := v.(int); ok {
		return i
	}

	if f, ok := v.(float64); ok {
		return int(f)
	}
	return 0
}

func toFloat(v any) float64 {
	if v == nil {
		return 0.0
	}

	if f, ok := v.(float64); ok {
		return f
	}

	if i, ok := v.(int); ok {
		return float64(i)
	}
	return 0.0
}

func toOptionalInt(v any) optional.Option[int] {
	if v == nil {
		return optional.None[int]()
	}

	if i, ok := v.(int); ok {
		return optional.Some(i)
	}

	if f, ok := v.(float64); ok {
		return optional.Some(int(f))
	}
	return optional.None[int]()
}

func toOptionalFloat(v any) optional.Option[float64] {
	if v == nil {
		return optional.None[float64]()
	}

	if f, ok := v.(float64); ok {
		return optional.Some(f)
	}

	if i, ok := v.(int); ok {
		return optional.Some(float64(i))
	}
	return optional.None[float64]()
}

func toOptionalString(v any) optional.Option[string] {
	if v == nil {
		return optional.None[string]()
	}

	if s, ok := v.(string); ok {
		return optional.Some(s)
	}
	return optional.None[string]()
}

package parser

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

func ptr[T any](v T) *T {
	return &v
}

func toPtrInt(v any) *int {
	if v == nil {
		return nil
	}

	if i, ok := v.(int); ok {
		return ptr(i)
	}

	if f, ok := v.(float64); ok {
		return ptr(int(f))
	}
	return nil
}

func toPtrFloat(v any) *float64 {
	if v == nil {
		return nil
	}

	if f, ok := v.(float64); ok {
		return ptr(f)
	}

	if i, ok := v.(int); ok {
		return ptr(float64(i))
	}
	return nil
}

func toPtrString(v any) *string {
	if v == nil {
		return nil
	}

	if s, ok := v.(string); ok {
		return ptr(s)
	}
	return nil
}

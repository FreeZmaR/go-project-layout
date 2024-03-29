package utils

func OneOf[T comparable](needle T, haystack ...T) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}

func WithPtr[T any](value T) *T {
	return &value
}

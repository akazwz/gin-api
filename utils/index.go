package utils

func ArrayIncludes(arr []string, value string) bool {
	set := make(map[string]struct{})

	for _, value := range arr {
		set[value] = struct{}{}
	}

	_, ok := set[value]
	return ok
}

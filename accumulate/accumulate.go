package accumulate

func Accumulate(in []string, operation func(string) string) []string {
	result := make([]string, len(in))
	for i, element := range in {
		result[i] = operation(element)
	}
	return result
}

// package accumulate provides function Accumulate
package accumulate

// Accumulate can apply various operations to the elements of an input []string
func Accumulate(input []string, operation func(string) string) []string {
	output := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = operation(input[i])
	}
	return output
}

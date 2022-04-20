package funcmap

import "fmt"

// RunSliceMapDemo executes the `Map Int to String Representation Demo`
func RunSliceMapDemo() {
	fmt.Println("------Map Int to String Representation Demo------")

	int2word := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		0: "zero",
	}

	nums := []int{8, 6, 7, 5, 3, 0, 9}

	words := Map[int, string](nums, func(e int) string { return int2word[e] })

	fmt.Println(words)
	fmt.Println("")
}

// Map applies a "mapper" function to each element in slice unmapped of type T1 and returns a slice of type T2 with the in-place results of calling the "mapper" on each element
func Map[T1 any, T2 any](unmapped []T1, mapper func(T1) T2) []T2 {
	var mapped = make([]T2, len(unmapped))
	for i, e := range unmapped {
		mapped[i] = mapper(e)
	}
	return mapped
}

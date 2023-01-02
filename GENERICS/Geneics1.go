package main

func Reverse[T any](inpSlice []T) []T {
	sliceLength := len(inpSlice)
	revSlice := make([]T, sliceLength)

	for i, element := range inpSlice {
		revSlice[sliceLength-i-1] = element
	}

}

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	// 1. strings
	str := "hello there friends!"
	fmt.Println(strings.Contains(str, "hello"))
	fmt.Println(strings.ReplaceAll(str, "hello", "hi"))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Index(str, "ll"))

	// the original value is unchanged
	fmt.Println("Original str is: ", str)

	fmt.Println(strings.Repeat("***", 16))

	// 2. sort
	ages := []int{ 45, 23, 67, 2, 70 }
	fmt.Println("before sort", ages)
	sort.Ints(ages)
	fmt.Println("after sort", ages)

	// the original value is changed
	fmt.Println("Original slice is: ", ages)

	index := sort.SearchInts(ages, 2)
	// 如果找不到，返回值为len(ages)
	indexTwo := sort.SearchInts(ages, 100)
	fmt.Println("Search", index, indexTwo, len(ages))

	names := []string{ "yoshi", "mario", "peach", "bowser", "luigi" }
	// alphabetical order
	sort.Strings(names)
	indexThree := sort.SearchStrings(names, "bowser")
	fmt.Println("sort", names, indexThree)
}
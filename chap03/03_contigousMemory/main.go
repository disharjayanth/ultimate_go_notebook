package main

import "fmt"

func inspectSlice(slice []string) {
	fmt.Printf("Length[%v]\tCapacity[%v]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}

func main() {
	str := [5]string{"Hello", "Hey", "Hi", "Yo!", "Ayyy"}
	for i, v := range str {
		fmt.Printf("Value[%v]\tAddress[%v]\tIndexAddress[%v]\n", v, &v, &str[i])
	}

	slice := []int{2, 4, 5, 6}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", len(slice))

	fmt.Println("Length and Capacity of slice;")

	sliceOfStrings := make([]string, 5, 8)
	sliceOfStrings[0] = "Apple"
	sliceOfStrings[1] = "Banana"
	sliceOfStrings[2] = "Pinneapple"
	sliceOfStrings[3] = "Strawberry"
	sliceOfStrings[4] = "Watermelon"
	sliceOfStrings2 := sliceOfStrings[2:4:4]
	// sliceOfStrings2[0] = "CHANGED!!!!"
	sliceOfStrings2 = append(sliceOfStrings2, "CHANGED")

	fmt.Printf("Address of sliceOfStrings in main func: %p\n", &sliceOfStrings)

	inspectSlice(sliceOfStrings)
	inspectSlice(sliceOfStrings2)

	// var data []string

	// for record := 1; record <= 102401; record++ {
	// 	data = append(data, fmt.Sprintf("Record: %d", record))
	// }

	// fmt.Printf("Length:[%v] Capacity: [%d]\n", len(data), cap(data))

	someStrings := []string{"Hey", "Hello", "Hi", "Yo"}
	someStrings2 := make([]string, len(someStrings))

	copy(someStrings2, someStrings)
	fmt.Printf("someStrings slice Type: [%T]\t Length: [%d]\t Capacity: [%d]\n", someStrings, len(someStrings), cap(someStrings))
	fmt.Printf("someStrings2 slice Type: [%T]\t Length: [%d]\t Capacity: [%d]\n", someStrings2, len(someStrings2), cap(someStrings2))

	inspectSlice(someStrings)
	inspectSlice(someStrings2)
}

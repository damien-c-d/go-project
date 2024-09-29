package features

import (
	"embed"
	"encoding/json"
	"fmt"
)

var Files embed.FS

func DemonstrateGenerics() {
	demonstrateWithoutGenerics()
	demonstrateWithGenerics()
	demonstrateGenericsWithAny()
	demonstrateGenericsWithJson()
	demonstrateGenericsWithStructs()
}

func demonstrateWithoutGenerics() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of intSlice:", sumIntSlice(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Sum of float32Slice:", sumFloat32Slice(float32Slice))

	var float64Slice = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Sum of float64Slice:", sumFloat64Slice(float64Slice))
}

func sumIntSlice(intSlice []int) int {
	var sum int
	for _, value := range intSlice {
		sum += value
	}
	return sum
}

func sumFloat32Slice(float32Slice []float32) float32 {
	var sum float32
	for _, value := range float32Slice {
		sum += value
	}
	return sum
}

func sumFloat64Slice(float64Slice []float64) float64 {
	var sum float64
	for _, value := range float64Slice {
		sum += value
	}
	return sum
}

func demonstrateWithGenerics() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of intSlice:", sumSlice(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Sum of float32Slice:", sumSlice(float32Slice))

	var float64Slice = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Sum of float64Slice:", sumSlice(float64Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}

func demonstrateGenericsWithAny() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println("Is intSlice empty:", isEmpty(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Is float32Slice empty:", isEmpty(float32Slice))

	var float64Slice = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Is float64Slice empty:", isEmpty(float64Slice))
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func demonstrateGenericsWithJson() {
	var contacts []contactInfo = loadJSON[contactInfo]("files/contacts.json")
	fmt.Printf("\nLoaded contacts: %v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("files/purchases.json")
	fmt.Printf("\nLoaded purchases: %v", purchases)
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float64
	Amount int
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := Files.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

func demonstrateGenericsWithStructs() {

}

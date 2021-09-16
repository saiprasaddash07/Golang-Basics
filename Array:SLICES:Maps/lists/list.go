package lists

import "fmt"

type Product struct {
	title string
	id    string
	price int
}

func createProduct(title string, id string, price int) *Product {
	prod := Product{title, id, price}
	return &prod
}

func main() {
	arr := []Product{}
	var product1 Product = *createProduct("x", "y", 10)
	arr = append(arr, product1)
	var product2 Product = *createProduct("a", "b", 20)
	arr = append(arr, product2)
	fmt.Println(arr)

	prices := []float32{10.99, 9.78}
	// prices := []float32{}
	fmt.Println(prices[1])
	// prices[2] = 9.8 // error
	updatedPrices := append(prices, 6.78, 8.9, 4.2, 5.6, 7.8, 1.2)
	fmt.Println(updatedPrices)
	prices = append(prices, 8, 24)

	hello1 := []int{1, 2}
	hello2 := []int{4, 5}
	newh := append(hello1, hello2...)
	fmt.Println(newh)
}

// func main() {
// 	// var arr [5]string = [5]string{"jehjked"}
// 	prices := [5]float32{2.4, 10.5, 6.3, 5.8, 9.2}
// 	fmt.Println(prices)
// 	fmt.Println(prices[1])
// 	prices[0] = 6.8

// 	featuredPrices := prices[1:5]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(featuredPrices),cap(featuredPrices))
// 	featuredPrices = prices[:4]
// 	fmt.Println(featuredPrices)
// 	featuredPrices = prices[:1]
// 	fmt.Println(featuredPrices)
// 	featuredPrices = prices[2:]
// 	fmt.Println(featuredPrices)
// 	featuredPrices = prices[1:]
// 	fmt.Println(featuredPrices)
// 	highprices := featuredPrices[:1]
// 	fmt.Println(highprices)
// 	// slices also modify the original array when modified
// }

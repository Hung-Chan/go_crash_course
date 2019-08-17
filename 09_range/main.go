package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	for i, id := range ids {
		fmt.Printf("%d -ID : %d\n", i, id)
	}

	// 不使用索引
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	emails := map[string]string{"Bob": "bob@gmail.com", "Fa": "fa@gmail.com"}

	for key, val := range emails {
		fmt.Printf("%s: %s\n", key, val)
	}

	for key := range emails {
		fmt.Println("Name", key)
	}

}

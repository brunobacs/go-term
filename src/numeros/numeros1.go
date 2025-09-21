package main

import "fmt"

func main() {

	//parte 1 = exibir os divisieis por 3
	fmt.Print("============Parte 1============")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}

	}

	fmt.Print("============Parte 2============")
	//parte 2 = exibir pin em multiplo de 3, pan em multiplo de 5
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin Pan")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Pin")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Pan")
			continue
		}

		fmt.Println(i)
	}

}

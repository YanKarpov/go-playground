package main

import "fmt"

func main() {
	var n int
	fmt.Println("Сгенерировать последовательность Фибоначчи до какого числа?")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if n < 1 {
		fmt.Println("Ошибка:", n, "слишком маленькое число — последовательность не может быть сгенерирована.")
		return
	} else if n > int(^uint(0)>>1) {
		fmt.Println("Ошибка:", n, "слишком большое число! Попробуйте ввести число поменьше.")
		return
	}

	fmt.Println("Последовательность Фибоначчи до", n)
	fmt.Print("0, 1, ")

	p := 0
	q := 1
	for p+q <= n {
		if p < q {
			p += q
			fmt.Printf("%d", p)
		} else {
			q += p
			fmt.Printf("%d", q)
		}

		if p+q <= n {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
}
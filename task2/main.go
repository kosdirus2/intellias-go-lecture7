package main

import (
	"strconv"
	"strings"
)

/*
На вхід подано стрінг з цілими числа, котри розділені пробілами.
Треба повернути найбільше та найменше число.
Наприклад:
input := "1 2 3 4 5" // повертає "5 1"
input := "1 9 3 4 -5" // повертає "9 -5"
Уточнення:
1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
2. В стрінгі завжди буде принаймні одне число.
3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
число). Найбільше число має бути першим.
4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
*/

func main() {
	input := "1 9 3 4 -5"
	var (
		result   string
		min, max int32
	)

	inputSlice := strings.Split(input, " ")
	if len(inputSlice) == 1 {
		result = inputSlice[0]
		return
	}

	for i, v := range inputSlice {
		valueInt, _ := strconv.Atoi(v)
		valueInt32 := int32(valueInt)
		if i == 0 {
			min = valueInt32
			max = valueInt32
			continue
		}
		switch {
		case min > valueInt32:
			min = valueInt32
		case max < valueInt32:
			max = valueInt32
		}
	}

	result = strconv.Itoa(int(max)) + " " + strconv.Itoa(int(min))
	_ = result
}

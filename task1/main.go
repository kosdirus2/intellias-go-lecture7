package main

/*
Прибрати всі дублікати з слайсу int.
Приклад даних на вхід: [3, 4, 4, 3, 6, 3]
видаляємо 3 по індексу 0
видаляємо 4 по індексу 1
видаляємо 3 по індексу 3
Правильний результат: [3, 4, 6]
Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
*/

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	mapWithUniqueNumbers := make(map[int]int)

	for _, value := range arr {
		if _, ok := mapWithUniqueNumbers[value]; ok {
			continue
		}
		mapWithUniqueNumbers[value] = 0
		result = append(result, value)
	}

}

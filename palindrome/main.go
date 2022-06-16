package main

func main() {
	slice := []int{1, 6, 5, 4, 4, 5, 6, 1}
	isPalindrome(slice)
}

func isPalindrome(slice []int) bool {
	i := 0 // присваиваем 0
	j := len(slice) - 1
	for i < len(slice)/2 && j > len(slice)/2 { // цикл длиной половина слайса слева и справа
		if slice[i] != slice[j] { // если элементы не равны
			return false // возвращаем ложь, при возврате вся функция прерывается
		}
		i++ // увеличиваем i на  единицу
		j-- // уменьшаем j на единицу и повторяем цикл
	}

	return true // если все циклы истина, возвращаем true (т. е. у нас палиндром)
}

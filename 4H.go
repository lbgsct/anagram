package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(filePath string) {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1) 
	parts := strings.Split(line1, " ")

	// g и S
	g, _ := strconv.Atoi(parts[0])
	SLength, _ := strconv.Atoi(parts[1])

	// слово W
	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2) 
	W := strings.Split(line2, "")

	// карта для слова
	windowmap := make(map[string]int)
	for i := 0; i < len(W); i++ {
		windowmap[W[i]]++
	}

	// последовательность S
	line3, _ := reader.ReadString('\n')
	line3 = strings.TrimSpace(line3) // Удаляем лишние пробелы и символы новой строки
	S := strings.Split(line3, "")

	// Инициализация переменных для окна
	beg := 0
	end := 0
	count := 0

	// в 4 буквы последовательность 
	nowWindow := make(map[string]int)

	for end < SLength {
		// добавляем символ
		nowWindow[S[end]]++

		// Если размер окна больше допустимого, уменьшаем окно
		if end-beg+1 > g {
			nowWindow[S[beg]]--
			if nowWindow[S[beg]] == 0 {
				delete(nowWindow, S[beg])
			}
			beg++
		}

		// Проверяем, содержит ли окно все символы из W
		if end-beg+1 == g && containsAllChars(nowWindow, windowmap) {
			count++
		}
		end++
	}

	// Вывод результата - количество вхождений
	fmt.Println(count)
}

// Проверяем, содержит ли текущий словарь все символы целевого словаря
func containsAllChars(windowMap map[string]int, targetMap map[string]int) bool {
	for char, count := range targetMap {
		if windowMap[char] < count {
			return false
		}
	}
	return true
}

func main() {
	readLine("/home/sergey/practice/yandex/input.txt")
}

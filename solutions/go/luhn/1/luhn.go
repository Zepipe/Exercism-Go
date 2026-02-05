package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	// Удаляем ВСЕ пробелы
	cleanId := strings.ReplaceAll(id, " ", "")
	
	// Проверка длины (минимум 2 символа)
	if len(cleanId) <= 1 {
		return false
	}

	// Проверка, что все символы - цифры
	for _, r := range cleanId {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	sum := 0
	// Обрабатываем справа налево (для определения четности позиции)
	shouldDouble := len(cleanId)%2 == 0 // Четность определяем с конца

	for _, r := range cleanId {
		digit := int(r - '0') // Преобразуем руну в число

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble // Переключаем флаг для следующей цифры
	}

	return sum%10 == 0
}
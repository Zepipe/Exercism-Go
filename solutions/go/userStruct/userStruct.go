package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User) SecretIdentity() string {
	if u.FirstName == "" || u.LastName == "" {
		return ""
	}
	
	// Берем первую букву имени и фамилии
	firstLetter := string([]rune(u.FirstName)[0])
	lastLetter := string([]rune(u.LastName)[0])
	
	// Генерируем случайное число от 1 до 100
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100) + 1
	
	return fmt.Sprintf("%s%s%d", firstLetter, lastLetter, randomNum)
}

func (u User) Age() int {
	return time.Now().Year() - u.BirthYear
}

func (u *User) AddFavoriteLanguage(language string) error {
	if language == "" {
		return errors.New("empty language name")
	}
	
	for _, lang := range u.FavoriteLanguages {
		if lang == language {
			return errors.New("duplicate")
		}
	}
	
	u.FavoriteLanguages = append(u.FavoriteLanguages, language)
	return nil
}

func (u *User) RemoveFavoriteLanguage(language string) error {
	if len(u.FavoriteLanguages) == 0 || language == "" {
		return errors.New("not found")
	}
	
	for i, lang := range u.FavoriteLanguages {
		if lang == language {
			// Удаляем элемент слайса
			u.FavoriteLanguages = append(u.FavoriteLanguages[:i], u.FavoriteLanguages[i+1:]...)
			return nil
		}
	}
	
	return errors.New("not found")
}

func (u User) IsProgrammingLanguageFavorite(language string) bool {
	if language == "" || len(u.FavoriteLanguages) == 0 {
		return false
	}
	
	for _, lang := range u.FavoriteLanguages {
		if language == lang {
			return true
		}
	}
	
	return false
}

func (u User) RandomFavoriteLanguage() (string, error) {
	if len(u.FavoriteLanguages) == 0 {
		return "", errors.New("no favorite languages")
	}
	
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(u.FavoriteLanguages))
	return u.FavoriteLanguages[randomIndex], nil
}

func (u User) GenerateProfile() string {
	languages := strings.Join(u.FavoriteLanguages, ", ")
	if languages == "" {
		languages = "нет"
	}
	return fmt.Sprintf("Имя: %s.\nФамилия: %s.\nВозраст: %d.\nСписок любимых языков программирования: [%s].", 
		u.FirstName, u.LastName, u.Age(), languages)
}

func (u *User) UpdateName(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("empty data")
	}
	
	// Проверяем первый символ на заглавную букву с учетом unicode
	firstRunes := []rune(firstName)
	lastRunes := []rune(lastName)
	
	if len(firstRunes) == 0 || len(lastRunes) == 0 {
		return errors.New("invalid data")
	}
	
	if !unicode.IsUpper(firstRunes[0]) || !unicode.IsUpper(lastRunes[0]) {
		return errors.New("invalid data")
	}
	
	u.FirstName = firstName
	u.LastName = lastName
	
	return nil
}

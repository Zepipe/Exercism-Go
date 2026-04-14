package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
    "strings"
)

// User представляет пользователя
type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

// Address представляет адрес пользователя
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Product представляет продукт в корзине
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}


func printInfo(user User) {
    result := "" // result string with four sentences  
    var hasElectronic bool
    bucketOfGoods := []string{}
    totalBucketSum := 0
    
     result += fmt.Sprintf("Покупатель %s. Телефон: %s. Адрес: г. %s, %s.\n", user.Name, user.Phone, user.Address.City, user.Address.Street) // first sentence
    
    for _, cartItem := range user.Cart {
        if cartItem.Product.Category == "Электроника" { // check for second sentence
            hasElectronic = true
        }
        
        if cartItem.Product.Price >= 10000 { // products for third sentence
            bucketOfGoods = append(bucketOfGoods, cartItem.Product.Name)
        }
        
        totalBucketSum += cartItem.Quantity * cartItem.Product.Price // total sum for fourth sentence
    }
    
    if hasElectronic { // second sentence
        result += "Пользователь является покупателем электроники.\n"
    } else {
        result += "Пользователь не является покупателем электроники.\n"
    }
    
    if len(bucketOfGoods) == 0 { // third sentence
        result += "Товары в корзине, где цена 10000 и более: отсутствуют.\n"
    } else {
        result += fmt.Sprintf("Товары в корзине, где цена 10000 и более: %s.\n", strings.Join(bucketOfGoods, ", "))
    }
    
    result += fmt.Sprintf("Общая сумма покупки: %d руб.", totalBucketSum) // fourth sentence
    
    fmt.Println(result)
}

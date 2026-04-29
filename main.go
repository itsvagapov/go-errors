package main

import (
	// "errors"
	"fmt"
	"carservice/carservice"
)

// type Product struct {
// Name  string
// Price int // в копейках
// Stock int // сколько на складе
// }
//
// Склад магазина
// var warehouse = []Product{
// {Name: "Ноутбук", Price: 8999900, Stock: 5},
// {Name: "Мышь", Price: 249900, Stock: 50},
// {Name: "Клавиатура", Price: 499900, Stock: 30},
// }
//
// FindProduct ищет товар по имени и возвращает его
// func FindProduct(name string) (*Product, error) {
// for _, p := range warehouse {
// if p.Name == name {
// return &p, nil
// }
// }
// Что вернуть, если товар не найден?
// return nil, errors.New("товар не найден!") // пустой продукт — это плохо
// }
//
// Buy оформляет покупку: уменьшает остаток на складе
// func Buy(name string, quantity int) error {
// Что если товара нет?
// if quantity <= 0 {
// return errors.New("некорректное количество!")
// }
// Что если quantity больше, чем есть на складе?
// product, err := FindProduct(name)
// if err != nil {
// return err
// }
// Что если quantity отрицательный?
// if product.Stock < quantity {
// return errors.New("недостаточно товара на складе!")
// }
// product.Stock -= quantity
//
// return nil
// }
//
func main() {
// Покупатель хочет купить 3 ноутбука
// product, err := FindProduct("Ноутбук")
// if err != nil {
// fmt.Println("Ошибка:", err)
// } else {
//  fmt.Printf("Найден товар: %s, цена: %d копеек\n", product.Name, product.Price)
// }
//
//
// err = Buy("Ноутбук", 4)
// if err != nil {
// fmt.Println("Ошибка:", err)
// } else {
// fmt.Println("Покупка успешно завершена!")
// }
//
// А что если покупатель ищет товар, которого нет?
// product2, err := FindProduct("Мышь")
// if err != nil {
// fmt.Println("Ошибка:", err)
// } else {
//  fmt.Printf("Найден товар: %s, цена: %d копеек\n", product2.Name, product2.Price)
// }
//
//
//   err = Buy("Мышь", 2)
// if err != nil {
// fmt.Println("Ошибка:", err)
// } else {
// fmt.Println("Покупка успешно завершена!")
// }
// Что напечатается?



manager := &carservice.AppointmentManager{}
 
	fmt.Println("=== Создание записей ===")
 
	err := manager.CreateAppointment("А123БВ77", "Ахмед Сулейманов", "ТО")
	fmt.Println("Создана запись А123БВ77:", err)
 
	err = manager.CreateAppointment("Б456ГД99", "Муслим Анзоров", "Ремонт")
	fmt.Println("Создана запись Б456ГД99:", err)
 
	err = manager.CreateAppointment("", "Ибрах1им Тсуев", "ТО")
	fmt.Println("Пустой номер:", err)
 

	err = manager.CreateAppointment("В111ВВ77", "", "ТО")
	fmt.Println("Пустое имя:", err)
 

	err = manager.CreateAppointment("Г222ГГ77", "Мухаммад Вагапов", "Покраска")
	fmt.Println("Неизвестная услуга:", err)
 
	
	err = manager.CreateAppointment("А123БВ77", "Расул Бахарчиев", "Диагностика")
	fmt.Println("Дубликат:", err)
 
	fmt.Println("\n=== Поиск записей ===")
 
	appt, err := manager.FindAppointment("Б456ГД99")
	fmt.Println("Найдено:", appt, err)
 
	appt, err = manager.FindAppointment("Х000ХХ00")
	fmt.Println("Не найдено:", appt, err)
 
	appt, err = manager.FindAppointment("")
	fmt.Println("Пустой номер:", appt, err)
 
	fmt.Println("\n=== Отмена записей ===")
 
	err = manager.CancelAppointment("А123БВ77")
	fmt.Println("Отменена А123БВ77:", err)
 
	err = manager.CancelAppointment("А123БВ77")
	fmt.Println("Повторная отмена А123БВ77:", err)
 
	err = manager.CancelAppointment("")
	fmt.Println("Пустой номер:", err)
}
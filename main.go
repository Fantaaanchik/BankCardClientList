package main

import (
	"errors"
	"fmt"
	"log"
	"tests/funcs"
	"tests/models"
)

func main() {
	fmt.Println("Запуск...")
	var (
		cliNum int
		cl     models.Client
		card   models.Card
		err    error
	)

	fmt.Println("Введите количество клиентов!")
	_, err = fmt.Scan(&cliNum)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("неверный тип данных количества клиентов!")
		return
	}
	for i := 1; i <= cliNum; i++ {
		fmt.Println("Введите номер карты!")
		_, err = fmt.Scan(&card.Number)
		if err != nil {
			fmt.Print(errors.New("неверный номер карты"))
			return
		}
		fmt.Println(`Введите данные клиента! "ФИО и Возраст и Адресс и телефон!"`)
		_, err = fmt.Scan(&card.FullName, &card.Age, &card.Address, &card.Phone)
		if err != nil {
			fmt.Println(errors.New("неверный ввод данных, вы должны ввести ФИО и Возраст и Адресс и телефон"))
			return
		}

		if card.Number[:4] != models.HumoPrefix {
			card.Bank = "Другой банк"
		} else {
			card.Bank = "Хумо банк"
		}
		fmt.Println("Введите срок карты!")
		_, err = fmt.Scan(&card.ExpDate)
		if err != nil {
			fmt.Println(errors.New("не верный формат срока карты"))
			return
		}
		err = funcs.CheckDate(card.ExpDate)
		if err != nil {
			fmt.Println(errors.New("срок карты истек"))
			return
		}
		cl.Card = append(cl.Card, card)
		fmt.Println(cl)
	}

}

// RewriteTjToRu - replaces Tajik letters to Russian (қ = > к)
//func RewriteTjToRu(text string) string {
//	text = strings.Replace(text, "Ғ", "Г", -1)
//	text = strings.Replace(text, "ғ", "г", -1)
//	text = strings.Replace(text, "Ӣ", "И", -1)
//	text = strings.Replace(text, "ӣ", "и", -1)
//	text = strings.Replace(text, "Қ", "К", -1)
//	text = strings.Replace(text, "қ", "к", -1)
//	text = strings.Replace(text, "Ӯ", "У", -1)
//	text = strings.Replace(text, "ӯ", "у", -1)
//	text = strings.Replace(text, "Ҳ", "Х", -1)
//	text = strings.Replace(text, "ҳ", "х", -1)
//	text = strings.Replace(text, "Ҷ", "Ч", -1)
//	text = strings.Replace(text, "ҷ", "ч", -1)
//	return text
//}

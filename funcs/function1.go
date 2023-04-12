package funcs

import (
	"errors"
	"time"
)

// CheckDate - принимает текст в формате dd.mm.yyyy и проверяет срок годности
func CheckDate(date string) (err error) {
	expDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		return errors.New("Неверный формат времени! Подсказка: dd.mm.yyyy")
	}
	if expDate.Unix() < time.Now().Unix() {
		return errors.New("Срок карты истёк!")
	}
	return
}

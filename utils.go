package mapon

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"time"
)

// CheckError - вспомогательная функция для проверки и обработки ошибок
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// FormatDateTime форматирует дату в требуемый API формат
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

// ToJSON конвертирует структуру в JSON-строку
func ToJSON(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// FromJSON конвертирует JSON-строку в структуру
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

// AddQueryParam добавляет параметры в URL
func AddQueryParam(u *url.URL, key, value string) {
	q := u.Query()
	q.Add(key, value)
	u.RawQuery = q.Encode()
}

// IntToStr конвертирует int в string для использования в URL
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

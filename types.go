package mapon

type HistoryPoint struct {
	UnitID   int     `json:"unit_id"`
	DateTime string  `json:"datetime"`
	Mileage  float64 `json:"mileage"`
	// Добавьте другие поля, которые возвращает API
}

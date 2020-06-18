package models

type Car struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Color string `json:"color"`
}

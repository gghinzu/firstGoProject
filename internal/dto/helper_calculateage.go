package dto

import (
	"time"
)

func CalculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	if age < 0 {
		return 0
	}
	return age
}

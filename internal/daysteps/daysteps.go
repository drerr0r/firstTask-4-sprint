package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже

	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}

	countStep := parts[0]
	count, err := strconv.Atoi(countStep)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при парсинге количества")
	}

	durationTraning := parts[1]
	duration, err := time.ParseDuration(durationTraning)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при парсинге длительности")
	}
	return count, duration, nil

}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже

	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * StepLength
	distanceKm := distance / 1000
	walkingSpentcalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	result := fmt.Sprintf(
		"Количество шагов: %d. \n"+
			"Дистанция составила %.2f км. \n"+
			"Вы сожгли %.2f ккал. ",
		steps,
		distanceKm,
		walkingSpentcalories,
	)
	return result
}

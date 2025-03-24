package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, " ", 0, fmt.Errorf("invalid input format")
	}

	count := parts[0]
	countSteps, err := strconv.Atoi(count)
	if countSteps <= 0 {
		return 0, " ", 0, fmt.Errorf("negative steps error: %w", err)
	}
	if err != nil {
		return 0, " ", 0, fmt.Errorf("parsing quantity error: %w", err)
	}

	durationTraning := parts[2]
	duration, err := time.ParseDuration(durationTraning)
	if duration <= 0 {
		return 0, " ", 0, fmt.Errorf("negative duration error: %w", err)

	}
	if err != nil {
		return 0, " ", 0, fmt.Errorf("parsing duration error: %w", err)
	}

	activityType := parts[1]

	return countSteps, activityType, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже

	return (float64(steps) * lenStep) / mInKm

}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже

	if duration <= 0 {
		return 0
	}
	return distance(steps) / duration.Hours()

}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь

	speed := meanSpeed(steps, duration)

	calories := ((runningCaloriesMeanSpeedMultiplier * speed) - runningCaloriesMeanSpeedShift) * weight

	return calories

}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	speed := meanSpeed(steps, duration)

	calories := ((walkingCaloriesWeightMultiplier * weight) + (speed*speed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH

	return calories
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже

	countSteps, typeOfTraining, durationTraning, err := parseTraining(data)
	if err != nil {
		return fmt.Sprintf("Ошибка при обработке данных: %s", err.Error())
	}

	var speed float64
	var calories float64

	speed = meanSpeed(countSteps, durationTraning)

	switch typeOfTraining {
	case "Бег":
		calories = RunningSpentCalories(countSteps, weight, durationTraning)
	case "Ходьба":
		calories = WalkingSpentCalories(countSteps, weight, height, durationTraning)
	default:
		return fmt.Sprintf("Неизвестный тип тренировки: %s", typeOfTraining)
	}

	result := fmt.Sprintf("Количество шагов: %d\n"+
		"Тип тренировки: %s\n"+
		"Продолжительность: %.2f минут\n"+
		"Пройденная дистанция: %.2f км\n"+
		"Средняя скорость: %.2f км/ч\n"+
		"Сгоревшие калории: %.2f ккал\n",
		countSteps,
		typeOfTraining,
		durationTraning.Minutes(),
		distance(countSteps),
		speed,
		calories)

	return result

}

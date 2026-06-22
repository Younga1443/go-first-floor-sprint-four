package spentcalories // расчитывает потраченные калории

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// возвращает кол-во шагов, вид активности, продолжительность активности
func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("len slices is not 3")
	}
	step, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, err
	}
	timeTraining, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, err
	}

	if step <= 0 {
		fmt.Println("steps are zero")
		return 0, "", 0, fmt.Errorf("steps are zero")
	}

	if timeTraining <= 0 {
		return 0, "", 0, fmt.Errorf("timeTraining are zero")
	}

	return step, parts[1], timeTraining, nil
}

// принимает шаги и рост, возвращает дистанцию в километрах
func distance(steps int, height float64) float64 {
	if steps <= 0 {
		fmt.Println("steps are zero")
		return 0
	}

	if height <= 0 {
		fmt.Println("height are zero")
		return 0
	}

	return ((height * stepLengthCoefficient) * float64(steps)) / mInKm
}

// возвращает среднюю скорость
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		fmt.Println("steps are zero")
		return 0
	}
	if height <= 0 {
		fmt.Println("height are zero")
		return 0
	}
	if duration <= 0 {
		fmt.Println("duration are zero")
		return 0
	}

	distanceTraining := distance(steps, height)
	durationHours := duration.Hours()

	return distanceTraining / durationHours
}

// возвращает информацию о тренировке
func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeActivity, durationActivity, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if durationActivity <= 0 {
		return "", fmt.Errorf("durationActivity are zero")
	}
	if height <= 0 {
		return "", fmt.Errorf("height are zero")
	}
	if weight <= 0 {
		return "", fmt.Errorf("weight are zero")
	}

	distanceTraining := distance(steps, height)
	averageSpeed := meanSpeed(steps, height, durationActivity)

	runningSpentCalories, err := RunningSpentCalories(steps, weight, height, durationActivity)
	if err != nil {
		return "", err
	}
	walkingSpentCalories, err := WalkingSpentCalories(steps, weight, height, durationActivity)
	if err != nil {
		return "", err
	}

	switch typeActivity {
	case "Бег":
		runOutput := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, durationActivity.Hours(), distanceTraining, averageSpeed, runningSpentCalories)
		return runOutput, nil
	case "Ходьба":
		walkOutput := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, durationActivity.Hours(), distanceTraining, averageSpeed, walkingSpentCalories)
		return walkOutput, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	//return " ", fmt.Errorf("неизвестный тип тренировки")
}

// возвращает кол-во потраченных каллорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps are zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight are zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height are zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration are zero")
	}

	durationMinutes := duration.Minutes()
	averageSpeed := meanSpeed(steps, height, duration)

	return (weight * averageSpeed * durationMinutes) / minInH, nil
}

// возвращает кол-во потраченных каллорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps are zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight are zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height are zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration are zero")
	}

	averageSpeed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	return ((weight * averageSpeed * durationMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

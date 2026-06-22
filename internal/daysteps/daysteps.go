package daysteps // учет активности в течении дня

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// возвращает кол-во шагов и продолжительность прогулки
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("len slices not 2")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps are zero")
	}
	timeWalk, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, timeWalk, nil
}

// возвращает информацию о дне
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return " "
	}
	if steps <= 0 {
		fmt.Println("steps are zero")
		return " "
	}
	if weight <= 0 {
		fmt.Println("weight are zero")
		return " "
	}

	if height <= 0 {
		fmt.Println("height are zero")
		return " "
	}

	distanceWalk := (float64(steps) * stepLength) / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return " "
	}

	infoDay := fmt.Sprintf("Количество шагов: %d.\n Дистанция составила %.2f.\n Вы сожгли %.2f.", steps, distanceWalk, calories)
	return infoDay
}

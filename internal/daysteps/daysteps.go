package daysteps // учет активности в течении дня

import (
	"fmt"
	"log"
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
		log.Println(err)
		return 0, 0, err
	}
	if steps <= 0 {
		log.Println(err)
		return 0, 0, fmt.Errorf("steps are zero")
	}
	timeWalk, err := time.ParseDuration(parts[1])
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if timeWalk <= 0 {
		log.Println(err)
		return 0, 0, fmt.Errorf("timeWalk are zero")
	}

	return steps, timeWalk, nil
}

// возвращает информацию о дне
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		log.Println("steps are zero")
		fmt.Println("steps are zero")
		return ""
	}
	if weight <= 0 {
		log.Println("weight are zero")
		fmt.Println("weight are zero")
		return ""
	}

	if height <= 0 {
		log.Println("height are zero")
		fmt.Println("height are zero")
		return ""
	}
	if duration <= 0 {
		log.Println("duration are zero")
		fmt.Println("duration are zero")
		return ""
	}

	distanceWalk := (float64(steps) * stepLength) / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	if calories <= 0 {
		log.Println("calories are zero")
		fmt.Println("calories are zero")
		return ""
	}
	if distanceWalk <= 0 {
		log.Println("distanceWalk are zero")
		fmt.Println("distanceWalk are zero")
		return ""
	}

	infoDay := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceWalk, calories)
	return infoDay
}

package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

var ()

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	if data == "" {
		return 0, 0, spentcalories.ErrNoString
	}
	parts := strings.Split(data, ",")
	if len(parts) == 2 {
		steps, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, err
		}
		if steps <= 0 {
			return 0, 0, spentcalories.ErrSteps
		}
		time, err := time.ParseDuration(parts[1])
		if err != nil {
			return 0, 0, err
		}
		if time <= 0 {
			return 0, 0, spentcalories.ErrTime
		}
		return steps, time, nil
	}
	return 0, 0, spentcalories.ErrSliceLen
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	if data == "" {
		log.Println(spentcalories.ErrNoString)
		return ""
	}
	steps, time, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps > 0 {
		distance := float64(steps) * stepLength / mInKm
		calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, time)
		if err != nil {
			log.Println(err)
			return ""
		}
		result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
		return result
	}
	log.Println(spentcalories.ErrSteps)
	return ""
}

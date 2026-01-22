package spentcalories

import (
	"errors"
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

var (
	ErrParameterZero = errors.New("ошибка. переданный параметр равняется нулю")
	ErrNoString      = errors.New("ошибка. пустая строка")
	ErrUnknownType   = errors.New("ошибка. неизвестный тип тренировки")
	ErrSliceLen      = errors.New("ошибка длины полученного слайса")
	ErrSteps         = errors.New("ошибка количества шагов")
	ErrTime          = errors.New("ошибка. нулевое или отрицательно время")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	if data == "" {
		return 0, "", 0, ErrNoString
	}
	parts := strings.Split(data, ",")
	if len(parts) == 3 {
		steps, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, "", 0, err
		}
		if steps <= 0 {
			return 0, "", 0, ErrSteps
		}
		time, err := time.ParseDuration(parts[2])
		if err != nil {
			return 0, "", 0, err
		}
		if time <= 0 {
			return 0, "", 0, ErrTime
		}
		return steps, parts[1], time, nil
	}
	return 0, "", 0, ErrSliceLen
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLen := stepLengthCoefficient * height
	dist := float64(steps) * stepLen / mInKm
	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration > 0 {
		dist := distance(steps, height)
		speed := dist / duration.Hours()
		return speed
	}
	return 0
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	if data == "" {
		log.Println(ErrNoString)
		return "", ErrNoString
	}
	if (weight <= 0) || (height <= 0) {
		log.Println(ErrParameterZero)
		return "", ErrParameterZero
	}
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err, "ошибка парсинга")
		return "", err
	}

	switch {
	case activityType == "Бег":
		dist := distance(steps, height)
		avgSpeed := meanSpeed(steps, height, duration)
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", activityType, duration.Hours(), dist, avgSpeed, calories)
		return str, nil
	case activityType == "Ходьба":
		dist := distance(steps, height)
		avgSpeed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", activityType, duration.Hours(), dist, avgSpeed, calories)
		return str, nil
	default:
		log.Println(ErrUnknownType)
		return "", ErrUnknownType
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if (steps <= 0) || (weight <= 0) || (height <= 0) || (duration <= 0) {
		return 0, ErrParameterZero
	}
	avgSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	return weight * avgSpeed * minutes / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if (steps <= 0) || (weight <= 0) || (height <= 0) || (duration <= 0) {
		return 0, ErrParameterZero
	}
	avgSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	return weight * avgSpeed * minutes / minInH * walkingCaloriesCoefficient, nil
}

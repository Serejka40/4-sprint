package spentcalories

import (
	"fmt"
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

func parseTraining(data string) (int, string, time.Duration, error) {
	// Разделить строку на слайс строк.
	dataSl := strings.Split(data, ",")
	// Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов, вид активности и продолжительность.
	if len(dataSl) != 3 {
		return 0, "", 0, fmt.Errorf("ожидается строка формата '3456,Ходьба,3h00m', получено: '%s'", data)
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	steps, err := strconv.Atoi(dataSl[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов")
	}
	if steps == 0 {
		return 0, "", 0, fmt.Errorf("ошибка: 0 шагов")
	}
	// Преобразовать третий элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	duration, err := time.ParseDuration(dataSl[3])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования продолжительности")
	}
	// Если всё прошло без ошибок, верните количество шагов, продолжительность и nil (для ошибки).
	return steps, dataSl[2], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

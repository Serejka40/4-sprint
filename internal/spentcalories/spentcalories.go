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
	// Рассчитайте длину шага. Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient.
	// Соответствующая константа уже определена в пакете.
	// умножьте пройденное количество шагов на длину шага.
	// разделите полученное значение на число метров в километре (mInKm, константа определена в пакете).
	distance := (float64(steps) * (height * stepLengthCoefficient)) / float64(mInKm)
	return distance

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 {
		fmt.Println("Ошибка: продолжительность не больше 0")
		return 0
	}
	// Вычислить дистанцию с помощью distance().
	distance := distance(steps, height)
	// Вычислить и вернуть среднюю скорость. Для этого разделите дистанцию на продолжительность в часах.
	// Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	averageSpeed := distance / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	//

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 {
		return 0, fmt.Errorf("некорректное количество шагов: %d (требуется > 0)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("некорректный вес: %.1f кг (требуется > 0)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("некорректный рост: %.2f м (требуется > 0)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("некорректная продолжительность: %v (требуется > 0)", duration)
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	averageSpeed := meanSpeed(steps, height, duration)
	// Рассчитать и вернуть количество калорий.
	calories := (weight * averageSpeed * duration.Minutes()) / 60
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

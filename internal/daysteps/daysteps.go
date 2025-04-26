package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Разделить строку на слайс строк.
	dataSl := strings.Split(data, ",")
	// Проверить, чтобы длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.
	if len(dataSl) != 2 {
		return 0, 0, fmt.Errorf("ожидается строка формата '678,0h50m', получено: '%s'", data)
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	steps, err := strconv.Atoi(dataSl[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования шагов")
	}
	if steps == 0 {
		return 0, 0, fmt.Errorf("0 шагов")
	}
	// Преобразовать второй элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	duration, err := time.ParseDuration(dataSl[1])
	if err != 0 {
		return 0, 0, fmt.Errorf("ошибка преобразования продолжительности")
	}
	// Если всё прошло без ошибок, верните количество шагов, продолжительность и nil (для ошибки).
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// Получить данные о количестве шагов и продолжительности прогулки с помощью функции parsePackage().
	// В случае возникновения ошибки вывести её на экран и вернуть пустую строку.
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}
	// Проверить, чтобы количество шагов было больше 0. В противном случае вернуть пустую строку.
	if steps <= 0 {
		fmt.Println("Ошибка: шагов не больше 0")
		return ""
	}
	// Вычислить дистанцию в метрах. Дистанция равна произведению количества шагов на длину шага.
	// Константа stepLength (длина шага) уже определена в коде.
	distance := float64(steps) * stepLength / float64(mInKm)
	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
}

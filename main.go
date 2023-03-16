package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	user_ie int = 0
	user_sn int = 0
	user_tf int = 0
	user_jp int = 0

	user_i = "I"
	user_n = "N"
	user_f = "F"
	user_p = "P"

	user_nnnn = ""
	user_full = "full description after test"
	info      = "Для ответа пишите а или б"

	question_array, err = readLines("questions.txt")
	answers_map         = make(map[int]bool, 70)
	types_map, errs     = mapFilling("types.txt")
)

func mapFilling(path string) (map[string]string, error) {
	file, err := os.Open(path)
	err_handler(err)
	defer file.Close()

	var temp = make(map[string]string, 16)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp_arr := strings.Split(scanner.Text(), "_")
		temp[temp_arr[0]] = temp_arr[1]
	}
	return temp, scanner.Err()
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	err_handler(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func err_handler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func calculations(mapp map[int]bool) {
	for i := range mapp {
		if (i%7 == 0) && mapp[i] {
			user_ie += 1
		}
		if (i%7 == 1 || i%7 == 2) && mapp[i] {
			user_sn += 1
		}
		if (i%7 == 3 || i%7 == 4) && mapp[i] {
			user_tf += 1
		}
		if (i%7 == 5 || i%7 == 6) && mapp[i] {
			user_jp += 1
		}
	}
	if user_ie > 5 {
		user_i = "E"
	} else {
		user_i = "I"
	}
	if user_sn > 10 {
		user_n = "S"
	} else {
		user_n = "N"
	}
	if user_tf > 10 {
		user_f = "T"
	} else {
		user_f = "F"
	}
	if user_jp > 10 {
		user_p = "J"
	} else {
		user_p = "P"
	}
	user_nnnn = user_i + user_n + user_f + user_p
	user_full = types_map[user_nnnn]
}

func main() {
	fmt.Println("Привет! Это Тест 8. Инструкция к тесту: Этот вопросник предназначен для определения типичных способов поведения и личностных характеристик. Он состоит из 70 утверждений (вопросов), каждое из которых имеет два варианта ответа. Вам необходимо выбрать ОДИН. Все ответы равноценны, среди них нет 'правильных' или 'неправильных'! Поэтому не нужно 'угадывать' ответ. Выберите ответ, который свойствен вашему поведению в большинстве жизненных ситуаций. Работайте последовательно, не пропуская вопросов. Отвечайте правдиво, если вы хотите узнать что-то о себе, а не о какой-то мифической личности.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите ваше имя:")
	_, err := reader.ReadString('\n')
	err_handler(err)
	fmt.Println("Введите вашу почтy:")
	_, errs := reader.ReadString('\n')
	err_handler(errs)

	for j, i := range question_array {
	exit:
		fmt.Println(j + 1)
		fmt.Println(i + " " + info)
		input, err := reader.ReadString('\n')
		err_handler(err)
		if ([]byte(input)[1] != []byte("а")[1]) && ([]byte(input)[1] != []byte("б")[1]) {
			fmt.Println("ошибка ввода")
			goto exit
		}
		answers_map[j] = (input == "а")
	}
	calculations(answers_map)
	fmt.Println("Спасибо! Тестирование окончено.", user_nnnn, user_full)
}

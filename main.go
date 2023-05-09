package main

import (
	"KKTwatcher/internal/fptr10"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func main() {
	logger.Out = os.Stdout

	numOfKkt, err := findKkt()
	if err != nil {
		logger.Println(err)
	}
	if numOfKkt == 1 {
		kktType, err := getTypeOfKkt()
		if err != nil {
			logger.Println(err)
		}
		fmt.Println(numOfKkt, kktType)
	}
}

func findKkt() (int, error) {
	var discovered int

	//поищем драйвер и ККТ Атол
	fptr_lib, err := fptr10.NewWithID("KKT1")
	if err != nil {
		return 0, err
		logger.Println("всё пропало")
	}
	if fptr_lib != "" {
		logger.Println("найден и успешно создан экземпляр драйвера Атол")
		discovered++
	} else {
		logger.Println("драйвер Атол не найден")
	}
	if discovered == 0 {
		return 0, errors.New("ничего не найдено")
	}
	return discovered, nil
}

func getTypeOfKkt() (int, error) {
	return 0, errors.New("не удалось определить тип ККТ")
}

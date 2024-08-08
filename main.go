package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputNumbers, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputNumbersTrim := strings.TrimSpace(inputNumbers)
	romanFirstNumber := ""
	romanSecondNumber := ""
	result := 0
	romanNumbersMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XII": 12, "XIV": 14, "XV": 15, "XVI": 16, "XVIII": 18, "XX": 20, "XXI": 21, "XXIV": 24, "XXV": 25, "XXVII": 27, "XXVIII": 28, "XXX": 30, "XXXII": 32, "XXXV": 35, "XXXVI": 36,
		"XL": 40, "XLII": 42, "XLV": 45, "XLVIII": 48, "XLIX": 49, "L": 50, "LIV": 54, "LVI": 56, "LX": 60, "LXIII": 63, "LXIV": 64,
		"LXX": 70, "LXXII": 72, "LXXX": 80, "LXXXI": 81, "XC": 90, "C": 100,
	}

	numbersArr := strings.Split(inputNumbersTrim, " ")

	operator, operatorCount := findOperator(inputNumbersTrim)

	numbersArr = strings.Split(strings.ReplaceAll(inputNumbersTrim, " ", ""), operator)

	if len(numbersArr) <= 1 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	firstNum, errFirstNum := strconv.Atoi(numbersArr[0])

	if errFirstNum != nil {
		var ok bool
		firstNum, ok = romanNumbersMap[numbersArr[0]]

		if ok == true {
			romanFirstNumber = findKeyByValue(firstNum, romanNumbersMap)
		} else {
			romanFirstNumber = "0"
		}
	}

	secondNum, errSecondNum := strconv.Atoi(numbersArr[1])

	if errSecondNum != nil {
		var ok bool
		secondNum, ok = romanNumbersMap[numbersArr[1]]

		if ok == true {
			romanSecondNumber = findKeyByValue(secondNum, romanNumbersMap)
		} else {
			romanSecondNumber = "0"
		}
	}

	if operatorCount > 1 || len(numbersArr) > 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	if (errFirstNum == nil && errSecondNum != nil) || (errFirstNum != nil && errSecondNum == nil) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	if romanFirstNumber == "0" || romanSecondNumber == "0" {
		panic("Выдача паники, так как в римской системе нет таких чисел.")
	}

	result = calculate(firstNum, secondNum, operator)

	if errFirstNum != nil && result < 0 {
		panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
	}

	if errFirstNum != nil && errSecondNum != nil {
		fmt.Println(romanFirstNumber, operator, romanSecondNumber, " = ", findKeyByValue(result, romanNumbersMap))
	} else {
		fmt.Println(firstNum, operator, secondNum, " = ", result)
	}

	//testCalculator(romanNumbersMap, (errFirstNum != nil && errSecondNum != nil))
}

func findOperator(input string) (string, int) {
	operatorPlus := "+"
	operatorMinus := "-"
	operatorMultiply := "*"
	operatorDivide := "/"
	currentOperator := ""
	operatorCount := 0

	arr := strings.Split(input, "")

	for i := range len(arr) {
		switch arr[i] {
		case operatorPlus:
			currentOperator = operatorPlus
			operatorCount++
		case operatorMinus:
			currentOperator = operatorMinus
			operatorCount++
		case operatorMultiply:
			currentOperator = operatorMultiply
			operatorCount++
		case operatorDivide:
			currentOperator = operatorDivide
			operatorCount++
		}
	}

	return currentOperator, operatorCount
}

func calculate(firstNumber, secondNumber int, operator string) int {
	var result int

	switch operator {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		result = firstNumber / secondNumber
	}

	return result
}

func findKeyByValue(value int, hashMap map[string]int) string {
	for i := range hashMap {
		if hashMap[i] == value {
			return i
		}
	}

	return ""
}

func testCalculator(romanNumbersMap map[string]int, isRomanNumber bool) {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if isRomanNumber {
				fmt.Print(findKeyByValue(i*j, romanNumbersMap), "\t")
			} else {
				fmt.Print(i*j, "\t")
			}
		}
		fmt.Println()
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func resultOfOperation(a, b int, sign string) int { // вывод результата действия с учётом знака

	switch sign {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}

func isRomanNumber(rom string) (int, bool) { // проверка на римское число и одновременный перевод из римского в арабское

	switch rom {
	case "I":
		return 1, true
	case "II":
		return 2, true
	case "III":
		return 3, true
	case "IV":
		return 4, true
	case "V":
		return 5, true
	case "VI":
		return 6, true
	case "VII":
		return 7, true
	case "VIII":
		return 8, true
	case "IX":
		return 9, true
	case "X":
		return 10, true
	default:
		return 0, false
	}
}

func arabicTensToRoman(art int) string { // перевод арабских десятков в римские

	switch art {
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	case 10:
		return "C"
	default:
		return ""
	}
}

func arabicUnitsToRoman(aru int) string { // перевод арабских единиц в римские

	switch aru {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	case 10:
		return "X"
	default:
		return ""
	}
}

func main() {

	fmt.Println("Ведите действие с двумя числами в одну строку через пробелы:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	splitedStr := strings.Split(str, " ")

	if len(splitedStr) != 3 {
		fmt.Println("Ошибка! (неверный ввод данных, действия происходят ровно с двумя операндами)")
		return
	}

	if splitedStr[1] != "+" && splitedStr[1] != "-" && splitedStr[1] != "*" && splitedStr[1] != "/" {
		fmt.Println("Ошибка (знак оператора может быть только символом \"+\" или \"-\" или \"*\" или \"/\")")
		return
	}

	n1, err1 := strconv.Atoi(splitedStr[0])
	n2, err2 := strconv.Atoi(splitedStr[2])

	if err1 == nil && err2 == nil { // если числа арабские то...
		if n1 < 1 || n1 > 10 || n2 < 1 || n2 > 10 {
			fmt.Println("Ошибка! (входные данные должны быть не менее 1(I) и не более 10(X))")
		} else {
			fmt.Println(resultOfOperation(n1, n2, splitedStr[1]))
		}
	} else if err1 != nil && err2 == nil || err1 == nil && err2 != nil { // если введены разные системы счисления то...
		fmt.Println("Ошибка! (нельзя использовать в одном действие разные системы счисления)")
	} else { // если числа не арабские то...
		n1r, isrom1 := isRomanNumber(splitedStr[0])
		n2r, isrom2 := isRomanNumber(splitedStr[2])
		if isrom1 && isrom2 && (n1r-n2r > 0 || splitedStr[1] != "-") {
			res := resultOfOperation(n1r, n2r, splitedStr[1])
			fmt.Println(arabicTensToRoman(res/10) + arabicUnitsToRoman(res%10))
		} else if isrom1 && isrom2 {
			fmt.Println("Ошибка! (в римской системе счисления нет отрицательных чисел)")
		} else {
			fmt.Println("Ошибка! (входные данные должны быть не менее 1(I) и не более 10(X))")
		}
	}
}

package main

import (
	"fmt"
	"strconv"
)

func poluchenie() (first string, oper string, second string) {
	//read := bufio.NewReader(os.Stdin)
	//first, _ = read.ReadString('\n')
	//fmt.Println(first)
	var test, test1 string
	fmt.Scanf("%s%s%s%s%s", &first, &oper, &second, &test, &test1)
	if test1 != "" || test != "" {
		panic("Ошибка Много символов")
	}
	if first == "" || oper == "" || second == "" {
		panic(" Ошибка Некорректный ввод")
	}
	fmt.Println(test, test1)

	return first, oper, second

}

func main() {
	var (
		first     string
		firstConv int
		firstType bool

		oper string

		preResult int

		second     string
		secondConv int
		secondType bool
	)

	rim := map[string]int{
		"I":        1,
		"II":       2,
		"III":      3,
		"IV":       4,
		"V":        5,
		"VI":       6,
		"VII":      7,
		"VIII":     8,
		"IX":       9,
		"X":        10,
		"XI":       11,
		"XII":      12,
		"XIII":     13,
		"XIV":      14,
		"XV":       15,
		"XVI":      16,
		"XVII":     17,
		"XVIII":    18,
		"XIX":      19,
		"XX":       20,
		"XXI":      21,
		"XXII":     22,
		"XXIII":    23,
		"XXIV":     24,
		"XXV":      25,
		"XXVI":     26,
		"XXVII":    27,
		"XXVIII":   28,
		"XXIX":     29,
		"XXX":      30,
		"XXXI":     31,
		"XXXII":    32,
		"XXXIII":   33,
		"XXXIV":    34,
		"XXXV":     35,
		"XXXVI":    36,
		"XXXVII":   37,
		"XXXVIII":  38,
		"XXXIX":    39,
		"XL":       40,
		"XLI":      41,
		"XLII":     42,
		"XLIII":    43,
		"XLIV":     44,
		"XLV":      45,
		"XLVI":     46,
		"XLVII":    47,
		"XLVIII":   48,
		"XLIX":     49,
		"L":        50,
		"LI":       51,
		"LII":      52,
		"LIII":     53,
		"LIV":      54,
		"LV":       55,
		"LVI":      56,
		"LVII":     57,
		"LVIII":    58,
		"LIX":      59,
		"LX":       60,
		"LXI":      61,
		"LXII":     62,
		"LXIII":    63,
		"LXIV":     64,
		"LXV":      65,
		"LXVI":     66,
		"LXVII":    67,
		"LXVIII":   68,
		"LXIX":     69,
		"LXX":      70,
		"LXXI":     71,
		"LXXII":    72,
		"LXXIII":   73,
		"LXXIV":    74,
		"LXXV":     75,
		"LXXVI":    76,
		"LXXVII":   77,
		"LXXVIII":  78,
		"LXXIX":    79,
		"LXXX":     80,
		"LXXXI":    81,
		"LXXXII":   82,
		"LXXXIII":  83,
		"LXXXIV":   84,
		"LXXXV":    85,
		"LXXXVI":   86,
		"LXXXVII":  87,
		"LXXXVIII": 88,
		"LXXXIX":   89,
		"XC":       90,
		"XCI":      91,
		"XCII":     92,
		"XCIII":    93,
		"XCIV":     94,
		"XCV":      95,
		"XCVI":     96,
		"XCVII":    97,
		"XCVIII":   98,
		"XCIX":     99,
		"C":        100,
	}

	first, oper, second = poluchenie()

	firstConv, firstType, secondConv, secondType = proverka(rim, first, second)

	switch oper {
	case "+":
		preResult = firstConv + secondConv
	case "-":
		preResult = firstConv - secondConv
	case "*":
		preResult = firstConv * secondConv
	case "/":
		preResult = firstConv / secondConv
	default:
		panic("Неизвестная операция. Доступные операции +,-,*,/")
	}

	if firstType == true && secondType == true {
		fmt.Println(preResult)
	} else {
		if preResult >= 0 {
			fmt.Println(conv(rim, preResult))
		} else {
			fmt.Println("ошибка, так как в римской системе нет отрицательных чисел")
		}

	}

}
func conv(Rim map[string]int, preResult int) string {
	var result string
	for i, g := range Rim {
		if g == preResult {
			result = i
		}
	}
	return result
}
func proverka(Rim map[string]int, first string, second string) (int, bool, int, bool) {
	var (
		//first      string
		firstError error
		firstType  bool // if true  Арабские if false  Римские
		firstConv  int
		firstCount bool

		//second      string
		secondError error
		secondType  bool // if true  Арабские if false  Римские
		secondConv  int
		secondCount bool
	)
	_, firstError = strconv.Atoi(first)
	if firstError != nil {
		for i, g := range Rim {
			if i == first {
				firstConv = g
				firstCount = true
				firstType = false
			}

		}
		if firstCount == false {
			panic("Первый введенный элемент не соответствует используемым системам счисления")

		}
	} else {
		firstConv, _ = strconv.Atoi(first)
		firstType = true
	}
	if firstConv > 10 {
		fmt.Println("Ошибка Первое число больше 10")
	}

	_, secondError = strconv.Atoi(second)
	if secondError != nil {
		for i, g := range Rim {
			if i == second {
				secondConv = g
				secondCount = true
				secondType = false
			}

		}
		if secondCount == false {
			panic("Второй введенный элемент не соответствует используемым системам счисления")

		}
	} else {
		secondConv, _ = strconv.Atoi(second)
		secondType = true
	}
	if secondConv > 10 {
		fmt.Println("Ошибка Второе число больше 10")
	}
	if secondType != firstType {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления")
	}
	return firstConv, firstType, secondConv, secondType
}

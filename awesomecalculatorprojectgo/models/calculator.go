package main

import (
	"strconv"
	"strings"
	"unicode"
)

func main() {

}
func calculator(expression string) string {
	if expression == "" {
		return "0"
	}
	for _, char := range expression {
		if unicode.IsLetter(char) {
			return "invalid syntax"
		}
	}
	noSpaceExpression := strings.ReplaceAll(expression, " ", "")
	if strings.Contains(noSpaceExpression, "++") {
		return "invalid syntax"
	}
	if strings.Contains(noSpaceExpression, "--") {
		return "invalid syntax"
	}
	if strings.Contains(noSpaceExpression, "//") {
		return "invalid syntax"
	}
	if strings.Contains(noSpaceExpression, "**") {
		return "invalid syntax"
	}
	if strings.Contains(noSpaceExpression, "%%") {
		return "invalid syntax"
	}
	if len(noSpaceExpression) <= 2 {
		return noSpaceExpression
	}
	if !strings.Contains(noSpaceExpression, "%") || !strings.Contains(noSpaceExpression, "*") || !strings.Contains(noSpaceExpression, "/") || !strings.Contains(noSpaceExpression, "-") || !strings.Contains(noSpaceExpression, "+") {
		return noSpaceExpression
	}

	noSpaceExpression = multiply(noSpaceExpression)
	noSpaceExpression = divide(noSpaceExpression)
	noSpaceExpression = modulus(noSpaceExpression)
	noSpaceExpression = add(noSpaceExpression)
	noSpaceExpression = subtract(noSpaceExpression)

	return noSpaceExpression
}

func add(expression string) string {
	var addedNumberString string
	var initialNumberString string

	for index := 0; index < len(expression); index++ {
		if expression[index] == '+' {
			firstDigit := ""
			secondDigit := ""
			for count := index - 1; count >= 0; count-- {
				if unicode.IsDigit(rune(expression[count])) {
					firstDigit = string(expression[count]) + firstDigit
				} else {
					break
				}
			}
			for count := index + 1; count < len(expression); count++ {
				if unicode.IsDigit(rune(expression[count])) {
					secondDigit += string(expression[count])
				} else {
					break
				}
			}
			firstDigitInt, err1 := strconv.Atoi(firstDigit)
			secondDigitInt, err2 := strconv.Atoi(secondDigit)
			if err1 != nil || err2 != nil {
				return "invalid syntax"
			}
			addedNumber := firstDigitInt + secondDigitInt
			addedNumberString = strconv.Itoa(addedNumber)
			initialNumberString = firstDigit + "+" + secondDigit
			expression = strings.Replace(expression, initialNumberString, addedNumberString, 1)
			index = 0
		}

	}
	return expression
}

func subtract(expression string) string {
	var subtractedNumberString string
	var initialNumberString string

	for index := 0; index < len(expression); index++ {
		if expression[index] == '-' {
			firstDigit := ""
			secondDigit := ""
			for count := index - 1; count >= 0; count-- {
				if unicode.IsDigit(rune(expression[count])) {
					firstDigit = string(expression[count]) + firstDigit
				} else {
					break
				}
			}
			for count := index + 1; count < len(expression); count++ {
				if unicode.IsDigit(rune(expression[count])) {
					secondDigit += string(expression[count])
				} else {
					break
				}
			}
			firstDigitInt, err1 := strconv.Atoi(firstDigit)
			secondDigitInt, err2 := strconv.Atoi(secondDigit)

			if err1 != nil || err2 != nil {
				return "invalid syntax"
			}
			subtractedNumber := firstDigitInt - secondDigitInt
			subtractedNumberString = strconv.Itoa(subtractedNumber)
			initialNumberString = firstDigit + "-" + secondDigit
			expression = strings.Replace(expression, initialNumberString, subtractedNumberString, 1)
			index = 0
		}
	}
	return expression
}

func divide(expression string) string {
	var dividedNumberString string
	var initialNumberString string

	for index := 0; index < len(expression); index++ {
		if expression[index] == '/' {
			firstDigit := ""
			secondDigit := ""

			for count := index - 1; count >= 0; count-- {
				if unicode.IsDigit(rune(expression[count])) {
					firstDigit = string(expression[count]) + firstDigit
				} else {
					break
				}
			}

			for count := index + 1; count < len(expression); count++ {
				if unicode.IsDigit(rune(expression[count])) {
					secondDigit += string(expression[count])
				} else {
					break
				}
			}

			firstDigitInt, err1 := strconv.Atoi(firstDigit)
			secondDigitInt, err2 := strconv.Atoi(secondDigit)

			if err1 != nil || err2 != nil {
				return "invalid syntax"
			}

			if secondDigitInt == 0 {
				return "division by zero"
			}

			dividedNumber := firstDigitInt / secondDigitInt
			dividedNumberString = strconv.Itoa(dividedNumber)
			initialNumberString = firstDigit + "/" + secondDigit

			expression = strings.Replace(expression, initialNumberString, dividedNumberString, 1)
			index = 0
		}
	}
	return expression
}

func multiply(expression string) string {
	var multipliedNumberString string
	var initialNumberString string

	for index := 0; index < len(expression); index++ {
		if expression[index] == '*' {
			firstDigit := ""
			secondDigit := ""

			for count := index - 1; count >= 0; count-- {
				if unicode.IsDigit(rune(expression[count])) {
					firstDigit = string(expression[count]) + firstDigit
				} else {
					break
				}
			}

			for count := index + 1; count < len(expression); count++ {
				if unicode.IsDigit(rune(expression[count])) {
					secondDigit += string(expression[count])
				} else {
					break
				}
			}

			firstDigitInt, err1 := strconv.Atoi(firstDigit)
			secondDigitInt, err2 := strconv.Atoi(secondDigit)

			if err1 != nil || err2 != nil {
				return "invalid syntax"
			}

			multipliedNumber := firstDigitInt * secondDigitInt
			multipliedNumberString = strconv.Itoa(multipliedNumber)
			initialNumberString = firstDigit + "*" + secondDigit

			expression = strings.Replace(expression, initialNumberString, multipliedNumberString, 1)
			index = 0
		}
	}
	return expression
}
func modulus(expression string) string {
	var resultNumberString string
	var initialNumberString string

	for index := 0; index < len(expression); index++ {
		if expression[index] == '%' {
			firstDigit := ""
			secondDigit := ""

			for count := index - 1; count >= 0; count-- {
				if unicode.IsDigit(rune(expression[count])) {
					firstDigit = string(expression[count]) + firstDigit
				} else {
					break
				}
			}

			for count := index + 1; count < len(expression); count++ {
				if unicode.IsDigit(rune(expression[count])) {
					secondDigit += string(expression[count])
				} else {
					break
				}
			}

			firstDigitInt, err1 := strconv.Atoi(firstDigit)
			secondDigitInt, err2 := strconv.Atoi(secondDigit)

			if err1 != nil || err2 != nil {
				return "invalid syntax"
			}

			if secondDigitInt == 0 {
				return "modulus by zero"
			}

			resultNumber := firstDigitInt % secondDigitInt
			resultNumberString = strconv.Itoa(resultNumber)
			initialNumberString = firstDigit + "%" + secondDigit

			expression = strings.Replace(expression, initialNumberString, resultNumberString, 1)
			index = 0
		}
	}
	return expression
}

func bracket(expression string) string {
	//expression = strings.ReplaceAll(expression, "\n", " ")
	return ""
}

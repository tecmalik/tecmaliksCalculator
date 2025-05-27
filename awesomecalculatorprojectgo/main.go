package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

// type calculator struct {
// }
func calculate(expression string) string {
	//value string json `json:"value"`

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
		return "invalid syntax, wait for upgrade"
	}
	if strings.Contains(noSpaceExpression, "--") {
		return "invalid syntax"
	}
	if strings.Contains(noSpaceExpression, "//") {
		return "invalid syntax, wait for upgrade"
	}
	if strings.Contains(noSpaceExpression, "**") {
		return "invalid syntax, wait for upgrade"
	}
	if strings.Contains(noSpaceExpression, "%%") {
		return "invalid syntax, wait for upgrade"
	}

	if strings.Contains(noSpaceExpression, "+-") {
		return "invalid syntax, wait for upgrade"
	}
	if strings.Contains(noSpaceExpression, "-+") {
		return "invalid syntax"
	}
	if strings.Contains(noSpaceExpression, "/-") {
		return "invalid syntax, wait for upgrade"
	}
	if strings.Contains(noSpaceExpression, "*-") {
		return "invalid syntax, wait for upgrade"
	}
	if strings.Contains(noSpaceExpression, "%+") {
		return "invalid syntax, wait for upgrade"
	}

	if len(noSpaceExpression) <= 2 {
		return noSpaceExpression
	}
	if !strings.Contains(noSpaceExpression, "%") || !strings.Contains(noSpaceExpression, "*") || !strings.Contains(noSpaceExpression, "/") || !strings.Contains(noSpaceExpression, "-") || !strings.Contains(noSpaceExpression, "+") {
		return noSpaceExpression
	}

	//noSpaceExpression = evaluateBrackets(noSpaceExpression)

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

func evaluateBrackets(expression string) string {
	for strings.Contains(expression, "(") {
		start := strings.LastIndex(expression, "(")
		end := strings.Index(expression[start:], ")") + start
		if end < start {
			return "invalid syntax: missing closing bracket"
		}
		subExpr := expression[start+1 : end]
		result := calculate(subExpr)
		expression = expression[:start] + result + expression[end+1:]
	}
	return expression
}

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.POST("/calculate", handleCalculate)
	router.Run(":9090")
}

func handleCalculate(c *gin.Context) {
	var input struct {
		Expression string `json:"expression"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request format",
			"details": err.Error(),
		})
		return
	}
	if strings.TrimSpace(input.Expression) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "expression cannot be empty",
		})
		return
	}
	result := calculate(input.Expression)

	c.JSON(http.StatusOK, gin.H{
		"message": "Calculated successfully!",
		"result":  result,
	})
	if result == "division by zero" ||
		result == "modulus by zero" ||
		strings.HasPrefix(result, "invalid syntax") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result,
		})
		return
	}
}

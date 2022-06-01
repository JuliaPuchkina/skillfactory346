package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readSortWrite("input.txt", "output.txt")
}

func readSortWrite(inputName string, outputName string) {
	f, err := os.OpenFile(inputName, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileReader := bufio.NewReader(f)

	_ = os.Remove(outputName)
	ff, err := os.OpenFile(outputName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer ff.Close()

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		mathRegex := regexp.MustCompile("[0-9]+[-+*/]{1}[0-9]+")
		isMath := mathRegex.MatchString(string(line))
		if isMath {
			finRes := calcMath(line)

			writer := bufio.NewWriter(ff)
			writer.WriteString(string(finRes))
			writer.WriteString(string("\n"))
			writer.Flush()
		}
	}

}

func calcMath(expr []byte) (result []byte) {
	numRegex := regexp.MustCompile("[0-9]+")
	nums := numRegex.FindAllString(string(expr), -1)

	signRegex := regexp.MustCompile("[-+*/]")
	sign := signRegex.FindAllString(string(expr), -1)

	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	var math int

	if sign[0] == "+" {
		math = num1 + num2
	} else if sign[0] == "-" {
		math = num1 - num2
	} else if sign[0] == "*" {
		math = num1 * num2
	} else if sign[0] == "/" {
		math = num1 / num2
	}

	finNum := strconv.Itoa(math)

	var resString []string
	resString = append(resString, string(nums[0]), string(sign[0]), string(nums[1]), "=", finNum)

	stringByte := strings.Join(resString, "\x20")

	result = []byte(stringByte)

	return result
}

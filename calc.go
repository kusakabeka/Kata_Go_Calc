package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	operators = "+-*/"
)

// func convert roman number to arabic numbers
//func roman_to_arabic(str string) int {
//	n := 1
//	var c int = n * 2
//	return c
//}

func calculate(x, y int, operator string) {
	switch operator {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		fmt.Println(x / y)
	}
}

func animation(str string) {
	for _, char := range str {
		fmt.Print(string(char))
		time.Sleep(1 * time.Millisecond)
	}
}

func whichOS() string {
	return runtime.GOOS
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

//func isOperator(str string) bool {
//
//}

func clear() {
	if whichOS() == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	text := "\033[1;36m\n░█████╗░░█████╗░██╗░░░░░░█████╗░██╗░░░██╗██╗░░░░░░█████╗░████████╗░█████╗░██████╗░\n██╔══██╗██╔══██╗██║░░░░░██╔══██╗██║░░░██║██║░░░░░██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗\n██║░░╚═╝███████║██║░░░░░██║░░╚═╝██║░░░██║██║░░░░░███████║░░░██║░░░██║░░██║██████╔╝\n██║░░██╗██╔══██║██║░░░░░██║░░██╗██║░░░██║██║░░░░░██╔══██║░░░██║░░░██║░░██║██╔══██╗\n╚█████╔╝██║░░██║███████╗╚█████╔╝╚██████╔╝███████╗██║░░██║░░░██║░░░╚█████╔╝██║░░██║\n░╚════╝░╚═╝░░╚═╝╚══════╝░╚════╝░░╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝\033[0m\n"
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(1 * time.Millisecond)
	}

	reader := bufio.NewReader(os.Stdin)

	welcom := "                            Добро Пожаловать!\n"
	animation(welcom)
	list_oper := "                    Список операций, которые доступны:\n"
	animation(list_oper)
	oper_1 := "\033[1;34m[+]\033[0m - \033[4;32mОперация сложения.\033[0m\n"
	animation(oper_1)
	oper_2 := "\033[1;34m[-]\033[0m - \033[4;32mОперация вычитания.\033[0m\n"
	animation(oper_2)
	oper_3 := "\033[1;34m[*]\033[0m - \033[4;32mОперация умножения.\033[0m\n"
	animation(oper_3)
	oper_4 := "\033[1;34m[/]\033[0m - \033[4;32mОперация деления.\033[0m\n"
	animation(oper_4)
	info_1 := "\033[1;33m[?]\033[0m Вы можете использовать \033[4;33mАрабские\033[0m числа(1,2,3 ... 10) либо же \033[4;33mРимские\033[0m (I, II, IV ... X)!\n"
	animation(info_1)
	info_2 := "Введите \033[1;33m--help\033[0m, чтобы ознакомиться с командами.\n"
	animation(info_2)
	console := "\033[5;35m~> \033[0m"

	stringForParsing := ""
	for stringForParsing != ":q" {

		// if enter '--help'
		if stringForParsing == "--help" {
			fmt.Print("\033[5;36m\033[0m   \034 \033[1;32m:q\033[0m - Выход из программы.\n   \034 \033[1;32m:o\033[0m - Напомнить доступные операции.\n\033[0m   \034 \033[1;32mclear\u001B[0m - Очистка консоли программы.\n")
		} else if stringForParsing == ":q" {
			break
		} else if stringForParsing == ":o" {
			fmt.Print("   \034 \033[1;32m[+]\033[0m - Операция сложения.\n")
			fmt.Print("   \034 \033[1;32m[-]\033[0m - Операция вычитания.\n")
			fmt.Print("   \034 \033[1;32m[*]\033[0m - Операция умножения.\n")
			fmt.Print("   \034 \033[1;32m[/]\033[0m - Операция деления.\n")
		} else if stringForParsing == "clear" {
			clear()
			fmt.Println("\033[1;36m\n░█████╗░░█████╗░██╗░░░░░░█████╗░██╗░░░██╗██╗░░░░░░█████╗░████████╗░█████╗░██████╗░\n██╔══██╗██╔══██╗██║░░░░░██╔══██╗██║░░░██║██║░░░░░██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗\n██║░░╚═╝███████║██║░░░░░██║░░╚═╝██║░░░██║██║░░░░░███████║░░░██║░░░██║░░██║██████╔╝\n██║░░██╗██╔══██║██║░░░░░██║░░██╗██║░░░██║██║░░░░░██╔══██║░░░██║░░░██║░░██║██╔══██╗\n╚█████╔╝██║░░██║███████╗╚█████╔╝╚██████╔╝███████╗██║░░██║░░░██║░░░╚█████╔╝██║░░██║\n░╚════╝░╚═╝░░╚═╝╚══════╝░╚════╝░░╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝\033[0m")
			fmt.Println("                            Добро Пожаловать!")
			fmt.Println("                    Список операций, которые доступны:")
			fmt.Println("\033[1;34m[+]\033[0m - \033[4;32mОперация сложения.\033[0m")
			fmt.Println("\033[1;34m[-]\033[0m - \033[4;32mОперация вычитания.\033[0m")
			fmt.Println("\033[1;34m[*]\033[0m - \033[4;32mОперация умножения.\033[0m")
			fmt.Println("\033[1;34m[/]\033[0m - \033[4;32mОперация деления.\033[0m")
			fmt.Println("\033[1;33m[?]\033[0m Вы можете использовать \033[4;33mАрабские\033[0m числа(1,2,3 ... 10) либо же \033[4;33mРимские\033[0m (I, II, IV ... X)!")
			fmt.Println("Введите \033[1;33m--help\033[0m, чтобы ознакомиться с командами.")
		}
		fmt.Print(console)
		var err error
		stringForParsing, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		stringForParsing = strings.TrimSpace(stringForParsing)

		s := strings.Split(stringForParsing, " ")

		number1, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println("Error")
		}

		number2, err := strconv.Atoi(s[2])
		if err != nil {
			fmt.Println("Error")
		}

		if (len(s) == 3) && isNumber(s[0]) && isNumber(s[2]) && strings.Contains(operators, s[1]) {
			calculate(number1, number2, s[1])
		} else {
			fmt.Println("Error")
		}
	}

}

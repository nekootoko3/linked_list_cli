package main

import (
	"bufio"
	"fmt"
	"linked_list/linked_list"
	"os"
	"strconv"
)

func main() {
	fmt.Println("You can use choose order below")
	fmt.Println("1 prepend | 2 removeHead | 3 isExist | 4 remove | 5 getLength | 6 printAll")
	fmt.Println("Which order?")
	l := linked_list.LinkedList{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		order, err := strconv.Atoi(input)
		if err != nil || (order < 1 && 6 < order) {
			fmt.Println("Wrong order")
			fmt.Println("next")
			fmt.Println("1 prepend | 2 removeHead | 3 isExist | 4 remove | 5 getLength | 6 printAll")
		}

		switch order {
		case 1:
			num, err := getNumber(scanner)
			if err != nil {
				fmt.Println("Your input is not number")
			}
			l.Prepend(num)
		case 2:
			l.RemoveHead()
		case 3:
			num, err := getNumber(scanner)
			if err != nil {
				fmt.Println("Your input is not number")
			}
			if l.IsExist(num) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case 4:
			num, err := getNumber(scanner)
			if err != nil {
				fmt.Println("Your input is not number")
			}
			l.Remove(num)
		case 5:
			l.GetLength()
		case 6:
			l.PrintAll()
		}

		fmt.Println("")
		fmt.Println("---- next")
		fmt.Println("1 prepend | 2 removeHead | 3 isExist | 4 remove | 5 getLength | 6 printAll")
	}
}

func getNumber(scanner *bufio.Scanner) (int, error) {
	fmt.Println("Please input number")

	scanner.Scan()
	input := scanner.Text()
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return num, nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your birth date as (DD/MM/YYYY):")
	birthDate, err := reader.ReadString('\n')
	birthDate = strings.ReplaceAll(birthDate, "\n", "")
	if err != nil {
		fmt.Printf("error while reading the input :: %v", err)
		return
	}

	if len(birthDate) != 10 {
		fmt.Printf("Please enter the birth date in correct format\nE.g : 03/02/2000")
	}

	birthDateSlice := strings.Split(birthDate, "/")
	year, _ := strconv.Atoi(birthDateSlice[2])
	month, _ := strconv.Atoi(birthDateSlice[1])
	day, _ := strconv.Atoi(birthDateSlice[0])

	birthDateInTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	age := now.Year() - birthDateInTime.Year()

	fmt.Printf("You are %d years old.\n", age)
}

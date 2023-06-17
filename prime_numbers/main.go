package main

import(
	stuff "primeNumbers/prime"
	"bufio"
	"log"
	"os"
	"fmt"
	"strconv"
)

func input(text string) string{
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s", text)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return line[:len(line)-1]
}

func main() {
	a := input("Start: ")
	b := input("Stop: ")

	aInt, err := strconv.Atoi(a)
	bInt, err := strconv.Atoi(b)

	if err != nil {
		log.Fatal(err)
	}

	stuff.Sequence(aInt,bInt)
}
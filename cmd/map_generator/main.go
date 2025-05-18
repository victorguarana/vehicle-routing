package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const filename = "map_105"
const warehouseQnt = 3
const customerQnt = 100

func main() {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writeRandomWarehouses(writer)
	writeNewLine(writer)
	writeRandomCustomers(writer)
}

func writeRandomWarehouses(writer *bufio.Writer) {
	for i := 1; i <= warehouseQnt; i++ {
		line := fmt.Sprintf("Warehouse%d;%d;%d;%d", i, randomPosition(), randomPosition(), 0)
		if i < warehouseQnt {
			line += "\n"
		}
		if _, err := writer.WriteString(line); err != nil {
			panic(err)
		}
	}
}

func writeRandomCustomers(writer *bufio.Writer) {
	for i := 1; i <= customerQnt; i++ {
		line := fmt.Sprintf("Customer%d;%d;%d;%d", i, randomPosition(), randomPosition(), randomPackage())
		if i < customerQnt {
			line += "\n"
		}
		if _, err := writer.WriteString(line); err != nil {
			panic(err)
		}
	}
}

func writeNewLine(writer *bufio.Writer) {
	if _, err := writer.WriteString("\n"); err != nil {
		panic(err)
	}
}

func randomPackage() int {
	return (rand.Int() % 13) + 1
}

func randomPosition() int {
	return (rand.Int() % 100) - 50
}

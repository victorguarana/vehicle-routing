package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const filename = "map_105"
const depositQnt = 3
const clientQnt = 100

func main() {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writeRandomWarehouses(writer)
	writeRandomClients(writer)
}

func writeRandomWarehouses(writer *bufio.Writer) {
	for i := 1; i <= depositQnt; i++ {
		line := fmt.Sprintf("Deposito%d;%d;%d;%d\n", i, randomPosition(), randomPosition(), 0)
		if _, err := writer.WriteString(line); err != nil {
			panic(err)
		}
	}
}

func writeRandomClients(writer *bufio.Writer) {
	for i := 1; i <= clientQnt; i++ {
		line := fmt.Sprintf("Cliente%d;%d;%d;%d\n", i, randomPosition(), randomPosition(), randomPackage())
		if _, err := writer.WriteString(line); err != nil {
			panic(err)
		}
	}
}

func randomPackage() int {
	return (rand.Int() % 13) + 1
}

func randomPosition() int {
	return (rand.Int() % 100) - 50
}

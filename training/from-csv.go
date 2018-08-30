package training

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// FromCSV obtem os dados de treino
func FromCSV(filename string, input, output int) []Data {
	file, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(file))

	var data []Data
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		count := len(line)
		d := Data{
			Input:  make([]int, input),
			Output: make([]int, output),
		}
		for i := 0; i < count; i++ {
			n, err := strconv.Atoi(line[i])
			if err != nil {
				log.Fatal(err)
			}
			if i < input {
				d.Input[i] = n
			} else {
				d.Output[i-input] = n
			}
		}
		data = append(data, d)
	}
	return data
}

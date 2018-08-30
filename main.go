package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/andersonmarin/neural/neuron"
	"github.com/andersonmarin/neural/training"
)

func main() {
	rand.Seed(time.Now().Unix())

	data := training.FromCSV("training.csv", 6, 1)

	trainedNetwork := training.Run(func() neuron.Network {
		return neuron.NewNetwork(6, 12, 8, 1)
	}, data)

	// bico, olhos, penas, bipede, asas, voa, result
	result := trainedNetwork.Run([]int{0, 1, 0, 0, 0, 0})
	fmt.Printf("test tatu Result = %d\n", result)

	save(fmt.Sprintf("trained-network-%d.xml", time.Now().Unix()), trainedNetwork)
}

func save(filename string, data interface{}) {
	str, _ := xml.MarshalIndent(data, "", "\t")
	ioutil.WriteFile(filename, str, 0644)
}

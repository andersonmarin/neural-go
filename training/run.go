package training

import (
	"fmt"
	"reflect"

	"github.com/andersonmarin/neural/neuron"
)

// Run executa o treinamento
func Run(networkFactory func() neuron.Network, data []Data) neuron.Network {
	var (
		bestNetwork neuron.Network
		bestPoints  int
	)

	count := len(data)
	for d := 0; d < 512; d++ {
		currentNetwork := networkFactory()
		currentPoints := 0

		for i := 0; i < count; i++ {
			output := currentNetwork.Run(data[i].Input)
			if reflect.DeepEqual(data[i].Output, output) {
				currentPoints++
			}
		}

		if currentPoints > bestPoints {
			bestNetwork = currentNetwork
			bestPoints = currentPoints
			fmt.Printf("best network %.2f%%\n", float64(currentPoints)/float64(count)*100)

			if currentPoints == count {
				fmt.Printf("stop training on interaction %d\n", d+1)
				break
			}
		}
	}
	return bestNetwork
}

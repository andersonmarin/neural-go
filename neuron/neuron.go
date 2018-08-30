package neuron

import (
	"math"
	"math/rand"
)

// Neuron representa um neurônio
type Neuron struct {
	Weights []int
}

// NewNeuron cria um novo neurônio
func NewNeuron(count int) Neuron {
	return Neuron{
		Weights: randomWeights(count),
	}
}

// Run executa um neurônio
func (n *Neuron) Run(input []int) int {
	output := 0
	for i := 0; i < len(input); i++ {
		output += input[i] * n.Weights[i]
	}
	// if output > 0 {
	// 	return 1
	// }
	if sigmoid(float64(output)) > 0.5 {
		return 1
	}
	return 0
}

func randomWeights(count int) []int {
	weights := make([]int, count)
	for i := 0; i < count; i++ {
		if rand.Float32() > 0.5 {
			weights[i] = 1
		} else {
			weights[i] = -1
		}
	}
	return weights
}

func sigmoid(x float64) float64 {
	return x / (1.0 + math.Abs(x))
}

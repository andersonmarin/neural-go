package neuron

// Layer representa uma camada de neurônios
type Layer struct {
	Neurons []Neuron
}

// NewLayer cria uma camada de neurônios
func NewLayer(width, weights int) Layer {
	layer := Layer{
		Neurons: make([]Neuron, width),
	}
	for i := 0; i < width; i++ {
		layer.Neurons[i] = NewNeuron(weights)
	}
	return layer
}

// Run executa uma camada de neurônios
func (l *Layer) Run(input []int) []int {
	count := len(l.Neurons)
	output := make([]int, count)
	for i := 0; i < count; i++ {
		output[i] = l.Neurons[i].Run(input)
	}
	return output
}

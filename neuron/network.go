package neuron

// Network representa uma rede de neurônios
type Network struct {
	Input  int
	Width  int
	Depth  int
	Output int
	Layers []Layer
}

// NewNetwork cria uma rede de neurônios
func NewNetwork(input, width, depth, output int) Network {
	network := Network{
		Input:  input,
		Width:  width,
		Depth:  depth,
		Output: output,
		Layers: make([]Layer, depth+1),
	}

	network.Layers[0] = NewLayer(width, input)

	for d := 1; d < depth; d++ {
		layer := NewLayer(width, width)
		network.Layers[d] = layer
	}

	network.Layers[depth] = NewLayer(output, width)

	return network
}

// Run executa uma rede de neurônios
func (n Network) Run(input []int) []int {
	count := len(n.Layers)
	var output []int
	output = input
	for i := 0; i < count; i++ {
		output = n.Layers[i].Run(input)
	}
	return output
}

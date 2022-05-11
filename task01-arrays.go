package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	for el := 0; el < len(input); el++ {
		sum += input[el]
	}

	return sum / float32(len(input))
}

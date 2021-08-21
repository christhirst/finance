package indicator

func avarage(j []float64) float64 {
	var sum float64
	var timesum int
	for t, c := range j {
		sum += c
		timesum += t
	}

	return (sum/float64(timesum))

}


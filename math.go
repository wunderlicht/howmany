package main

// could have been a generic but clear is better than clever.
// will return +Inf when total is 0
func percent(value, total int) float64 {
	return float64(value) / float64(total) * 100
}

// could have been a generic but clear is better than clever.
func average(d []int) float64 {
	sum := 0.0
	if len(d) == 0 {
		return 0.0
	}
	for i := range d {
		sum += float64(d[i])
	}
	return sum / float64(len(d))
}

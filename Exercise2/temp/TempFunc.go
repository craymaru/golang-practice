package temp

func DivideCall(beDivNum int, divNum float64) float64 {
	return dividing(beDivNum, divNum)
}

func dividing(beDivNum int, divNum float64) float64 {
	var i float64
	i = float64(beDivNum) / divNum
	return i
}

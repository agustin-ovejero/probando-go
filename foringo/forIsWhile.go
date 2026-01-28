package foringo



func ForEsWhile() int{
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

package internal

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c Calculator) Add(num1, num2 int) int {
	result := num1 + num2
	panic("implement me")
	return result
}

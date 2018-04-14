package _type

type AlgebraOfSets interface {
	Subtract(minuend []Timerange, subtrahend []Timerange) []Timerange
}

package model

type Ordered interface {
	~string | ~int | ~float64
}

type OrderedStringer interface {
	Ordered
	String() string
}

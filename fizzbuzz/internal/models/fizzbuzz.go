package models

type FizzBuzz struct {
	Int1    int    `json:"int1"`
	Int2    int    `json:"int2"`
	Int3    int    `json:"int3"`
	String1 string `json:"string1"`
	String2 string `json:"string2"`
}

var DefaultFizzBuzz = FizzBuzz{
	Int1:    3,
	Int2:    5,
	Int3:    100,
	String1: "fizz",
	String2: "buzz",
}

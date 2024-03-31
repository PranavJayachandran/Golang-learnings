package main

type Person struct {
	NAME string `json:"name"`
	AGE  string `json:"age"`
	ID   string `json:"id"`
}

var data = []Person{
	{NAME: "One", AGE: "10", ID: "12"},
	{NAME: "Two", AGE: "20", ID: "2"},
}

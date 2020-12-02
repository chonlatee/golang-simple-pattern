package main

import "fmt"

func main() {
	d := developer{
		FirstName: "John",
		LastName:  "Doe",
		Rank: 2,
	}

	b := &bonus{}
	d.Accept(b)
	d.FullName()
	fmt.Println("Developer bonus", b.GetBonus())

	dt := director{
		FirstName: "Jane",
		LastName:  "Doe",
		Rank: 4,
	}

	dtBonus := &bonus{}
	dt.Accept(dtBonus)
	dt.FullName()
	fmt.Println("Director bonus", dtBonus.GetBonus())

}

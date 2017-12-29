package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string
}

func main() {
	p1 := person{
		First: "a",
		Last:  "z",
	}
	p2 := person{
		First: "b",
		Last:  "y",
	}
	xp := []person{p1, p2}

	fmt.Println("go data", xp)

	bs, err := json.Marshal(xp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("json", string(bs))
	}
	fmt.Printf("---------------\n");
	j := `[{"first":"a","Last":"z"},{"First":"b","Last":"y"}]`
	fmt.Println("json", j)
	xp1 := []person{}
	errs := json.Unmarshal([]byte(j), &xp1)
	if errs != nil {
		fmt.Println(errs)
	} else {
    fmt.Printf("%+v\n", xp1) // + also print the field names of the struct
  }
}

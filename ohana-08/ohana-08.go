package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string
	Email   string
	Age     int
	Address []string
}

func main() {
	a := Person{
		"Billy Crystal",
		"billy@billy.billy",
		43,
		[]string{
			"Alabama",
			"USA",
		},
	}
	fmt.Println(a)

	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	var js []byte
	js, err = json.Marshal(users)
	fmt.Println("Marshaling JSON:", string(js))

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var ujs []Bond

	err = json.Unmarshal([]byte(s), &ujs)

	if err != nil {
		fmt.Println("Error Unmarshaling JSON:", err)
	}

	fmt.Printf("%v\n\n", ujs)

	err = json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println("Error when encoding:", err)
	}

}

type user struct {
	First string
	Age   int
}

type Bond struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

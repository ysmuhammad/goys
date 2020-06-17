package main

import (
	"fmt"
)

type heaven struct{
	country string
	city string
}

type football struct{
	clubName string
	country string
	region string
	players []string
}

func main() {
	heav1 := heaven{
		country: "Hawaii",
		city: "Oahu",
	}

	heav2 := heaven{
		country: "Indonesia",
		city: "Mentawai",
	}
	
	fmt.Println(heav1, heav2)
	fmt.Printf("Best city to surf in Indonesia? \n- %v\n", heav2.city)
	fmt.Printf("Best city to surf in Hawaii? \n- %v\n", heav1.city)

	nestedStruct()

	structSlice()

	structMap()
}

func nestedStruct() {
	type boss struct{
		firstName string
		lastName string
	}
	type kids struct{
		firstName string
		lastName string
	}
	type fam struct{
		boss
		kids
		status string
	}

	boss1 := boss{
		firstName: "John",
		lastName: "Doe",
	}

	kids1 := kids{
		firstName: "Greenwood",
		lastName: "Mason",
	}

	all := fam{
		boss: boss1,
		kids: kids1,
		status: "Happy Family!",
	}

	fmt.Printf("\nBoss first name: %v\n", all.boss.firstName)
	fmt.Printf("Kids last name: %v\n", all.kids.lastName)
	fmt.Printf("Family status: %v\n", all.status)
}

func structSlice(){
	type jkt struct{
		area string
		station string
		population []int
	}

	v1 := jkt{
		area: "Lebak Bulus",
		station: "Fatmawati MRT Station",
		population: []int{
			15094,
			99999,
		},
	}
	fmt.Printf("=========\nArea: %v\nStation Name: %v\nPopulation: %d\n", v1.area, v1.station,v1.population[1])
	fmt.Println("========")
}

func structMap() {
	f1 := football{
		clubName: "Manchester United",
		country: "England",
		region: "Europe",
		players: []string{"Rooney","Cristiano Ronaldo","Beckham","Cantona","Best"},
	}
	fmt.Println(f1)

	mu := map[string]football{
		f1.clubName: f1,
	}

	fmt.Println(mu)

	for k, v := range mu{
		fmt.Println(k, v)
		for _, vl := range v.players {
		fmt.Println("Top United Players: ", vl)
		}
	}
}
package arraysmapsslice

import "fmt"

type floatMap map[string]float64

func Maps() {
	fmt.Println("Maps")

	//websites := []string{"https://google.com", "https://aws.com"}
	websites := map[string]string{
		"google":             "https://google.com",
		"amzon web services": "https://aws.com",
	} /// [string] is the value type

	fmt.Println(websites["google"])

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 5.0

	for index, value := range courseRatings {
		fmt.Println(index)
		fmt.Println(value)
	}

}

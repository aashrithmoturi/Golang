package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["Py"] = "python"
	languages["Rb"] = "Ruby"
	fmt.Println(languages)

	delete(languages, "Py")
	fmt.Println(languages)

	fmt.Println(len(languages))

	//loops are interesting in goloang
	for key, value := range languages {
		fmt.Println(key, value)
	}

}

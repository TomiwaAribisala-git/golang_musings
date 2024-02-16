package main

//"io"
//"net/http"
//"net/url"
//"encoding/json"
//"strings"
//"time"

func main() {
	/*
			var smallVal bool = false
			fmt.Println(smallVal)
			fmt.Printf("Variable is of type: %T \n", smallVal)

			presentTime := time.Now()
			fmt.Println(presentTime)

			fruitlist := []string{}
			fmt.Printf("Fruitlist is of type: %T\n", fruitlist)

			languages := make(map[string]int)
			languages["js"] = 1
			fmt.Println("List of all languages: ", languages)

			for _, value := range languages {
				fmt.Printf("The value of languages is %v\n", value)
			}


			content := "I am writing this text into a file"
			file, err := os.Create("./log.txt")
			if err != nil {
				fmt.Printf("The file opening is giving %v\n:", err)
			}
			length, err := io.WriteString(file, content)
			if err != nil {
				panic(err)
			}
			fmt.Println("length is:", length)
			readfile("./log.txt")
			defer file.Close()


			const url = "https://lco.dev"
			response, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			databytes, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println("The body of the url: ", string(databytes))

			const myurl = "http://localhost:3000/get"
			response, err := http.Get(myurl)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()

			fmt.Println("Status code is: ", response.StatusCode)
			fmt.Println("Content length: ", response.ContentLength)
			content, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))

			const myurl = "http://localhost:3000/post"
			// fake json payload
			requestBody := strings.NewReader(`
				{
					"coursename": "Let's go with golang",
					"price": 0,
					"platform": "learncodeonline"
				}
			`)
			response, err := http.Post(myurl, "application/json", requestBody)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			content, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))


			const myurl = "http://localhost:3000/postform"
			//how to send form data in golang
			data := url.Values{}
			data.Add("firstname", "tomiwa")
			data.Add("lastname", "aribisala")
			data.Add("age", "20")

			response, err := http.PostForm(myurl, data)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			content, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))


		// Creating(marshalling) JSON data in golang
		type course struct {
			Name  string   `json:"Coursename"`
			Price int      `json:"Price"`
			Tags  []string `json:"Tags"`
		}

		lcoCourses := []course{
			{"reactjs bootcamp", 530, []string{"react", "native"}},
			{"nodejs bootcamp", 480, []string{"javascript", "nodejs"}},
			{"devops bootcamp", 390, []string{"docker", "golang"}},
		}

		finaljson, err := json.MarshalIndent(lcoCourses, "lco", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", finaljson)
		// Consuming(unmarshalling) JSON data in golang
		jsonWebData := []byte(`
		{
			"Coursename": "devops bootcamp",
			"Price": 390,
			"Tags": ["docker","golang"]
			}
		`)

		var lcoCourse course

		checkValid := json.Valid(jsonWebData)
		if checkValid {
			fmt.Println("Json data is valid")
			json.Unmarshal(jsonWebData, &lcoCourse)
			fmt.Printf("%#v\n", lcoCourse)
		} else {
			fmt.Println("Json data is not valid")
		}
	*/
	// Read gorilla/mux and Go http package integration
	// Building an API in Golang
}

/*
func readfile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is:", string(databyte))
}

func main() {
	working("chase")

	go working("dune")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	messages := make(chan string)
	go func() {
		messages <- "maker"
	}()
	msg := <-messages
	fmt.Println(msg)

	time.Sleep(time.Second)
	fmt.Println("DONE")
}
*/

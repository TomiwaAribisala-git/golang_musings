package main

import (
	//"io"
	//"math/rand"
	"fmt"
	"log"
	"net/http"

	//"strconv"
	//"net/url"
	"encoding/json"
	//"strings"
	//"time"
	"github.com/gorilla/mux"
)

// Building API in Golang

// Commands to run API and testing with Postman
// go run main.go &
// localhost:4000			GET: default route
// localhost:4000/courses	GET: get all courses
// localhost:4000/course/2	GET: get one course
// localhost:4000/course  	POST: create a course
// localhost:4000/course/2  PUT: update a course
// localhost:4000/course/2  DELETE: delete a course

type Course struct {
	CourseId string		`json:"courseid"`
	CourseName string 	`json:"coursename"`
	CoursePrice int 	`json:"price"`
	Author *Author		`json:"author"`
}
type Author struct {
	Fullname string		`json:"fullname"`
	Website string		`json:"website"`
}

// fake DB
var courses []Course

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to writing an API</h1>"))
}

// serve API json response for all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
// serve API json response for one course based on request id
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	
	// grab id from request
	params := mux.Vars(r)

	// loop through courses, finding matching id and return a response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
// Add a course controller
func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		// what if: body is empty
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about {}
	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)
	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	//rand.Seed(time.Now().UnixNano())
	//newCourse.CourseId = strconv.Itoa(rand.Intn(100))

	// append course into courses
	courses = append(courses, newCourse)
	// Respond with the created course
	_ = json.NewEncoder(w).Encode(newCourse)
	return
}
// Update a course controller
func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[:index+1]... )
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			_ = json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("No course is found with the given id")
	return
}
// Delete a course controller
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	//loop, id, remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[:index+1]... )
			break
		}
	}
	json.NewEncoder(w).Encode("No course is found with the given id")
	return
}

func main() {
	fmt.Println("API Demo")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Tomiwa Aribisala", Website: "go.dev"}})
	
	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

/*
// gorilla/mux router and Go http package integration
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the golang series</h1>"))
}
/*

/*
func main() {
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
}

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

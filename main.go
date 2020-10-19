package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Id string `json:"Id"`
}

type ErrorMessage struct {
	Message string `json:"Message"`
}
type Option struct {
	Message string `json:"Message"`
}

//Articles - local DataBase
var Articles []Article

//GET request for /articles
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hint: getAllArticles woked.....")
	text := "встав рядом ссылкой http://localhost:8000/articles/ ='{сыслка нас скачивыемый снимок}'"
	json.NewEncoder(w).Encode(text)
}

const lenPath = len("/download/")

//GET request for article with ID
func GetArticleWithId(w http.ResponseWriter, r *http.Request) {
	name := "<form method='POST' action='/add'>URL: <input type='text' name='url'><input type='submit' value='Add'></form>"
	fmt.Fprintf(w, name)

	//if len(os.Args) != 3 {
	//	fmt.Println(vars)
	//	os.Exit(1)
	//}
	//url := os.Args[1]
	//filename := os.Args[2]
	//
	//err := DownloadFile(url, filename)
	//if err != nil {
	//	panic(err)
	//}
	//if !find {
	//	var erM = ErrorMessage{Message: "Not found article with that ID"}
	//	json.NewEncoder(w).Encode(erM)
	//}
}

//func DownloadFile(url string, filepath string) error {
//	// Create the file
//	out, err := os.Create(filepath)
//	if err != nil {
//		return err
//	}
//	defer out.Close()
//
//	// Get the data
//	resp, err := http.Get(url)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// Write the body to file
//	_, err = io.Copy(out, resp.Body)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func main() {
	fmt.Println("REST API V2.0 worked....")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", GetAllArticles).Methods("GET")
	myRouter.HandleFunc("/download/", GetArticleWithId).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

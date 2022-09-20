package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"example.com/demo/models"
)

func CreateBook(book *models.Book) error {

	return nil
}

func GetAllBooks() ([]models.Book, error) {
	content, err := ioutil.ReadFile("./bookInfo.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	fmt.Println(content)
	var book []models.Book
	err = json.Unmarshal(content, &book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(book)
	return book, nil
	// fmt.Printf("%v\n", softwares)
	// data := models.Book{}
	// // Now let's unmarshall the data into `payload`
	// var payload map[string]interface{}
	// err = json.Unmarshal(content, &data)
	// if err != nil {
	// 	log.Fatal("Error during Unmarshal(): ", err)
	// }

	// _ = json.Unmarshal([]byte(file), &data)
}

func GetBookByName(string) (bookname *models.Book) {
	return nil
}

func DeleteBook(bookname *models.Book) error {
	return nil
}

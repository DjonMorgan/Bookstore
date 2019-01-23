package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../Bookstore/db"
)

func main() {
	db.Init()
	fmt.Println("pitohui")
	http.HandleFunc("/new", Bookstore)
	http.HandleFunc("/up", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/send", IndexHandler)
	http.ListenAndServe("localhost:8181", nil)
}

type Books struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication int64  `json:"publication"`
}

//добавление книги в бд
func Bookstore(w http.ResponseWriter, r *http.Request) {
	new, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var bk Books
	err = json.Unmarshal(new, &bk)
	fmt.Println(bk)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	database := db.Manager()
	fmt.Println(database)

	_, err = database.Exec("insert into book.books (book_name, author, publication_date) values (?, ?, ?)",
		bk.Name, bk.Author, bk.Publication)

	fmt.Println(err)
}

//обновление данных про книги в бд
func Update(w http.ResponseWriter, r *http.Request) {
	up, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var bk Books
	err = json.Unmarshal(up, &bk)
	fmt.Println(bk)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	database := db.Manager()
	_, err = database.Exec("update book.books set book_name=?, author=?, publication_date=? where id_book = ?",
		bk.Name, bk.Author, bk.Publication, bk.Id)

	fmt.Println(err)
}

//удаление книги
func Delete(w http.ResponseWriter, r *http.Request) {
	del, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var bk Books
	err = json.Unmarshal(del, &bk)
	fmt.Println(bk)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	database := db.Manager()
	_, err = database.Exec("delete from book.books where id_book = ?", bk.Id)

	fmt.Println(err)
}

//добавление книга из бд
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	database := db.Manager()
	rows, err := database.Query("select * from book.books")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	books := []Books{} // срез с типом Books

	for rows.Next() {
		b := Books{}
		err := rows.Scan(&b.Id, &b.Name, &b.Author, &b.Publication)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}
	js, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// //alternative
// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         var bk Books
//         if r.Body == nil {
//             http.Error(w, "Please send a request body", 400)
//             return
//         }
//         err := json.NewDecoder(r.Body).Decode(&bk)
//         if err != nil {
//             http.Error(w, err.Error(), 400)
//             return
//         }
//         fmt.Println(bk.Id)
//     })
//     log.Fatal(http.ListenAndServe(":8181", nil))
// }

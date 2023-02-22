package models

type Book struct {
	BOOKID int    `json:"bookid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

# Zocket Internship Assignment


## Problem Statement
TASK-1- Implement a basic CRUD API: The intern should create a basic CRUD (Create, Read, Update, Delete) API using Go that allows a user to manage a list of items (e.g., books, movies, etc.). The API should include endpoints for creating, reading, updating, and deleting items.


Implemented as follows 
Item - book
Database - MongoDB
Language - Go
Library - Gin Gonic

- View all books - GET /books
- View a book - GET /books?bookid=1 (bookid is the unique id of the book)
- Add a book - POST /books (book details in the body)
- Update a book - Patch /book (book details in the body)
- Delete a book - DELETE /book (bookid in the body)

createdatabase.ipynb - contains the code to create and populate the collection
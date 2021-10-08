package service

import (
	"backend/model"
)

type BookService struct {}

func (BookService) SetBook(book *model.Book) error {
	 _, err := DbEngine.Insert(book)
	 if err != nil {
		 return err
	 }
	 return nil
}

func (BookService) GetBookList() []model.Book {
	bookLists := make([]model.Book, 0)
	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&bookLists)
	if err != nil {
		panic(err)
	}
	return bookLists
}

func (BookService) UpdateBook(newBook *model.Book) error {
	 _, err := DbEngine.Id(newBook.Id).Update(newBook)
	 if err != nil {
		return err
	}
	return nil
}

func (BookService) DeleteBook(id int) error {
	book := new(model.Book)
	 _, err := DbEngine.Id(id).Delete(book)
	 if err != nil {
		return err
	}
	return nil
}
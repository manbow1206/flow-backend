package service

import (
	"backend/model"
)

type BookService struct {}

func (BookService) SetBook(book *model.Book) error {
	 _, err := DbEngin.Insert(book)
	 if err != nil {
		 return err
	 }
	 return nil
}

func (BookService) GetBookList() []model.Book {
	bookLists := make([]model.Book, 0)
	err := DbEngin.Distinct("id", "title", "content").Limit(10, 0).Find(&bookLists)
	if err != nil {
		return err
	}
	return bookLists
}

func (BookService) UpdateBook(newBoook *model.Book) error {
	 _, err := DbEngin.Id(newBoook.Id).Update(newBook)
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
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
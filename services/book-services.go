package services

import (
	"context"
	"fmt"
	"log"

	"github.com/dhruv-vavliya/BookStore/ent"
	"github.com/dhruv-vavliya/BookStore/models"
	"github.com/dhruv-vavliya/BookStore/repos"
)

func GetAllBooks(ctx context.Context, client *ent.Client) ([]*ent.Book, error) {
	books, err := repo.GetBooks(client, ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed Registering Book: %w", err)
	}
	return books, nil
}

func GetBooksByAuthorName(ctx context.Context, client *ent.Client, authorName string) ([]*ent.Book, error) {
	
	author, err := repo.GetAuthorByAuthorName(client, authorName, ctx)
	if err != nil {
		return nil, fmt.Errorf("Author Doesn't Exist: %w", err)
	}

	books, err := repo.GetBooksByAuthor(author, ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed Registering Book: %w", err)
	}

	return books, nil
}

func CreateBook(ctx context.Context, client *ent.Client, params *models.Book, author *ent.Author) (*ent.Book, error) {
	book, err := repo.CreateBook(client, params, author, ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed Registering Book: %w", err)
	}
	
	log.Println("Book Successfully Added to BookStore.")
	return book, nil
}


func DeleteBookByID(ctx context.Context, client *ent.Client) (error) {
	affected, err := repo.DeleteBook(client, ctx)
	if err != nil {
		return fmt.Errorf("Failed Deleting Book: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("Book Not Found: %w", err)
	}
	log.Println("Book Successfully Removed to BookStore.")
	return nil
}



func UpdateBookByID(ctx context.Context, client *ent.Client, bookID int, params *models.Book) (error) {
	affected, err := repo.UpdateBook(client, bookID, params, ctx)
	if err != nil {
		return fmt.Errorf("Failed Deleting Book: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("Book Not Found: %w", err)
	}
	log.Println("Book Successfully Updated.")
	return nil
}
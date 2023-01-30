package services

import (
	"context"
	"fmt"
	"log"

	"github.com/dhruv-vavliya/BookStore/ent"
	"github.com/dhruv-vavliya/BookStore/ent/author"
	"github.com/dhruv-vavliya/BookStore/ent/book"
	"github.com/dhruv-vavliya/BookStore/models"
)

func GetAllBooks(ctx context.Context, client *ent.Client) ([]*ent.Book, error) {
	books, err := client.Book.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed Registering Book: %w", err)
	}

	return books, nil
}

func GetBooksByAuthorName(ctx context.Context, client *ent.Client, authorName string) ([]*ent.Book, error) {
	
	author, err := client.Author.Query().Where(
		author.Name(authorName),
	).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("Author Doesn't Exist: %w", err)
	}

	books, err := author.QueryBooks().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed Registering Book: %w", err)
	}

	return books, nil
}

func CreateBook(ctx context.Context, client *ent.Client, params *models.Book, user *ent.Author) (*ent.Book, error) {
	book, err := client.Book.Create().
		SetName(params.Name).
		SetPrice(params.Price).
		SetAuthor(user).
		Save(ctx)						// save to DB & return to HTTP response.
	
	if err != nil {
		return nil, fmt.Errorf("Failed Registering Book: %w", err)
	}
	
	log.Println("Book Successfully Added to BookStore.")
	return book, nil
}


func DeleteBookByID(ctx context.Context, client *ent.Client, bookID int) (error) {
	affected, err := client.Book.Delete().
		Where(
			book.ID(bookID),
		).Exec(ctx)
	
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
	affected, err := client.Book.Update().
		Where(
			book.ID(bookID),
		).
		SetName(params.Name).
		SetPrice(params.Price).
		Save(ctx)
	
	if err != nil {
		return fmt.Errorf("Failed Deleting Book: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("Book Not Found: %w", err)
	}
	log.Println("Book Successfully Updated.")
	return nil
}
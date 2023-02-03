package repo

import (
	"context"
	"fmt"
	// "strconv"

	"github.com/dhruv-vavliya/BookStore/ent"
	"github.com/dhruv-vavliya/BookStore/ent/book"
	"github.com/dhruv-vavliya/BookStore/models"
)

type BookStoreAccess interface{
	CreateBook(client *ent.Client, params *models.Book, author *ent.Author, ctx context.Context) (*ent.Book, error)
	DeleteBook(client *ent.Client, bookID int, ctx context.Context) (int, error)
	UpdateBook(client *ent.Client, bookID int, params *models.Book, ctx context.Context) (int, error)
}

func CreateBook(client *ent.Client, params *models.Book, author *ent.Author, ctx context.Context) (*ent.Book, error) {
	return client.Book.Create().
		SetName(params.Name).
		SetPrice(params.Price).
		SetAuthor(author).
		Save(ctx)						// save to DB & return to HTTP response.
}

func DeleteBook(client *ent.Client, ctx context.Context) (int, error) {
	// bookID, _ := strconv.Atoi("bookID")
	fmt.Print(ctx)
	return client.Book.Delete().
		Where(
			book.ID(1),
		).
		Exec(ctx)
}

func UpdateBook(client *ent.Client, bookID int, params *models.Book, ctx context.Context) (int, error) {
	return client.Book.Update().
		Where(
			book.ID(bookID),
		).
		SetName(params.Name).
		SetPrice(params.Price).
		Save(ctx)
}



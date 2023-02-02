package repo

import (
	"context"

	"github.com/dhruv-vavliya/BookStore/ent"
	"github.com/dhruv-vavliya/BookStore/ent/author"
	"github.com/dhruv-vavliya/BookStore/models"
)

type AuthorAccess interface{
	CreateAuthor(client *ent.Client, params *models.Author, hashedPassword string, ctx context.Context) (*ent.Author, error)
	GetAuthorByEmail(client *ent.Client, params *models.Credentials, ctx context.Context) (*ent.Author, error)
	DeleteAuthor(client *ent.Client, authorID int, ctx context.Context) (int, error)
	GetBooks(client *ent.Client, ctx context.Context) ([]*ent.Book, error)
	GetAuthorByAuthorName(client *ent.Client, authorName string, ctx context.Context) (*ent.Author, error)
	GetBooksByAuthor(author *ent.Author, ctx context.Context) ([]*ent.Book, error)
}

func CreateAuthor(client *ent.Client, params *models.Author, hashedPassword string, ctx context.Context) (*ent.Author, error) {
	return client.Author.Create().
		SetEmail(params.Email).
		SetName(params.Name).
		SetPassword(hashedPassword).
		Save(ctx)						// save to DB & return to HTTP response.
}

func GetAuthorByEmail(client *ent.Client, params *models.Credentials, ctx context.Context) (*ent.Author, error) {
	return client.Author.Query().
		Where(
			author.Email(params.Email),
		).
		Only(ctx)
}


func DeleteAuthor(client *ent.Client, authorID int, ctx context.Context) (int, error) {
	return client.Author.Delete().
		Where(
			author.ID(authorID),
		).
		Exec(ctx)
}

func GetBooks(client *ent.Client, ctx context.Context) ([]*ent.Book, error) {
	return client.Book.Query().All(ctx)
}

func GetAuthorByAuthorName(client *ent.Client, authorName string, ctx context.Context) (*ent.Author, error) {
	return client.Author.Query().Where(
		author.Name(authorName),
	).First(ctx)
}

func GetBooksByAuthor(author *ent.Author, ctx context.Context) ([]*ent.Book, error) {
	return author.QueryBooks().All(ctx)
}
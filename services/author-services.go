package services

import (
	"context"
	"fmt"
	"log"

	"github.com/dhruv-vavliya/BookStore/ent"
	"github.com/dhruv-vavliya/BookStore/ent/author"
	"github.com/dhruv-vavliya/BookStore/models"
	"github.com/dhruv-vavliya/BookStore/utils"
)


func RegisterAuthor(ctx context.Context, client *ent.Client, params *models.Author ) (*ent.Author, error) {
	
	hashedPassword, _ := utils.HashPassword(params.Password)
	user, err := client.Author.Create().
		SetEmail(params.Email).
		SetName(params.Name).
		SetPassword(hashedPassword).
		Save(ctx)						// save to DB & return to HTTP response.
	
	if err != nil {
		return nil, fmt.Errorf("Failed Registering Author: %w", err)
	}
	log.Println("Author Created Successfully.")

	return user, nil
}

func ValidateAuthor(ctx context.Context, client *ent.Client, params *models.Credentials ) (int, error) {
	
	// validate user
	author, err := client.Author.Query().
	Where(
		author.Email(params.Email),
	).
	Only(ctx)
	if err != nil {
		return -1, err
	}

	// validate password
	if !utils.CheckPasswordHash(author.Password, params.Password) {
		return -1, fmt.Errorf("Wrong Password! please try again.")
	}
	return author.ID, nil
}



func DeleteAuthorByID(ctx context.Context, client *ent.Client, authorID int) (error) {
	affected, err := client.Author.Delete().
		Where(
			author.ID(authorID),
		).Exec(ctx)
		
	if err != nil {
		return fmt.Errorf("Failed Deleting Author: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("Author Not Found: %w", err)
	}
	log.Println("Author Deleted Successfully.")
	return nil
}



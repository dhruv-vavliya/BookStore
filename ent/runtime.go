// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dhruv-vavliya/BookStore/ent/author"
	"github.com/dhruv-vavliya/BookStore/ent/book"
	"github.com/dhruv-vavliya/BookStore/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authorFields := schema.Author{}.Fields()
	_ = authorFields
	// authorDescEmail is the schema descriptor for email field.
	authorDescEmail := authorFields[0].Descriptor()
	// author.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	author.EmailValidator = authorDescEmail.Validators[0].(func(string) error)
	// authorDescPassword is the schema descriptor for password field.
	authorDescPassword := authorFields[1].Descriptor()
	// author.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	author.PasswordValidator = authorDescPassword.Validators[0].(func(string) error)
	// authorDescName is the schema descriptor for name field.
	authorDescName := authorFields[2].Descriptor()
	// author.DefaultName holds the default value on creation for the name field.
	author.DefaultName = authorDescName.Default.(string)
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescName is the schema descriptor for name field.
	bookDescName := bookFields[0].Descriptor()
	// book.NameValidator is a validator for the "name" field. It is called by the builders before save.
	book.NameValidator = bookDescName.Validators[0].(func(string) error)
	// bookDescPrice is the schema descriptor for price field.
	bookDescPrice := bookFields[1].Descriptor()
	// book.DefaultPrice holds the default value on creation for the price field.
	book.DefaultPrice = bookDescPrice.Default.(int)
	// bookDescDate is the schema descriptor for date field.
	bookDescDate := bookFields[2].Descriptor()
	// book.DefaultDate holds the default value on creation for the date field.
	book.DefaultDate = bookDescDate.Default.(time.Time)
}
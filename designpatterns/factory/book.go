package factory
// book is the product that the client requires
// It doesnt matter to the client how the concrete implementation of the book is created
// It will just pass some identifier to tell which type of book it needs and the factory will create that
type Book interface{
  GetBookName() string
}
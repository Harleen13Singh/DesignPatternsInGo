package factory
// Concrete implementation of book interface
type IntellectualBook struct{}
func (I *IntellectualBook) GetBookName() string {
  return "Book of Mirdad"
}
// Concrete factory of Intellectual book type
func NewIntellectualBook() *IntellectualBook {
  return &IntellectualBook{}
}
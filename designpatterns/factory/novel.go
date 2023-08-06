package factory
// concrete implementation of Book interface
type Novel struct{}
func (N *Novel) GetBookName() string {
  return "Pavittar Paapi"
}
// Concrete factory for the Novel book type
func NewNovel() *Novel{
  return &Novel{}
}


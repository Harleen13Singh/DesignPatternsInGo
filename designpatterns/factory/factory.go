package factory

// factory struct that has all the concrete books
type BookFactory struct {
  novel *Novel
  intellectuaBook *IntellectualBook
}

// factory method for the BookFactory
func NewBookFactory(novel *Novel,ib *IntellectualBook) *BookFactory{
  return &BookFactory{
    novel: novel,
    intellectuaBook: ib,
  }
} 

// Method based upon which concreted book will be returned
func (cf *BookFactory) FetchConcreteObject(randomnesFactor int) Book{
  if randomnesFactor < 100{
    return cf.intellectuaBook
  }
  return cf.novel
}

package factory
import "fmt"

func ClientCode() {
  ib := NewIntellectualBook()
	n := NewNovel()
	cf := NewBookFactory(n,ib)
	book := cf.FetchConcreteObject(24)

	fmt.Println(book.GetBookName())
	book = cf.FetchConcreteObject(1222)
	fmt.Println(book.GetBookName())
}
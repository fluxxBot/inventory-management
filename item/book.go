package item

type Book struct {
	Title string
	id    string
}

func (b Book) GetName() string {
	return b.Title
}

func (b Book) GetId() string {
	return b.id
}

func (b Book) Create(id string) any {
	return Book{Title: b.GetName(), id: id}
}

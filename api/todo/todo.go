package todo

type Todo struct {
	ID          uint `gorm: "primaryKey;unique"`
	Title       string
	Description string
}

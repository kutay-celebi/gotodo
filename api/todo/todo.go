package todo

type Todo struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4(); unique" json:"id"`
	Title       string `validate:"not_empty;min=15" json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

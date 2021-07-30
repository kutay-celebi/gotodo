package todo

import "github.com/google/uuid"

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4(); unique" json:"id"`
	Title       string    `validate:"not_empty;min=15" json:"title"`
	Description string    `json:"description"`
}

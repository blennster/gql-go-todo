package model

type Todo struct {
	ID      string `json:"id" db:"id"`
	Text    string `json:"text" db:"text"`
	Done    bool   `json:"done" db:"done"`
	UserRef string `json:"userref" db:"userref"`
}

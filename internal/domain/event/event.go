package event

type Event struct {
	ID          string `db:"id,omitempty" json:"id,omitempty"`
	Name        string `db:"name,omitempty" json:"name,omitempty"`
	Description string `db:"description,omitempty" json:"description,omitempty"`
	DateAndTime string `db:"date_and_time,omitempty" json:"date_and_time,omitempty"`
}

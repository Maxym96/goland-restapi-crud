package event

type Event struct {
	ID          string `db:"id,omitempty"`
	Name        string `db:"name,omitempty"`
	Description string `db:"description,omitempty"`
	DateAndTime string `db:"date_and_time,omitempty"`
}

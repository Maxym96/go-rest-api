package event

type Event struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	DateAndTime string `db:"date_and_time"`
}

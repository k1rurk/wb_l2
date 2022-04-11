package event

import (
	"encoding/json"
	"net/http"
	"time"
)

type Event struct {
	EventID int    `json:"event_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Data    Time   `json:"date"`
}

func UnmarshalRequest(r *http.Request) (Event, error) {
	var calendar Event
	err := json.NewDecoder(r.Body).Decode(&calendar)
	if err != nil {
		return Event{}, err
	}
	defer r.Body.Close()
	return calendar, nil
}

type Time struct {
	time.Time
}

const LayoutIOS = "2006-01-02"

func (t *Time) UnmarshalJSON(byte []byte) error {
	date, err := time.Parse("\""+LayoutIOS+"\"", string(byte))
	if err != nil {
		return err
	}

	*t = Time{date}
	return nil
}

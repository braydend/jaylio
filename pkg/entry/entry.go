package entry

import (
	"encoding/json"
	"time"
)

type entry struct {
	Date       time.Time `json:"date"`
	Mood       string   `json:"mood"`
	Activities []string `json:"activities"`
	NoteTitle string `json:"note_title"`
	Note      string `json:"note"`
}

type Entries struct {
	Entries []entry `json:"entries"`
}

func ParseEntriesFromCSV(csv []byte) Entries {
	return Entries{}
}

func (e Entries) toString() (string, error) {
	serialised, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(serialised), nil
}

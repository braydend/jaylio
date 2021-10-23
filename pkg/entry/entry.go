package entry

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type entry struct {
	Date       time.Time `json:"date"`
	Mood       string    `json:"mood"`
	Activities []string  `json:"activities"`
	NoteTitle  string    `json:"note_title"`
	Note       string    `json:"note"`
}

type Entries struct {
	Entries []entry `json:"input"`
}

var (
	CsvHeaderIndex   = 0
	FullDateIndex    = 0
	TimeIndex        = 3
	MoodIndex        = 4
	ActivitiesIndex  = 5
	NoteTitleIndex   = 6
	NoteContentIndex = 7
)

func ParseEntriesFromCSV(csvData []byte) (Entries, error) {
	records, err := csv.NewReader(bytes.NewReader(csvData)).ReadAll()

	if err != nil {
		return Entries{}, err
	}

	var entries []entry

	for i, record := range records {
		if i != CsvHeaderIndex {
			date, err := parseDateFromEntry(record[FullDateIndex], record[TimeIndex])

			if err != nil {
				return Entries{}, nil
			}

			entries = append(entries, entry{
				Date:       date,
				Mood:       record[MoodIndex],
				Activities: parseActivitiesFromEntry(record[ActivitiesIndex]),
				NoteTitle:  record[NoteTitleIndex],
				Note:       record[NoteContentIndex],
			})
		}
	}

	return Entries{Entries: entries}, nil
}

func parseDateFromEntry(date string, timeRecorded string) (time.Time, error) {
	dateStrings := strings.Split(date, "-")
	year, err := strconv.Atoi(dateStrings[0])
	month, err := strconv.Atoi(dateStrings[1])
	day, err := strconv.Atoi(dateStrings[2])

	isAm := strings.Contains(timeRecorded, "am")
	var hour, minute int
	if isAm {
		timeStrings := strings.Split(strings.Replace(timeRecorded, " am", "", 1), ":")
		hour, err = strconv.Atoi(timeStrings[0])
		minute, err = strconv.Atoi(timeStrings[1])
	} else {
		timeStrings := strings.Split(strings.Replace(timeRecorded, " pm", "", 1), ":")
		hour, err = strconv.Atoi(timeStrings[0])
		hour = (hour + 12) % 24
		minute, err = strconv.Atoi(timeStrings[1])
	}

	if err != nil {
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local), nil
}

func parseActivitiesFromEntry(activities string) []string {
	if activities == "" {
		return []string{}
	}
	activitySplits := strings.Split(activities, "|")
	var trimmedActivities []string

	for _, activity := range activitySplits {
		trimmedActivities = append(trimmedActivities, strings.TrimSpace(activity))
	}

	return trimmedActivities
}

func (e Entries) toString() (string, error) {
	serialised, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(serialised), nil
}

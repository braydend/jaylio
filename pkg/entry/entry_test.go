package entry

import (
	"fmt"
	"testing"
	"time"
)

func TestParseEntriesFromCSV(t *testing.T) {
	t.Run("converts CSV correctly", func(t *testing.T) {
		amRecordCsv := []byte("full_date,Date,weekday,time,Mood,Activities,note_title,Note\n2021-10-18,18 October,Monday,11:25 am,Energised,\"home | appreciate\",\"\",\"\"")
		amRecords := Entries{[]entry{{time.Date(2021,10,18,11,25, 0, 0, time.Local), "Energised", []string{"home", "appreciate"}, "", ""}}}

		pmRecordCsv := []byte("full_date,Date,weekday,time,Mood,Activities,note_title,Note\n2021-10-18,18 October,Monday,11:25 pm,Energised,\"home | appreciate\",\"\",\"\"")
		pmRecords := Entries{[]entry{{time.Date(2021,10,18,23,25, 0, 0, time.Local), "Energised", []string{"home", "appreciate"}, "", ""}}}

		noActivityRecordsCsv := []byte("full_date,Date,weekday,time,Mood,Activities,note_title,Note\n2021-10-18,18 October,Monday,11:25 pm,Energised,\"\",\"\",\"\"")
		noActivityRecords := Entries{[]entry{{time.Date(2021,10,18,23,25, 0, 0, time.Local), "Energised", []string{}, "", ""}}}

		noteRecordsCsv := []byte("full_date,Date,weekday,time,Mood,Activities,note_title,Note\n2021-10-18,18 October,Monday,11:25 pm,Energised,\"home | appreciate\",\"\",\"Today I wrote some tests\"")
		noteRecords := Entries{[]entry{{time.Date(2021,10,18,23,25, 0, 0, time.Local), "Energised", []string{"home", "appreciate"}, "", "Today I wrote some tests"}}}

		tests := []struct {
			name     string
			input    []byte
			expected Entries
		}{
			{"parses entry recorded in am", amRecordCsv, amRecords},
			{"parses entry recorded in pm", pmRecordCsv, pmRecords},
			{"parses entry with no activities", noActivityRecordsCsv, noActivityRecords},
			{"parses entry with a note", noteRecordsCsv, noteRecords},
		}
		for _, testcase := range tests {
			t.Run(testcase.name, func(t *testing.T) {
				result, _ := ParseEntriesFromCSV(testcase.input)
				expectedString, _ := testcase.expected.toString()
				if got, _ := result.toString(); got != expectedString {
					t.Errorf("recieved %v\n expected %v", result, testcase.expected)
				}
			})
		}
	})
}

func TestEntries_toString(t *testing.T) {
	date := time.Date(2021,10,18,11,25, 0, 0, time.Local)
	withActivities := Entries{[]entry{{date, "Energized", []string{"home", "appreciate"}, "", ""}}}
	withoutActivities := Entries{[]entry{{date, "Energized", []string{}, "", ""}}}
	withNote := Entries{[]entry{{date, "Energized", []string{}, "Note's title", "Note's content"}}}

	dateString, _ := date.MarshalJSON()
	withActivitiesString := fmt.Sprintf(`{"input":[{"date":%s,"mood":"Energized","activities":["home","appreciate"],"note_title":"","note":""}]}`, dateString)
	withoutActivitiesString := fmt.Sprintf(`{"input":[{"date":%s,"mood":"Energized","activities":[],"note_title":"","note":""}]}`, dateString)
	withNoteString := fmt.Sprintf(`{"input":[{"date":%s,"mood":"Energized","activities":[],"note_title":"Note's title","note":"Note's content"}]}`, dateString)

	tests := []struct {
		name string
		json Entries
		want string
	}{
		{"serialises to a string", withActivities, withActivitiesString},
		{"serialises to a string without Activities", withoutActivities, withoutActivitiesString},
		{"serialises to a string with Note", withNote, withNoteString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.json
			if got, _ := d.toString(); got != tt.want {
				t.Errorf("toString() = %v, expected %v", got, tt.want)
			}
		})
	}
}

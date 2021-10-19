package entry

import (
	"fmt"
	"testing"
	"time"
)

func TestParseCSV(t *testing.T) {
	t.Run("converts CSV correctly", func(t *testing.T) {
		input := []byte("full_date,Date,weekday,time,Mood,Activities,note_title,Note\n2021-10-18,18 October,Monday,11:25 pm,Energised,\"home | appreciate\",\"\",\"\"")
		expected := Entries{Entries: []entry{
			{
				Date:       time.Date(2021,10,18,11,25, 0, 0, time.Local),
				Mood:       "Energized",
				Activities: []string{"home", "appreciate"},
				NoteTitle:  "",
				Note:       "",
			},
		}}
		result := ParseEntriesFromCSV(input)

		expectedString, _ := expected.toString()
		resultString, _ := result.toString()

		if expectedString != resultString {
			t.Fatalf("expected: %s\ngot: %s\n", expectedString, resultString)
		}
	})
}

func TestEntries_toString(t *testing.T) {
	date := time.Date(2021,10,18,11,25, 0, 0, time.Local)
	withActivities := Entries{Entries: []entry{
		{date, "Energized", []string{"home", "appreciate"}, "", ""},
	}}
	withoutActivities := Entries{Entries: []entry{
		{date, "Energized", []string{}, "", ""},
	}}
	withNote := Entries{Entries: []entry{
		{date, "Energized", []string{}, "Note's title", "Note's content"},
	}}

	dateString, _ := date.MarshalJSON()
	withActivitiesString := fmt.Sprintf(`{"entries":[{"date":%s,"mood":"Energized","activities":["home","appreciate"],"note_title":"","note":""}]}`, dateString)
	withoutActivitiesString := fmt.Sprintf(`{"entries":[{"date":%s,"mood":"Energized","activities":[],"note_title":"","note":""}]}`, dateString)
	withNoteString := fmt.Sprintf(`{"entries":[{"date":%s,"mood":"Energized","activities":[],"note_title":"Note's title","note":"Note's content"}]}`, dateString)

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
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
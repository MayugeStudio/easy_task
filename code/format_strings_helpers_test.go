package code

import "testing"

func TestGetStatusString(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"ValidTaskStringGoodFormat_Done":   {"- [ ] TaskName", " ", false},
		"ValidTaskStringGoodFormat_Undone": {"- [X] TaskName", "X", false},
		"ValidTaskStringBadFormat_Done":    {"-[]TaskName", " ", false},
		"ValidTaskStringBadFormat_Undone":  {"-[X]TaskName", "X", false},
		// Error cases
		"InValidTaskString_NoDash":         {"[ ] TaskName", "", true},
		"InValidTaskString_NoBracketStart": {"- ] TaskName", "", true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := GetStatusString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatusString() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStatusString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		// Success cases
		"ValidGroupStringGoodFormat": {"- GroupTitle", "GroupTitle", false},
		"ValidGroupStringBadFormat":  {"-GroupTitle", "GroupTitle", false},
		// Error cases
		"InvalidGroupString_NoDash": {"GroupTitle", "", true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := GetGroupTitle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetGroupTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ValidGroupTitleString":   {"- GroupTitle", true},
		"InvalidGroupTitleString": {"GroupTitle", false},
		"TaskString":              {"- [ ] TaskName", false},
		"GroupTaskString":         {"  - [ ] TaskName", false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := IsGroupTitle(tt.in)
			if got != tt.want {
				t.Errorf("IsGroupTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGroupTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ValidGroupTaskString":                  {"  - [ ] TaskName", true},
		"InvalidGroupTaskString_NoDash":         {"  [ ] TaskName", false},
		"InvalidGroupTaskString_NoBracketStart": {"  -  ] TaskName", false},
		"InvalidGroupTaskString_NoBracketEnd":   {"  - [  TaskName", false},
		"SingleTaskString":                      {"- [ ] TaskName", false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := IsGroupTaskString(tt.in)
			if got != tt.want {
				t.Errorf("IsGroupTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

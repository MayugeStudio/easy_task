package code

import "testing"

func TestGetStatusString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want string
	}{
		"ValidTaskStringGoodFormat_Done":   {"- [ ] TaskName", " "},
		"ValidTaskStringGoodFormat_Undone": {"- [X] TaskName", "X"},
		"ValidTaskStringBadFormat_Done":    {"-[]TaskName", " "},
		"ValidTaskStringBadFormat_Undone":  {"-[X]TaskName", "X"},
		"InValidTaskString_NoDash":         {"[ ] TaskName", ""},
		"InValidTaskString_NoBracketStart": {"- ] TaskName", ""},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := GetStatusString(tt.in)
			if got != tt.want {
				t.Errorf("GetStatusString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in   string
		want string
	}{
		"ValidGroupStringGoodFormat": {"- GroupTitle", "GroupTitle"},
		"ValidGroupStringBadFormat":  {"-GroupTitle", "GroupTitle"},
		"InvalidGroupString_NoDash":  {"GroupTitle", ""},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := GetGroupTitle(tt.in)
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

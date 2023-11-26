package share

import "testing"

func Test_isGroupTitle(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ValidGroupTitleString":   {in: "- GroupTitle", want: true},
		"InvalidGroupTitleString": {in: "GroupTitle", want: false},
		"TaskString":              {in: "- [ ] TaskName", want: false},
		"GroupTaskString":         {in: "  - [ ] TaskName", want: false},
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

func Test_isGroupTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ValidGroupTaskString":                  {in: "  - [ ] TaskName", want: true},
		"InvalidGroupTaskString_NoDash":         {in: "  [ ] TaskName", want: false},
		"InvalidGroupTaskString_NoBracketStart": {in: "  -  ] TaskName", want: false},
		"InvalidGroupTaskString_NoBracketEnd":   {in: "  - [  TaskName", want: false},
		"SingleTaskString":                      {in: "- [ ] TaskName", want: false},
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

func Test_isSingleTaskString(t *testing.T) {
	tests := map[string]struct {
		in   string
		want bool
	}{
		"SingleTask_Done":               {in: "- [ ] I am SingleTask!", want: true},
		"SingleTask_Undone":             {in: "- [X] I am SingleTask!", want: true},
		"GroupTask_Done":                {in: "  - [X] I am SingleTask!", want: false},
		"GroupTask_Undone":              {in: "  - [X] I am SingleTask!", want: false},
		"Invalid_Undone_NoBracketStart": {in: "- X] I am SingleTask!", want: false},
		"Invalid_Undone_NoBracketEnd":   {in: "- [X I am SingleTask!", want: false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got := IsSingleTaskString(tt.in)
			if got != tt.want {
				t.Errorf("IsSingleTaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package share

import "testing"

func Test_isGroupTitle(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ValidGroupTitleString":    {in: "- GroupTitle", want: true},
		"IndentedGroupTitleString": {in: "  - GroupTitle", want: true},
		"InvalidGroupTitleString":  {in: "GroupTitle", want: false},
		"TaskString":               {in: "- [ ] TaskName", want: false},
		"GroupTaskString":          {in: "  - [ ] TaskName", want: false},
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
	t.Parallel()
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
	t.Parallel()
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

func TestIsItemModificationString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   string
		want bool
	}{
		"ModificationString":         {in: "<- [Tag]: Feature", want: true},
		"Invalid_ModificationString": {in: "[Tag]: Feature", want: false},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := IsItemModificationString(tt.in); got != tt.want {
				t.Errorf("IsItemModificationString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndentLevel(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   string
		want int
	}{
		"0": {in: "Hello", want: 0},
		"2": {in: "  Hello", want: 2},
		"4": {in: "    Hello", want: 4},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := GetIndentLevel(tt.in); got != tt.want {
				t.Errorf("GetIndentLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

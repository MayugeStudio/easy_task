package format

import "testing"

func Test_toFormattedGroupTaskString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Valid":         {in: "  - [ ] Buy the milk.", want: "  - [ ] Buy the milk."},
		"OneIndent":     {in: " - [ ] Buy the milk.", want: "  - [ ] Buy the milk."},
		"ThreeIndent":   {in: "   - [ ] Buy the milk.", want: "    - [ ] Buy the milk."},
		"NoIndent":      {in: "- [ ] Buy the milk.", wantErr: true},
		"InvalidFormat": {in: "  - Buy the milk.", wantErr: true},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedGroupTaskString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTaskString() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatGroupTaskString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_toFormattedGroupTitle(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Valid":       {in: "- GroupTitle", want: "- GroupTitle"},
		"BadIndent":   {in: "-GroupTitle", want: "- GroupTitle"},
		"InvalidLine": {in: "GroupTitle", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedGroupTitle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGroupTitleString() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatGroupTitleString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractGroupTitle(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"ValidGroupStringGoodFormat": {in: "- GroupTitle", want: "GroupTitle"},
		"ValidGroupStringBadFormat":  {in: "-GroupTitle", want: "GroupTitle"},
		"InvalidGroupString_NoDash":  {in: "GroupTitle", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := extractGroupTitle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetGroupTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

package format

import "testing"

func Test_toFormattedModificationString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"Valid_NoIndent":              {in: "<- [Tag]: Feature", want: "<- [Tag]: Feature"},
		"Valid_4Indent":               {in: "    <- [Tag]: Feature", want: "    <- [Tag]: Feature"},
		"Valid_BadFormat_NoIndent":    {in: "<-[Tag] :Feature", want: "<- [Tag]: Feature"},
		"Valid_BadFormat_4Indent":     {in: "    <-[Tag] :Feature", want: "    <- [Tag]: Feature"},
		"Invalid_InvalidModification": {in: "[Tag]: Feature", wantErr: true},
		"Invalid_NoBracketStart":      {in: "<- Tag]: Feature", wantErr: true},
		"Invalid_NoBracketEnd":        {in: "<- [Tag: Feature", wantErr: true},
		"Invalid_NoColon":             {in: "<- [Tag] Feature", wantErr: true},
		"Invalid_InvalidAttribute":    {in: "<- [tag]: Feature", wantErr: true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := toFormattedModificationString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("toFormattedModificationString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toFormattedModificationString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

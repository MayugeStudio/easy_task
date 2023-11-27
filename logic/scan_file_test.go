package logic

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

type MockFileReader struct {
	ReadLinesFunc func(filename string) (lines []string, err error)
}

func (m MockFileReader) ReadLines(filename string) (lines []string, err error) {
	return m.ReadLinesFunc(filename)
}

func TestScanFile(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		fileName string
		reader   FileReader
		want     []string
		wantErr  bool
	}{
		"SuccessfulScan": {
			fileName: "testFile.txt",
			reader: MockFileReader{
				func(filename string) (lines []string, err error) {
					return []string{"line1", "line2"}, nil
				},
			},
			want:    []string{"line1", "line2"},
			wantErr: false,
		},
		"ErrorOpeningFile": {
			fileName: "nonexistent.txt",
			reader: MockFileReader{
				func(filename string) (lines []string, err error) {
					return nil, errors.New("file not found")
				},
			},
			want:    nil,
			wantErr: true,
		},
		"ErrorScanningFile": {
			fileName: "errorFile.txt",
			reader: MockFileReader{
				func(filename string) (lines []string, err error) {
					return nil, errors.New("error scanning file")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			lines, err := ScanFile(tc.fileName, tc.reader)

			if (err != nil) != tc.wantErr {
				t.Errorf("ScanFile() error = %v, wantErr = %v", err, tc.wantErr)
				return
			}

			if len(lines) != len(tc.want) {
				t.Errorf("ScanFile() lines length = %d, want %d", len(lines), len(tc.want))
			}
			for i, line := range lines {
				if line != tc.want[i] {
					t.Errorf("ScanFile() lines[%d] = %s, want %s", i, line, tc.want[i])
				}
			}
		})
	}
}

func TestFileScanner_ReadLines(t *testing.T) {
	t.Parallel()
	var fileName = "temp.txt"
	f, createErr := os.CreateTemp("", fileName)
	if createErr != nil {
		t.Errorf("CreateTemp() error = %v", createErr)
		return
	}
	if f == nil {
		t.Errorf("CreateTemp() returns nil")
		return
	}

	defer func(name string) {
		_ = os.Remove(f.Name())
	}(f.Name())

	data := "line1\nline2\nline3"
	if _, writeErr := f.Write([]byte(data)); writeErr != nil {
		t.Errorf("WriteString() error = %v", writeErr)
		return
	}

	tests := map[string]struct {
		fileName  string
		wantLines []string
		wantErr   bool
	}{
		"SuccessfulRead":  {fileName: f.Name(), wantLines: []string{"line1", "line2", "line3"}, wantErr: false},
		"NonexistentFile": {fileName: "nonexistent.txt", wantLines: nil, wantErr: true},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			fs := FileScanner{}

			gotLines, err := fs.ReadLines(tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("ReadLines() gotLines = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}

package code

import (
	"io"
)

// PrintErrorMessages prints error messages to the provided io.Writer.
// It returns an error if writing to the io.Writer fails.
func PrintErrorMessages(w io.Writer, messages []string) error {
	if len(messages) == 0 {
		return nil
	}

	for _, message := range messages {
		_, err := w.Write([]byte(message))
		if err != nil {
			return err
		}
	}

	_, err := w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
}

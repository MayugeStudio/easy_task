package code

type TestWriter struct {
	WrittenData []byte
}

func (tw *TestWriter) Write(p []byte) (n int, err error) {
	tw.WrittenData = append(tw.WrittenData, p...)
	return len(p), nil
}

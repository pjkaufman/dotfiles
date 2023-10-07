//go:build unit

package logger

type LoggerHandlerMock struct {
	InfoMsgs []string
	WarnMsgs []string
}

func NewMockLoggerHandler() *LoggerHandlerMock {
	return &LoggerHandlerMock{}
}

func (lg LoggerHandlerMock) WriteError(msg string) {
	panic(msg)
}

func (lg LoggerHandlerMock) WriteInfo(msg string) {
	lg.InfoMsgs = append(lg.InfoMsgs, msg)
}

func (lg LoggerHandlerMock) WriteWarn(msg string) {
	lg.WarnMsgs = append(lg.InfoMsgs, msg)
}

func (lg LoggerHandlerMock) GetInputString(prompt string) string {
	return ""
}

func (lg LoggerHandlerMock) GetInputInt(prompt string) int {
	return 0
}

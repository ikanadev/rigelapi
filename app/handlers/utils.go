package handlers

type errMsg struct {
  message string `json:"message"`
}
func newErrMsg(msg string) errMsg {
  return errMsg{msg}
}

package handlers

type ClientErr struct {
  Status int
  Message string
}
func (c ClientErr) Error() string {
  return c.Message
}
func NewClientErr(code int, msg string) ClientErr {
  return ClientErr{code, msg}
}

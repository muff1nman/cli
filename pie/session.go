package pie
import (
  "errors"
  "github.com/jmcvetta/napping"
)

type LoginReq struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

type Session struct {
  Token string `json:"token"`
  UserId int `json:"user_id"`
}

const (
  LOGIN_URL = URL_PREFIX + "/sessions"
)

func Login(email string, password string) (session *Session, err error) {
  payload := &LoginReq {
    Email: email,
    Password: password,
  }
  session = &Session{}
  resp, err := napping.Post(LOGIN_URL, payload, session, nil)
  if err != nil { return }
  if resp.Status() != 201 {
    err = errors.New("Wrong e-mail or password")
    return
  }
  return
}

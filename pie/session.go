package pie
import (
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
  LOGIN_URL = "/sessions"
)

func Login(email string, password string) (session *Session, err error) {
  payload := &LoginReq {
    Email: email,
    Password: password,
  }
  session = &Session{}
  req := &PiePostRequest{
    Url: LOGIN_URL,
    Payload: payload,
  }

  err = PostPieResource(req, session)
  return
}

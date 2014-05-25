package pie
import (
)

type loginReq struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

type Session struct {
  Token string `json:"token"`
  UserId int `json:"user_id"`
}

// Creates a new session. If the e-mail or password is wrong, it will return an error.
func Login(email string, password string) (session *Session, err error) {
  payload := &loginReq {
    Email: email,
    Password: password,
  }
  session = &Session{}
  req := &piePostRequest{
    Url: "/sessions",
    Payload: payload,
  }

  err = postPieResource(req, session)
  return
}

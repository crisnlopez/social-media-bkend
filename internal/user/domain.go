package user


type User struct {
  ID        int    `json:"id,omitempty"`
  Email     string `json:"email"`
  Pass      string `json:"pass"`
  Nick      string `json:"nick"`
  Name      string `json:"name"`
  Age       int    `json:"age"`
}

type UserUpdated struct {
  Email string `json:"email"`
  Pass  string `json:"pass"`
  Nick  string `json:"nick"`
  Name  string `json:"name"`
  Age   int    `json:"age"`
}


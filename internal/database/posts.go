package database

type Posts struct {
  ID string `json:"id"`
  UserEmail string `json:"userEmail"`
  Text string `json:"text"`
}

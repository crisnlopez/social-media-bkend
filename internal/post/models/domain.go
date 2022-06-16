package models

import "time"

type Post struct {
  ID int64 `json:"id"`
  UserID int64 `json:"user_id"`
  Text string `json:"text"`
  CreatedAt time.Time `json:"created_at"`
  UpdateAt time.Time `json:"update_at"`
}

type PostCreateRequest struct {
  UserID int64 `json:"user_id"`
  Text string `json:"text"`
}

type PostUpdateRequest struct {
  ID int64 `json:"id"`
  UserID int64 `json:"user_id"`
  Text string `json:"text"`
}

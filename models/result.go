package models

type (
  Recs interface{}
  Result struct {
    Error bool
    Msg   string
    Data  []Recs
  }
)
package services

import (
  "github.com/astaxie/beego"
  "github.com/parnurzeal/gorequest"
  "time"

  "cola/models"
)

func Read(q string, skip, limit int) (err error, rtvs []models.GoogleResult) {
  request := gorequest.New().Timeout(5000 * time.Millisecond)
  resp, _, errs := request.Get("https://www.google.com.hk/webhp?hl=en").End()
  if len(errs) != 0 {
    err = errs[0]
    return
  }

  beego.Info(resp.Body)

  return
}

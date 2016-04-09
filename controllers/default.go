package controllers

import (
	"github.com/astaxie/beego"
  "cola/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Welcome() {
  var rt models.Result

  rt.Error = false
  rt.Msg = "Successful"
  rt.Data = make([]models.Recs, 1)
  rt.Data[0] = "欢迎来到cola，这是beego从heroku给您返回的数据"

  this.Data["json"] = &rt
  this.ServeJSON()
}

func (this *MainController) Read() {
  var rt models.Result
  var q = this.GetString(":q")
  var skip = this.GetString(":skip")
  var limit = this.GetString(":limit")
  beego.Info(q)
  beego.Info(skip)
  beego.Info(limit)

  err, rtv := services.Read(q, skip, limit)
  if err != nil {
    rt.Error = true
    rt.Msg = err.Error()
  } else {
    rt.Error = false
    rt.Msg = "Successful"
  }
  this.Data["json"] = &rt
  this.ServeJSON()
}

package api

import (
	"encoding/json"
	"html/template"
	"net/http"
	"regexp"
	"time"
	"video_server/config"
	"video_server/gee"
	"video_server/models"
	"video_server/sessiongo"
)


func init()  {
	http.HandleFunc("/user_login", UserLoginController)
	http.HandleFunc("/user_signIn", UserSignInController)
}

func UserLoginController(c *gee.Context)   {
	// 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	err := r.ParseForm()
	if err != nil {
		context.Logger.Error(weberrors.Wrap(err))
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles(config.View_path + "login.html")
		if err != nil {
			context.Logger.Error(weberrors.Wrap(err))
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			context.Logger.Error(weberrors.Wrap(err))
			return
		}
	} else if r.Method == "POST" {
		respStruct := ResponseStruct{}
		err := handleLogin(w, r)
		if err != nil {
			respStruct.Status = ERROR
			errType := weberrors.Type(err)
			if errType == weberrors.UNKNOW_ERROR {	//未知错误，记录日志
				context.Logger.Error(err.Error())
				respStruct.Data = "system error"
				respStruct.Code = 0
			}else {	//否则只需返回错误信息
				respStruct.Data = err.Error()
				respStruct.Code = weberrors.Type(err)
			}
		}else {
			respStruct.Status = OK
			respStruct.Data = ""
			respStruct.Code = 0
		}
		data, err := json.Marshal(respStruct)
		if err != nil {
			context.Logger.Error(weberrors.Wrap(err).Error())
		}
		_, err = w.Write([]byte(data))
		if err != nil {
			context.Logger.Error(weberrors.Wrap(err))
		}
	}
	err = r.Body.Close()
	if err != nil {
		context.Logger.Error(weberrors.Wrap(err))
	}
}

func UserSignInController(w http.ResponseWriter,r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		context.Logger.Error(weberrors.Wrap(err))
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles(config.View_path + "signIn.html")
		if err != nil {
			context.Logger.Error(weberrors.Wrap(err))
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			context.Logger.Error(weberrors.Wrap(err))
			return
		}
	} else if r.Method == "POST" {
		context.Logger.Debug("signIn post")
		err := handleSignIn(w, r)
		if err != nil {
			if weberrors.Type(err) == weberrors.UNKNOW_ERROR {	//未知错误，记录日志
				context.Logger.Error(err)
			}else {	//否则只需返回错误信息
				data := make(map[string]string)
				data["error"] = err.Error()
				t, err := template.ParseFiles(config.View_path + "error.html")
				if err != nil {
					context.Logger.Error(weberrors.Wrap(err))
					return
				}
				err = t.Execute(w, data)
				if err != nil {
					context.Logger.Error(weberrors.Wrap(err))
					return
				}
			}
		}else {		//注册成功
			data := make(map[string]string)
			data["name"] = r.Form.Get("name")
			t, err := template.ParseFiles(config.View_path + "signInSuccess.html")
			if err != nil {
				context.Logger.Error(weberrors.Wrap(err))
				return
			}
			err = t.Execute(w, data)
			if err != nil {
				context.Logger.Error(weberrors.Wrap(err))
				return
			}
		}
	}
	err = r.Body.Close()
	if err != nil {
		context.Logger.Error(weberrors.Wrap(err))
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request)  error{
	name := r.Form.Get("name")
	pwd := r.Form.Get("pwd")
	err := validateLogin(name, pwd)
	if err != nil {
		return err
	}
	//验证密码是否正确
	user, err := models.GetUserByName(name)
	if err != nil {
		return weberrors.Wrap(err)
	}
	if user == nil {
		return weberrors.New("user is not exist", 1)
	}
	if user.Pwd != pwd {
		return weberrors.New("pwd error", 2)
	}
	//验证通过，创建session
	//给登录的用户简历session
	session, err := sessiongo.NewSession(time.Minute * 10) //10分钟过期
	if err != nil {
		return err
	}
	//设置cookie
	expire := time.Now().Add(time.Minute * 10)	//过期时间
	cookie := http.Cookie{Name: "SessionId", Value: session.GetId(), Expires: expire}
	http.SetCookie(w, &cookie)
	return nil
}

func handleSignIn(w http.ResponseWriter, r *http.Request)  error{
	name := r.Form.Get("name")
	pwd := r.Form.Get("pwd")
	err := validateLogin(name, pwd)
	context.Logger.Debug("signIn : name="+name+"pwd="+pwd)
	if err != nil {
		return weberrors.Wrap(err)
	}
	//验证通过，存入数据库
	err = models.AddUser(name, pwd)
	if err != nil {
		return weberrors.Wrap(err)
	}
	return nil
}

func validateLogin(name, pwd string) error {
	if len(name) == 0 || len(pwd) == 0 {
		return weberrors.New("name or pwd can not be empty", weberrors.PARA_ERROR)
	}
	b, err := regexp.MatchString("^[a-zA-Z0-9_]+$", name)
	if err != nil {
		return weberrors.New(err.Error(), weberrors.UNKNOW_ERROR)
	}
	if !b {	//不匹配
		return weberrors.New("name have illegal character", weberrors.PARA_ERROR)
	}
	b, err = regexp.MatchString("^[a-zA-Z0-9_]+$", name)
	if err != nil {
		return weberrors.New(err.Error(), weberrors.UNKNOW_ERROR)
	}
	if !b {	//不匹配
		return weberrors.New("pwd have illegal character", weberrors.PARA_ERROR)
	}
	return nil
}



package controllers

import (
	"Go_authentication/authtutorial/models"
	// "authtutorial/utils"
	// "github.com/OneOfOne/go-utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
)
var sessionName = beego.AppConfig.String("SessionName")

type LoginController struct {
	beego.Controller
}

func (this *LoginController) RegisterView() {

	this.TplNames = "register.html"
}

func (this *LoginController) Register() {
	username := this.GetString("username")
	password := this.GetString("password")
	passwordre := this.GetString("passwordre")
	fmt.Println(username)
	test := models.RegisterForm{Username: username, Password: password, PasswordRe: passwordre}

	valid := validation.Validation{}
	b, err := valid.Valid(&test)
	if err != nil {
	}
	if !b {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
		}
	} else {
		// salt := utils.GetRandomString(10)
		// encodedPwd := salt + "$" + utils.EncodePassword(password, salt)

		o := orm.NewOrm()
		o.Using("default")
		user := models.User{Username: username, Password:password}
		fmt.Printf(user.Username)
		// user.Password = *****
		// user.Rands = 'ssss'
		id, err := o.Insert(&user)

		fmt.Printf("ID: %d, ERR: %v\n", id, err)
		this.Redirect("/", 302)
	}
	this.TplNames = "register.html"
}
func (this *LoginController) LoginView() {

	this.TplNames = "login.html"
}
func (this *LoginController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	o := orm.NewOrm()
	user := models.User{Username:username, Password:password}

	err := o.Read(&user,"Username", "Password")
	fmt.Println(err)
	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
		this.Redirect("/login", 302)
	} else {
	    fmt.Println(user.Password, user.Username)
	    this.Redirect("/", 302)
	}



	// var user models.User
	// if VerifyUser(&user, username, password) {
	// 	v := this.GetSession(sessionName)
	// 	if v == nil {
	// 		this.SetSession(sessionName, user.Id)
	// 	}
	// 	this.Redirect("/index", 302)

	// } else {
	// 	this.Redirect("/register", 302)
	// }

}

// func VerifyUser(user *models.User, username, password string) (success bool) {
// 	// search user by username or email
// 	if HasUser(user, username) == false {
// 		return
// 	}
// 	if password == user.Password {
// 		// success
// 		success = true
// 	}
// 	return
// }
package service

import (
	"IMchat/models"
	"IMchat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// GetUserList
// @Summary 所有用户信息
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// LoginUserByNameAndPwd
// @Summary 登录用户
// @Tags 用户模块
// @param name query string false "name"
// @param password query string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/loginUserByNameAndPwd [post]
func LoginUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "User does not exist",
			"data":    data,
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "Incorrect password",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.LoginUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "Login successfully",
		"data":    data,
	})
}

// CreateUser
// @Summary 新增用户信息
// @Tags 用户模块
// @param name query string false "name"
// @param password query string false "password"
// @param repassword query string false "repassword"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if user.Name == "" || password == "" || repassword == "" {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "The username or password cannot be empty",
			"data":    user,
		})
		return
	}
	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "Username has been registered",
			"data":    user,
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "Different passwords",
			"data":    user,
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "New user successfully",
		"data":    user,
	})
}

// DeleteUser
// @Summary 删除用户信息
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "Delete user information successfully",
		"data":    user,
	})
}

// UpdateUser
// @Summary 修改用户信息
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		panic(err)
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "Parameter modification error",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0, //  0成功   -1失败
			"message": "Update user information successfully",
			"data":    user,
		})
	}
}

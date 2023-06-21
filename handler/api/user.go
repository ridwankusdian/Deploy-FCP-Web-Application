package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	// Buat variabel baru
	var pengguna model.UserLogin
	// Model Binding
	if salah := c.BindJSON(&pengguna); salah != nil {
		c.JSON(http.StatusBadRequest,
			model.NewErrorResponse("invalid decode json"))
		return
	}

	// Cek jika Email atau Password kosong
	if pengguna.Email == "" || pengguna.Password == "" {
		c.JSON(http.StatusBadRequest,
			model.NewErrorResponse("email or password is empty"))
		return
	}

	// Buat variabel baru atau Model User
	modelUser := &model.User{
		Email:    pengguna.Email,
		Password: pengguna.Password,
	}

	// function Login dari Service
	token,
		salah := u.userService.Login(modelUser)
	if salah != nil {
		c.JSON(http.StatusInternalServerError,
			model.NewErrorResponse("error internal server"))
		return
	}

	// Cookie
	cookie,
		salah := c.Cookie("session_token")
	if salah != nil {
		cookie = *token
		c.SetCookie("session_token", cookie, 3600, "/", "", false, true)
	}

	// Response MEssage
	c.JSON(http.StatusOK,
		model.NewSuccessResponse("login success"))

	// TODO: answer here
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// function GetUserTaskCategory dari Service
	penggunaTaskCategory,
		salah := u.userService.GetUserTaskCategory()

	if salah != nil {
		c.JSON(http.StatusInternalServerError,
			salah)
		return
	}

	// Respon Pesan
	c.JSON(http.StatusOK,
		penggunaTaskCategory)

	// TODO: answer here
}

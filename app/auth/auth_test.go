package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/vmkevv/rigelapi/app"
	"github.com/vmkevv/rigelapi/app/auth"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/enttest"
)

func errPanic(e error) {
	if e != nil {
		panic(e)
	}
}

var dioUser = auth.SignUpReq{
	Name:     "Dio",
	LastName: "Brandom",
	Email:    "dio@dio.com",
	Password: "123456",
}

var danUser = auth.SignUpReq{
	Name:     "Dan",
	LastName: "Vase",
	Email:    "dan@dan.com",
	Password: "654321",
}

type AuthTestSuite struct {
	suite.Suite
	appServer app.Server
	ent       *ent.Client
	teacher   auth.SignInRes
}

func (suite *AuthTestSuite) SetupTest() {
	ctx := context.Background()
	config := config.GetConfig()
	suite.ent = enttest.Open(suite.T(), "sqlite3", "file:ent?mode=memory&_fk=1")
	errPanic(suite.ent.Schema.Create(ctx))
	logger := log.New(os.Stdout, "", log.LstdFlags)
	suite.appServer = app.NewServer(suite.ent, config, logger, ctx)
	auth.Setup(suite.appServer)
}

func (suite *AuthTestSuite) TearDownTest() {
	suite.appServer.DB.Close()
}

func (suite *AuthTestSuite) makeAuthRequest(method, url string, data interface{}) *http.Response {
	reqData, err := json.Marshal(data)
	errPanic(err)
	req := httptest.NewRequest(method, "http://0.0.0.0:4000"+url, bytes.NewBuffer(reqData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", suite.teacher.JWT)
	resp, err := suite.appServer.App.Test(req)
	errPanic(err)
	return resp
}

func (suite *AuthTestSuite) makeRequest(method, url string, data interface{}) *http.Response {
	reqData, err := json.Marshal(data)
	errPanic(err)
	req := httptest.NewRequest(method, "http://0.0.0.0:4000"+url, bytes.NewBuffer(reqData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := suite.appServer.App.Test(req)
	errPanic(err)
	return resp
}

func (suite *AuthTestSuite) signupUser(user auth.SignUpReq) {
	suite.makeRequest("POST", "/signup", user)
}

func (suite *AuthTestSuite) signinUser(user auth.SignUpReq) auth.SignInRes {
	reqData := auth.SignInReq{Email: user.Email, Password: user.Password}
	resp := suite.makeRequest("POST", "/signin", reqData)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var teacherData auth.SignInRes
	suite.Nil(json.Unmarshal(respData, &teacherData))
	return teacherData
}

func (suite *AuthTestSuite) TestSignUp() {
	// new user
	resp := suite.makeRequest("POST", "/signup", dioUser)
	suite.Equal(fiber.StatusCreated, resp.StatusCode)
	// existing user
	reqData := auth.SignUpReq{Email: dioUser.Email}
	resp = suite.makeRequest("POST", "/signup", reqData)
	suite.Equal(fiber.StatusBadRequest, resp.StatusCode)
}

func (suite *AuthTestSuite) TestSignIn() {
	suite.makeRequest("POST", "/signup", dioUser)
	// sing in
	reqData := auth.SignInReq{Email: dioUser.Email, Password: dioUser.Password}
	resp := suite.makeRequest("POST", "/signin", reqData)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var teacherData auth.SignInRes
	suite.Nil(json.Unmarshal(respData, &teacherData))
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	// sing in bad credentials
	reqData = auth.SignInReq{Email: "bad@email.com"}
	resp = suite.makeRequest("POST", "/signin", reqData)
	suite.Equal(fiber.StatusBadRequest, resp.StatusCode)
}

func (suite *AuthTestSuite) TestProfile() {
	suite.signupUser(dioUser)
	suite.teacher = suite.signinUser(dioUser)
	resp := suite.makeAuthRequest("GET", "/auth/profile", nil)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var teacherData models.TeacherWithSubs
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	suite.Nil(json.Unmarshal(respData, &teacherData))
}

func (suite *AuthTestSuite) TestTeacherProfile() {
	suite.signupUser(danUser)
	suite.signupUser(dioUser)
	dan := suite.signinUser(danUser)
	suite.teacher = suite.signinUser(dioUser)
	dio, err := suite.ent.Teacher.Get(suite.appServer.DBCtx, suite.teacher.Teacher.ID)
	suite.Nil(err)
	dio.Update().SetIsAdmin(true).Save(suite.appServer.DBCtx)
	// Admin request
	resp := suite.makeAuthRequest("GET", "/admin/teachers/"+dan.Teacher.ID, nil)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var teacherData models.TeacherWithSubs
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	suite.Nil(json.Unmarshal(respData, &teacherData))
}

func (suite *AuthTestSuite) TestTeacherList() {
	suite.signupUser(danUser)
	suite.signupUser(dioUser)
	suite.signinUser(danUser)
	suite.teacher = suite.signinUser(dioUser)
	dio, err := suite.ent.Teacher.Get(suite.appServer.DBCtx, suite.teacher.Teacher.ID)
	suite.Nil(err)
	dio.Update().SetIsAdmin(true).Save(suite.appServer.DBCtx)
	// Admin request
	resp := suite.makeAuthRequest("GET", "/admin/teachers", nil)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var teachers []models.Teacher
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	suite.Nil(json.Unmarshal(respData, &teachers))
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

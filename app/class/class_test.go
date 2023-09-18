package class_test

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
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/vmkevv/rigelapi/app"
	"github.com/vmkevv/rigelapi/app/auth"
	"github.com/vmkevv/rigelapi/app/class"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/database"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/enttest"
)

func errPanic(e error) {
	if e != nil {
		panic(e)
	}
}

var teacherReq = auth.SignUpReq{
	Name:     "Dio",
	LastName: "Brandom",
	Email:    "dio@dio.com",
	Password: "123456",
}

var classReq = class.SaveClassReq{
	GradeID:   "1",
	SubjectID: "1",
	SchoolID:  "80390001",
	YearID:    "2",
	Parallel:  "Rojo",
}

var classTwoReq = class.SaveClassReq{
	GradeID:   "2",
	SubjectID: "2",
	SchoolID:  "80390001",
	YearID:    "2",
	Parallel:  "Verde",
}

type ClassTestSuite struct {
	suite.Suite
	appServer app.Server
	ent       *ent.Client
	teacher   auth.SignInRes
}

func (suite *ClassTestSuite) SetupTest() {
	ctx := context.Background()
	config := config.GetConfig()
	suite.ent = enttest.Open(suite.T(), "sqlite3", "file:ent?mode=memory&_fk=1")
	errPanic(suite.ent.Schema.Create(ctx))
	errPanic(database.PopulateStaticData(suite.ent, ctx))
	dpto := suite.ent.Dpto.Create().SetID("1").SetName("dpto test").SaveX(ctx)
	prov := suite.ent.Provincia.Create().
		SetID("1").
		SetName("provincia test").
		SetDepartamento(dpto).
		SaveX(ctx)
	mun := suite.ent.Municipio.Create().
		SetID("1").
		SetName("municipio test").
		SetProvincia(prov).
		SaveX(ctx)
	suite.ent.School.Create().
		SetID("80390001").
		SetName("School test").
		SetLat("").
		SetLon("").
		SetMunicipio(mun).
		SaveX(ctx)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	suite.appServer = app.NewServer(suite.ent, config, logger, ctx)
	auth.Setup(suite.appServer)
	class.Setup(suite.appServer)
}

func (suite *ClassTestSuite) TearDownTest() {
	suite.appServer.DB.Close()
}

func (suite *ClassTestSuite) makeAuthRequest(method, url string, data interface{}) *http.Response {
	reqData, err := json.Marshal(data)
	errPanic(err)
	req := httptest.NewRequest(method, "http://0.0.0.0:4000"+url, bytes.NewBuffer(reqData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", suite.teacher.JWT)
	resp, err := suite.appServer.App.Test(req)
	errPanic(err)
	return resp
}

func (suite *ClassTestSuite) makeRequest(method, url string, data interface{}) *http.Response {
	reqData, err := json.Marshal(data)
	errPanic(err)
	req := httptest.NewRequest(method, "http://0.0.0.0:4000"+url, bytes.NewBuffer(reqData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := suite.appServer.App.Test(req)
	errPanic(err)
	return resp
}

func (suite *ClassTestSuite) signinUser() {
	suite.makeRequest("POST", "/signup", teacherReq)
	reqData := auth.SignInReq{Email: teacherReq.Email, Password: teacherReq.Password}
	resp := suite.makeRequest("POST", "/signin", reqData)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var teacherData auth.SignInRes
	suite.Nil(json.Unmarshal(respData, &teacherData))
	suite.teacher = teacherData
}

func (suite *ClassTestSuite) TestCreateClass() {
	suite.signinUser()
	resp := suite.makeAuthRequest("POST", "/auth/class", classReq)
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var classes []models.ClassData
	suite.Nil(json.Unmarshal(respData, &classes))
	suite.Equal(1, len(classes))
}

func (suite *ClassTestSuite) TestTeacherClasses() {
	suite.signinUser()
	suite.makeAuthRequest("POST", "/auth/class", classReq)
	suite.makeAuthRequest("POST", "/auth/class", classTwoReq)
	resp := suite.makeAuthRequest("GET", "/auth/classes/year/2", nil)
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var classes []models.ClassData
	suite.Nil(json.Unmarshal(respData, &classes))
	suite.Equal(2, len(classes))
}

func (suite *ClassTestSuite) TestClassDetails() {
	suite.signinUser()
	postResp := suite.makeAuthRequest("POST", "/auth/class", classReq)
	postRespData, _ := io.ReadAll(postResp.Body)
	var classes []models.ClassData
	suite.Nil(json.Unmarshal(postRespData, &classes))
	suite.Equal(1, len(classes))
	class := classes[0]
	student := suite.ent.Student.Create().
		SetID("1").
		SetName("st name").
		SetLastName("st last Name").
		SetCi("888").
		SetClassID(class.ID).
		SaveX(suite.appServer.DBCtx)
	classPeriod := suite.ent.ClassPeriod.Create().
		SetID("1").
		SetStart(time.Now()).
		SetEnd(time.Now()).
		SetFinished(false).
		SetPeriodID("4").
		SetClassID(class.ID).
		SaveX(suite.appServer.DBCtx)
	activity := suite.ent.Activity.Create().
		SetID("1").
		SetName("act test").
		SetDate(time.Now()).
		SetClassPeriodID(classPeriod.ID).
		SetAreaID("6").
		SaveX(suite.appServer.DBCtx)
	attDay := suite.ent.AttendanceDay.Create().
		SetID("1").
		SetDay(time.Now()).
		SetClassPeriodID(classPeriod.ID).
		SaveX(suite.appServer.DBCtx)
	suite.ent.Attendance.Create().
		SetID("1").
		SetStudentID(student.ID).
		SetAttendanceDayID(attDay.ID).
		SetValue(attendance.ValueAtraso).
		SaveX(suite.appServer.DBCtx)
	suite.ent.Score.Create().
		SetID("1").
		SetStudentID(student.ID).
		SetActivityID(activity.ID).
		SetPoints(100).
		SaveX(suite.appServer.DBCtx)
	resp := suite.makeRequest("GET", "/class/"+class.ID, nil)
	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(ClassTestSuite))
}

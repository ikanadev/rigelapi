package extra_test

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
	"github.com/vmkevv/rigelapi/app/extra"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/database"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/enttest"
)

func errPanic(e error) {
	if e != nil {
		panic(e)
	}
}

type ExtraTestSuite struct {
	suite.Suite
	appServer app.Server
	ent       *ent.Client
}

func (suite *ExtraTestSuite) SetupTest() {
	ctx := context.Background()
	config := config.GetConfig()
	suite.ent = enttest.Open(suite.T(), "sqlite3", "file:ent?mode=memory&_fk=1")
	errPanic(suite.ent.Schema.Create(ctx))
	errPanic(database.PopulateStaticData(suite.ent, ctx))
	logger := log.New(os.Stdout, "", log.LstdFlags)
	suite.appServer = app.NewServer(suite.ent, config, logger, ctx)
	extra.Setup(suite.appServer)
}

func (suite *ExtraTestSuite) TearDownTest() {
	suite.appServer.DB.Close()
}

func (suite *ExtraTestSuite) makeRequest(method, url string, data interface{}) *http.Response {
	reqData, err := json.Marshal(data)
	errPanic(err)
	req := httptest.NewRequest(method, "http://0.0.0.0:4000"+url, bytes.NewBuffer(reqData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := suite.appServer.App.Test(req)
	errPanic(err)
	return resp
}

func (suite *ExtraTestSuite) TestYears() {
	resp := suite.makeRequest("GET", "/years", nil)
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var years []models.YearData
	suite.Nil(json.Unmarshal(respData, &years))
	suite.Equal(2, len(years))
}

func (suite *ExtraTestSuite) TestStatic() {
	resp := suite.makeRequest("GET", "/static", nil)
	suite.Equal(fiber.StatusOK, resp.StatusCode)
	respData, err := io.ReadAll(resp.Body)
	suite.Nil(err)
	var static extra.StaticDataRes
	suite.Nil(json.Unmarshal(respData, &static))
}

func (suite *ExtraTestSuite) TestStats() {
	resp := suite.makeRequest("GET", "/stats", nil)
	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func (suite *ExtraTestSuite) TestSaveErrors() {
	errors := []models.AppError{
		{
			ID:         "1",
			UserID:     "1",
			Cause:      "cause test",
			ErrorMsg:   "error msg",
			ErrorStack: "file.js 2.3",
		},
	}
	resp := suite.makeRequest("POST", "/errors", errors)
	suite.Equal(fiber.StatusNoContent, resp.StatusCode)
}

func TestExtraSuite(t *testing.T) {
	suite.Run(t, new(ExtraTestSuite))
}

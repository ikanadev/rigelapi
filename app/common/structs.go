package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/vmkevv/rigelapi/app/models"
)

/* Data included in the token */

type AppClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

/* request of any sync request */

type DBOperation = string

const (
	Insert DBOperation = "INSERT"
	Update             = "UPDATE"
	Delete             = "DELETE"
)

type BaseTx struct {
	ID       string      `json:"id"`
	DateTime int64       `json:"date_time"`
	Type     DBOperation `json:"type"`
}

type AppStudentTx struct {
	BaseTx
	Data models.AppStudent `json:"data"`
}

type AppClassPeriodTx struct {
	BaseTx
	Data models.AppClassPeriod `json:"data"`
}

type AppAttendanceDayTx struct {
	BaseTx
	Data models.AppAttendanceDay `json:"data"`
}

type AppAttendanceTx struct {
	BaseTx
	Data models.AppAttendance `json:"data"`
}

type AppActivityTx struct {
	BaseTx
	Data models.AppActivity `json:"data"`
}

type AppScoreTx struct {
	BaseTx
	Data models.AppScore `json:"data"`
}

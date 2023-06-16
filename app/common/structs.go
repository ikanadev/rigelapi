package common

import "github.com/golang-jwt/jwt/v4"

/* Data included in the token */

type AppClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

/* response of any sync request */

type SyncIDResp struct {
	LastSyncID string `json:"last_sync_id"`
}

/* request of any sync request */

type DBOperation = string

const (
	Insert DBOperation = "INSERT"
	Update             = "UPDATE"
	Delete             = "DELETE"
)

type SyncReqBase struct {
	ID       string      `json:"id"`
	DateTime int64       `json:"date_time"`
	Type     DBOperation `json:"type"`
}

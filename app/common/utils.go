package common

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
)

/* rollback a transaction */

func RollbackTx(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func GenClaims(ID string) AppClaims {
	return AppClaims{
		ID:               ID,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
}

/* Build teacher profile response */
func BuildTeacherProfile(teacher *ent.Teacher) models.TeacherWithSubs {
	resp := models.TeacherWithSubs{
		Teacher: models.Teacher{
			ID:       teacher.ID,
			Name:     teacher.Name,
			LastName: teacher.LastName,
			Email:    teacher.Email,
			IsAdmin:  teacher.IsAdmin,
		},
		Subscriptions: make([]models.SubWithYear, len(teacher.Edges.Subscriptions)),
	}
	for i, subs := range teacher.Edges.Subscriptions {
		resp.Subscriptions[i] = models.SubWithYear{
			Subscription: models.Subscription{
				ID:     subs.ID,
				Method: subs.Method,
				Qtty:   subs.Qtty,
				Date:   subs.Date.UnixMilli(),
			},
			Year: models.Year{
				ID:    subs.Edges.Year.ID,
				Value: subs.Edges.Year.Value,
			},
		}
	}
	return resp
}

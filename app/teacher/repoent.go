package teacher

import (
	"context"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/subscription"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

type TeacherEntRepo struct {
	ent *ent.Client
	ctx context.Context
}

func NewTeacherEntRepo(ent *ent.Client, ctx context.Context) TeacherEntRepo {
	return TeacherEntRepo{ent, ctx}
}

func (ter TeacherEntRepo) GetProfile(teacherID string) (
	models.TeacherWithSubs,
	error,
) {
	var profile models.TeacherWithSubs
	entTeacher, err := ter.ent.Teacher.
		Query().
		Where(teacher.ID(teacherID)).
		WithSubscriptions(func(sq *ent.SubscriptionQuery) {
			sq.WithYear()
			sq.Order(ent.Asc(subscription.FieldDate))
		}).
		First(ter.ctx)
	if err != nil {
		return profile, err
	}

	profile.Teacher = models.Teacher{
		ID:       entTeacher.ID,
		Name:     entTeacher.Name,
		LastName: entTeacher.LastName,
		Email:    entTeacher.Email,
		IsAdmin:  entTeacher.IsAdmin,
	}
	profile.Subscriptions = make(
		[]models.SubWithYear,
		len(entTeacher.Edges.Subscriptions),
	)
	for i, subs := range entTeacher.Edges.Subscriptions {
		profile.Subscriptions[i] = models.SubWithYear{
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
	return profile, nil
}

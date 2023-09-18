package subscription

import (
	"context"
	"time"

	"github.com/vmkevv/rigelapi/ent"
)

func NewSubscriptionEntRepo(
	ent *ent.Client,
	ctx context.Context,
	idGenerator func() string,
) SubscriptionEntRepo {
	return SubscriptionEntRepo{ent, ctx, idGenerator}
}

type SubscriptionEntRepo struct {
	ent         *ent.Client
	ctx         context.Context
	idGenerator func() string
}

func (ser SubscriptionEntRepo) SaveSubscription(
	teacherID string,
	yearID string,
	method string,
	qtty int,
) error {
	_, err := ser.ent.Subscription.Create().
		SetID(ser.idGenerator()).
		SetMethod(method).
		SetQtty(qtty).
		SetDate(time.Now()).
		SetTeacherID(teacherID).
		SetYearID(yearID).
		Save(ser.ctx)
	return err
}

func (ser SubscriptionEntRepo) UpdateSubscription(subID string, method string, qtty int) error {
	_, err := ser.ent.Subscription.UpdateOneID(subID).SetMethod(method).SetQtty(qtty).Save(ser.ctx)
	return err
}

func (ser SubscriptionEntRepo) DeleteSubscription(subID string) error {
	err := ser.ent.Subscription.DeleteOneID(subID).Exec(ser.ctx)
	return err
}

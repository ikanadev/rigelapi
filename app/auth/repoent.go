package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/subscription"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"golang.org/x/crypto/bcrypt"
)

type AuthEntRepo struct {
	ent    *ent.Client
	ctx    context.Context
	config config.Config
	genID  func() string
}

func NewAuthEntRepo(
	ent *ent.Client,
	ctx context.Context,
	config config.Config,
	genID func() string,
) AuthEntRepo {
	return AuthEntRepo{
		ent, ctx, config, genID,
	}
}

func (aer AuthEntRepo) Register(name, lastName, email, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = aer.ent.Teacher.Query().Where(teacher.EmailEQ(email)).First(aer.ctx)
	// err = nil means teacher exists
	if err == nil {
		return common.NewClientErr(
			fiber.StatusBadRequest,
			"Ups! ya existe una cuenta con ese correo.",
		)
	}
	_, err = aer.ent.Teacher.
		Create().
		SetID(aer.genID()).
		SetName(name).
		SetEmail(email).
		SetLastName(lastName).
		SetPassword(string(hashedPass)).
		Save(aer.ctx)
	return nil
}

func (aer AuthEntRepo) GetTeacher(email, password string) (models.TeacherWithSubs, string, error) {
	teacherRes := models.TeacherWithSubs{}
	tokenStr := ""
	entTeacher, err := aer.ent.Teacher.Query().
		Where(teacher.EmailEQ(email)).
		WithSubscriptions(func(sq *ent.SubscriptionQuery) {
			sq.WithYear()
			sq.Order(ent.Asc(subscription.FieldDate))
		}).
		First(aer.ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return teacherRes,
				tokenStr,
				common.NewClientErr(fiber.StatusBadRequest, "Credenciales incorrectas.")
		}
		return teacherRes, tokenStr, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(entTeacher.Password), []byte(password))
	if err != nil {
		return teacherRes,
			tokenStr,
			common.NewClientErr(fiber.StatusBadRequest, "Credenciales incorrectas.")
	}
	claims := common.GenClaims(entTeacher.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString([]byte(aer.config.App.JWTKey))
	if err != nil {
		return teacherRes, tokenStr, err
	}
	teacherRes = common.BuildTeacherProfile(entTeacher)
	return teacherRes, tokenStr, nil
}

func (aer AuthEntRepo) GetTeacherProfile(teacherID string) (models.TeacherWithSubs, error) {
	var profile models.TeacherWithSubs
	entTeacher, err := aer.ent.Teacher.
		Query().
		Where(teacher.ID(teacherID)).
		WithSubscriptions(func(sq *ent.SubscriptionQuery) {
			sq.WithYear()
			sq.Order(ent.Asc(subscription.FieldDate))
		}).
		First(aer.ctx)
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
	profile.Subscriptions = make([]models.SubWithYear, len(entTeacher.Edges.Subscriptions))
	for i, sub := range entTeacher.Edges.Subscriptions {
		profile.Subscriptions[i] = models.SubWithYear{
			Subscription: models.Subscription{
				ID:     sub.ID,
				Method: sub.Method,
				Qtty:   sub.Qtty,
				Date:   sub.Date.UnixMilli(),
			},
			Year: models.Year{
				ID:    sub.Edges.Year.ID,
				Value: sub.Edges.Year.Value,
			},
		}
	}
	return profile, nil
}

func (aer AuthEntRepo) GetTeachers() ([]models.Teacher, error) {
	entTeachers, err := aer.ent.Teacher.Query().Order(ent.Asc(teacher.FieldLastName)).All(aer.ctx)
	if err != nil {
		return nil, err
	}
	teachers := make([]models.Teacher, len(entTeachers))
	for i, teacher := range entTeachers {
		teachers[i] = models.Teacher{
			ID:       teacher.ID,
			Name:     teacher.Name,
			LastName: teacher.LastName,
			Email:    teacher.Email,
			IsAdmin:  teacher.IsAdmin,
		}
	}
	return teachers, nil
}

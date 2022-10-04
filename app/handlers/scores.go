package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type Score struct {
	ID         string `json:"id"`
	StudentId  string `json:"student_id"`
	ActivityId string `json:"activity_id"`
	Points     int    `json:"points"`
}

func SaveScores(db *ent.Client) func(*fiber.Ctx) error {
	type ScoreReq struct {
		SyncReqBase
		Data Score `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		scores := []ScoreReq{}
		err := c.BodyParser(&scores)
		if err != nil {
			return err
		}
		tx, err := db.Tx(c.Context())
		if err != nil {
			return err
		}

		toAdd := []*ent.ScoreCreate{}
		toUpdate := []Score{}
		for _, score := range scores {
			switch score.Type {
			case Insert:
				{
					toAdd = append(
						toAdd,
						tx.Score.
							Create().
							SetID(score.Data.ID).
							SetPoints(score.Data.Points).
							SetStudentID(score.Data.StudentId).
							SetActivityID(score.Data.ActivityId),
					)
				}
			case Update:
				{
					toUpdate = append(toUpdate, score.Data)
				}
			}
		}

		_, err = tx.Score.CreateBulk(toAdd...).Save(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		for _, score := range toUpdate {
			_, err = tx.Score.UpdateOneID(score.ID).SetPoints(score.Points).Save(c.Context())
			if err != nil {
				return rollback(tx, err)
			}
		}
		err = tx.Commit()
		if err != nil {
			return rollback(tx, err)
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func GetScores(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		yearID := c.Params("yearid")
		serverScores, err := db.Score.
			Query().
			Where(
				score.HasActivityWith(
					activity.HasClassPeriodWith(
						classperiod.HasClassWith(
							class.HasTeacherWith(teacher.IDEQ(teacherID)),
							class.HasYearWith(year.IDEQ(yearID)),
						),
					),
				),
			).
			WithActivity().
			WithStudent().
			All(c.Context())
		if err != nil {
			return err
		}
		scores := make([]Score, len(serverScores))
		for index, score := range serverScores {
			scores[index] = Score{
				ID:         score.ID,
				Points:     score.Points,
				StudentId:  score.Edges.Student.ID,
				ActivityId: score.Edges.Activity.ID,
			}
		}
		return c.JSON(scores)
	}
}

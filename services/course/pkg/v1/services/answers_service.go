package services

import (
	"context"
	"fmt"

	repository "gitlab.com/scalent/goxpert/services/course/pkg/v1/repositories"
)

// AnswersService describes the service.
type AnswersService interface {
	SubmitAnswer(context.Context, string, string, string) error
}

//AnswersServiceImpl **
type AnswersServiceImpl struct {
	answersRepo repository.AnswersRepository
	userRepo    repository.UserServiceRepo
}

//NewAnswersServiceImpl inject depedancies user repositiory
func NewAnswersServiceImpl(answersRepo repository.AnswersRepository, userRepo repository.UserServiceRepo) AnswersService {
	return &AnswersServiceImpl{answersRepo: answersRepo, userRepo: userRepo}
}

func (answersService *AnswersServiceImpl) SubmitAnswer(ctx context.Context, question_id string, answer string, uuid string) error {

	user, err := answersService.userRepo.GetUsersInfoByUUID(ctx, uuid)

	if err != nil {
		return err
	}

	fmt.Println(user)

	question, err := answersService.answersRepo.GetQuetionByID(ctx, question_id)

	if err != nil {
		return err
	}

	if question.QuestionType == "mcq" {
		option, err := answersService.answersRepo.GetCorrectOptionsByID(ctx, question_id)
		var iscorrect string

		if err != nil {
			//iscorrect = "false"
			return err
		}

		if option.Name == answer {
			iscorrect = "true"
			return nil
		}

		if err != nil {
			iscorrect = "false"
			return err
		}

		err = answersService.answersRepo.SubmitAnswer(ctx, question_id, answer, iscorrect, user.ID)

	} else {

		//testcases := question.TestCases

		//call service for run code using testcase and answer

	}

	return nil
}

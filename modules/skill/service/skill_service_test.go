package service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"skilledin-green-skills-api/mocks"
	"skilledin-green-skills-api/modules/skill/common"
	"skilledin-green-skills-api/modules/skill/domain/entity"
	"skilledin-green-skills-api/modules/skill/service"
	"testing"
)

type skillServiceTestSuite struct {
	suite.Suite
	skills      []*entity.GreenSkill
	skillGroups []*entity.GreenSkillGroup
}

func (s *skillServiceTestSuite) SetupSuite() {
	s.skills = []*entity.GreenSkill{
		{
			ID:   "123",
			Name: "lorem",
		},
	}

	s.skillGroups = []*entity.GreenSkillGroup{
		{
			ID:   "123",
			Name: "lorem",
		},
	}
}

func (s *skillServiceTestSuite) Test_GreenSkillList_ShouldSuccess() {
	repo := new(mocks.ISkillRepository)
	svc := service.NewSkillService(repo)
	repo.On("AggregateGreenSkill", mock.Anything, mock.Anything).
		Return([]*entity.GreenSkill{s.skills[0]}, nil).Once()
	repo.On("CountGreenSkill", mock.Anything, mock.Anything).
		Return(int64(1)).Once()
	data, err := svc.GreenSkillList(context.TODO(), 10, 20, "lorem")
	s.Nil(err)
	s.NotNil(data)
	repo.AssertExpectations(s.T())
}
func (s *skillServiceTestSuite) Test_GreenSkillList_ShouldError() {
	repo := new(mocks.ISkillRepository)
	svc := service.NewSkillService(repo)
	repo.On("AggregateGreenSkill", mock.Anything, mock.Anything).
		Return(nil, errors.New("LOREM")).Once()
	data, err := svc.GreenSkillList(context.TODO(), 10, 20, "lorem")
	s.NotNil(err)
	s.Nil(data)
	repo.AssertExpectations(s.T())
}

func (s *skillServiceTestSuite) Test_GreenSkillDetail_ShouldSuccess() {
	repo := new(mocks.ISkillRepository)
	svc := service.NewSkillService(repo)
	repo.On("AggregateGreenSkill", mock.Anything, mock.Anything).
		Return([]*entity.GreenSkill{s.skills[0]}, nil).Once()
	data, err := svc.GreenSkillDetail(context.TODO(), "lorem")
	s.Nil(err)
	s.NotNil(data)
	repo.AssertExpectations(s.T())
}
func (s *skillServiceTestSuite) Test_GreenSkillDetail_ShouldError() {
	repo := new(mocks.ISkillRepository)
	svc := service.NewSkillService(repo)
	s.T().Run("ERROR WHEN AGGREGATE DATA", func(t *testing.T) {
		repo.On("AggregateGreenSkill", mock.Anything, mock.Anything).
			Return(nil, errors.New("LOREM")).Once()
		data, err := svc.GreenSkillDetail(context.TODO(), "lorem")
		s.NotNil(err)
		s.Nil(data)
		repo.AssertExpectations(s.T())
	})
	s.T().Run("ERROR WHEN AGGREGATE DATA", func(t *testing.T) {
		repo.On("AggregateGreenSkill", mock.Anything, mock.Anything).
			Return([]*entity.GreenSkill{}, nil).Once()
		data, err := svc.GreenSkillDetail(context.TODO(), "lorem")
		s.NotNil(err)
		s.Nil(data)
		s.Equal(err, common.ErrNotDataFound)
		repo.AssertExpectations(s.T())
	})
}

func TestSkillService(t *testing.T) {
	suite.Run(t, new(skillServiceTestSuite))
}

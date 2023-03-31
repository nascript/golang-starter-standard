package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"skilledin-green-skills-api/modules/skill/common"
	"skilledin-green-skills-api/modules/skill/domain"
	"skilledin-green-skills-api/modules/skill/domain/response"
	mongoRepo "skilledin-green-skills-api/modules/skill/repository/mongo"
	"skilledin-green-skills-api/pkg/http/wrapper"
)

type skillService struct {
	repository domain.ISkillRepository
}

func (service *skillService) GreenSkillList(
	ctx context.Context,
	limit, offset int,
	search string,
) (item *wrapper.PaginationResponse[response.GreenSkillResponse], err error) {
	pipeline := mongoRepo.GeneralSkillAggregatePipeline(limit, offset, search)
	data, err := service.repository.AggregateGreenSkill(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var res []*response.GreenSkillResponse
	var count int64
	if len(data) > 0 {
		for _, skill := range data {
			res = append(res, &response.GreenSkillResponse{
				ID:          skill.ID,
				Name:        skill.Name,
				Description: skill.Description,
			})
		}
		countPipeline := mongo.Pipeline{}
		if search != "" {
			countPipeline = append(countPipeline, bson.D{{
				Key: "$match", Value: bson.D{
					{Key: "name", Value: bson.M{
						"$regex": search, "$options": "i",
					}}},
			}})
		}
		count = service.repository.CountGreenSkill(ctx, countPipeline)
	}

	return &wrapper.PaginationResponse[response.GreenSkillResponse]{
		List:  res,
		Total: count,
	}, nil
}

func (service *skillService) GreenSkillDetail(
	ctx context.Context,
	id string,
) (item *response.GreenSkillResponse, err error) {
	data, err := service.repository.AggregateGreenSkill(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}},
		{{Key: "$limit", Value: 1}},
	})
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, common.ErrNotDataFound
	}

	skill := data[0]
	return &response.GreenSkillResponse{
		ID:          skill.ID,
		Name:        skill.Name,
		Description: skill.Description,
	}, nil
}

func NewSkillService(
	repository domain.ISkillRepository,
) domain.ISkillService {
	return &skillService{
		repository,
	}
}

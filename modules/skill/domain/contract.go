package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"skilledin-green-skills-api/modules/skill/domain/entity"
	"skilledin-green-skills-api/modules/skill/domain/response"
	"skilledin-green-skills-api/pkg/http/wrapper"
)

type (
	ISkillRepository interface {
		/*
		 * Part of green skills
		 * get count data and then aggregate the skills lists
		 */
		CountGreenSkill(
			ctx context.Context,
			pipeline mongo.Pipeline,
		) int64
		AggregateGreenSkill(
			ctx context.Context,
			pipeline mongo.Pipeline,
		) (items []*entity.GreenSkill, err error)

		/*
		 * Part of skills collection
		 * has 2 function first is Aggregate
		 *		this function can be use at green cv builder
		 *		(education, employment history, certification, etc)
		 * Find Or Insert
		 *		when ever the user search and data
		 *		did not found then call this function
		 */
		AggregateSkill(
			ctx context.Context,
			pipeline mongo.Pipeline,
		) (items []*entity.Skill, err error)
		FindOrInsertSkills(
			ctx context.Context,
			skills []string,
		) error
	}

	ISkillService interface {
		GreenSkillList(
			ctx context.Context,
			limit, offset int,
			search string,
		) (
			item *wrapper.PaginationResponse[response.GreenSkillResponse],
			err error,
		)

		GreenSkillDetail(
			ctx context.Context,
			id string,
		) (
			item *response.GreenSkillResponse,
			err error,
		)
	}
)

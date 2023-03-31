package mongo

import (
	"context"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"skilledin-green-skills-api/modules/skill/domain"
	"skilledin-green-skills-api/modules/skill/domain/entity"
	"strings"
	"sync"
)

type skillMongoRepository struct {
	skillCollection, greenSkillCollection,
	greenSkillGroupCollection *mongo.Collection
}

func (repo *skillMongoRepository) CountGreenSkill(
	ctx context.Context,
	pipeline mongo.Pipeline,
) int64 {
	pipeline = append(pipeline, bson.D{{Key: "$count", Value: "count"}})
	cursor, err := repo.greenSkillCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return int64(0)
	}
	defer func() { _ = cursor.Close(ctx) }()
	var result struct {
		Count int64 `bson:"count"`
	}
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return int64(0)
		}
	}
	if err := cursor.Err(); err != nil {
		return int64(0)
	}
	return result.Count
}

func (repo *skillMongoRepository) AggregateGreenSkill(
	ctx context.Context,
	pipeline mongo.Pipeline,
) (items []*entity.GreenSkill, err error) {
	cursor, err := repo.greenSkillCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *skillMongoRepository) AggregateSkill(
	ctx context.Context,
	pipeline mongo.Pipeline,
) (items []*entity.Skill, err error) {
	cursor, err := repo.skillCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *skillMongoRepository) FindOrInsertSkills(
	ctx context.Context,
	skills []string,
) error {
	var wg sync.WaitGroup
	for _, skill := range skills {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			_ = repo.skillCollection.FindOneAndUpdate(
				ctx, bson.M{"name": name},
				bson.M{"$set": bson.M{"_id": fiberUtils.UUIDv4(), "name": name}},
				options.FindOneAndUpdate().SetUpsert(true),
			).Err()
		}(strings.ToLower(skill))
	}
	wg.Wait()

	return nil
}

func NewSkillMongoRepository(
	opts ...SkillMongoRepositoryOption,
) domain.ISkillRepository {
	repo := &skillMongoRepository{}

	for _, opt := range opts {
		opt(repo)
	}

	return repo
}

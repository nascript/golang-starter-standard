package mongo_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	mongoRepo "skilledin-green-skills-api/modules/skill/repository/mongo"
	"testing"
)

func Test_CountGreenSkills_ShouldSuccess(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("test count should return data from aggregate", func(mt *mtest.T) {
		mt.AddMockResponses(
			mtest.CreateCursorResponse(1,
				"count.data",
				mtest.FirstBatch,
				bson.D{
					{Key: "count", Value: 12},
				}),
		)
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll))
		data := repo.CountGreenSkill(context.TODO(), nil)
		assert.NotNil(mt, data)
		assert.NotZero(mt, data)
	})
}
func Test_CountGreenSkill_ShouldError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("test count doc should error", func(mt *mtest.T) {
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll))
		count := repo.CountGreenSkill(context.TODO(), nil)
		assert.Zero(mt, count)
	})

	mt.Run("test count should error from aggregate", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    123,
			Message: "READ_FAILED",
		}))
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll))
		count := repo.CountGreenSkill(context.TODO(), nil)
		assert.Zero(mt, count)
	})

	mt.Run("test count should error from decode", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			"count.data",
			mtest.FirstBatch,
			bson.D{
				{Key: "count", Value: "12"},
			}))
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll))
		count := repo.CountGreenSkill(context.TODO(), nil)
		assert.Zero(mt, count)
	})

	mt.Run("test count should error from cursor", func(mt *mtest.T) {
		mt.AddMockResponses(
			mtest.CreateCursorResponse(50, "foo.bar", mtest.FirstBatch),
			mtest.CreateSuccessResponse(),
		)
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll))
		count := repo.CountGreenSkill(context.TODO(), nil)
		assert.Zero(mt, count)
	})
}

func Test_AggregateGreenSkill_ShouldSuccess(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	id1 := primitive.NewObjectID().String()
	id2 := primitive.NewObjectID().String()
	mt.Run("test find should return data", func(mt *mtest.T) {
		first := mtest.CreateCursorResponse(1,
			"greenskill.data",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: id1},
			})
		second := mtest.CreateCursorResponse(1,
			"greenskill.data",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: id2},
			})
		killCursors := mtest.CreateCursorResponse(
			0, "greenskill.data", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll))
		data, err := repo.AggregateGreenSkill(context.TODO(),
			mongoRepo.GeneralSkillAggregatePipeline(1, 2, "lorem"))
		assert.Nil(mt, err)
		assert.NotNil(mt, data)
		assert.Equal(mt, data[0].ID, id1)
		assert.Equal(mt, data[1].ID, id2)
	})
}
func Test_AggregateGreenSkill_ShouldError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	id1 := primitive.NewObjectID().String()
	mt.Run("test all should return data and error on decode", func(mt *mtest.T) {
		first := mtest.CreateCursorResponse(1,
			"greenskill.data",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: id1},
			})
		second := mtest.CreateCursorResponse(1,
			"greenskill.data",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: 123},
			})
		mt.AddMockResponses(first, second)
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll),
			mongoRepo.WithGreenSkillGroupCollection(mt.Coll))
		data, err := repo.AggregateGreenSkill(context.TODO(), nil)
		assert.NotNil(mt, err)
		assert.Nil(mt, data)
		assert.Equal(mt, err.Error(), "error decoding key _id: cannot decode 32-bit integer into a string type")
	})

	mt.Run("test all should error aggregate", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    123,
			Message: "READ_FAILED",
		}))
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithGreenSkillCollection(mt.Coll))
		data, err := repo.AggregateGreenSkill(context.TODO(), nil)
		assert.Nil(mt, data)
		assert.NotNil(mt, err)
		assert.Equal(mt, err.Error(), "write command error: [{write errors: [{READ_FAILED}]}, {<nil>}]")
	})
}

func Test_AggregateSkill_ShouldSuccess(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	id1 := primitive.NewObjectID().String()
	id2 := primitive.NewObjectID().String()
	mt.Run("test find should return data", func(mt *mtest.T) {
		first := mtest.CreateCursorResponse(1,
			"skill.data",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: id1},
			})
		second := mtest.CreateCursorResponse(1,
			"skill.data",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: id2},
			})
		killCursors := mtest.CreateCursorResponse(
			0, "skill.data", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithSkillCollection(mt.Coll))
		data, err := repo.AggregateSkill(context.TODO(), nil)
		assert.Nil(mt, err)
		assert.NotNil(mt, data)
		assert.Equal(mt, data[0].ID, id1)
		assert.Equal(mt, data[1].ID, id2)
	})
}
func Test_AggregateSkill_ShouldError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	id1 := primitive.NewObjectID().String()
	mt.Run("test all should return data and error on decode", func(mt *mtest.T) {
		first := mtest.CreateCursorResponse(1,
			"skill.data",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: id1},
			})
		second := mtest.CreateCursorResponse(1,
			"skill.data",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: 123},
			})
		mt.AddMockResponses(first, second)
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithSkillCollection(mt.Coll))
		data, err := repo.AggregateSkill(context.TODO(), nil)
		assert.NotNil(mt, err)
		assert.Nil(mt, data)
		assert.Equal(mt, err.Error(), "error decoding key _id: cannot decode 32-bit integer into a string type")
	})

	mt.Run("test all should error aggregate", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    123,
			Message: "READ_FAILED",
		}))
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithSkillCollection(mt.Coll))
		data, err := repo.AggregateSkill(context.TODO(), nil)
		assert.Nil(mt, data)
		assert.NotNil(mt, err)
		assert.Equal(mt, err.Error(), "write command error: [{write errors: [{READ_FAILED}]}, {<nil>}]")
	})
}

func Test_FindOrInsertSkills_ShouldSuccess(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("test find should return data", func(mt *mtest.T) {
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				"skill.data",
				mtest.FirstBatch,
				bson.D{
					{Key: "_id", Value: "123"},
				},
			),
		)
		repo := mongoRepo.NewSkillMongoRepository(
			mongoRepo.WithSkillCollection(mt.Coll))
		err := repo.FindOrInsertSkills(context.TODO(), []string{"lorem"})
		assert.Nil(mt, err)
	})
}

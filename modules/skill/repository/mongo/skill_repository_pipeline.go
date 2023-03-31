package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GeneralSkillAggregatePipeline(
	limit, offset int,
	search string,
) mongo.Pipeline {
	pipeline := mongo.Pipeline{}
	if search != "" {
		pipeline = append(pipeline, bson.D{{
			Key: "$match", Value: bson.D{
				{Key: "name", Value: bson.M{
					"$regex": search, "$options": "i",
				}}},
		}})
	}
	pipeline = append(pipeline, mongo.Pipeline{
		bson.D{{Key: "$skip", Value: offset}},
		bson.D{{Key: "$limit", Value: limit}},
	}...)
	return pipeline
}

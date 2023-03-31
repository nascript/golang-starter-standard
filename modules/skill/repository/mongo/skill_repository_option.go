package mongo

import "go.mongodb.org/mongo-driver/mongo"

type SkillMongoRepositoryOption func(*skillMongoRepository)

func WithSkillCollection(
	collection *mongo.Collection,
) SkillMongoRepositoryOption {
	return func(repository *skillMongoRepository) {
		repository.skillCollection = collection
	}
}

func WithGreenSkillCollection(
	collection *mongo.Collection,
) SkillMongoRepositoryOption {
	return func(repository *skillMongoRepository) {
		repository.greenSkillCollection = collection
	}
}

func WithGreenSkillGroupCollection(
	collection *mongo.Collection,
) SkillMongoRepositoryOption {
	return func(repository *skillMongoRepository) {
		repository.greenSkillGroupCollection = collection
	}
}

package entity

type (
	GreenSkill struct {
		ID          string             `bson:"_id"`
		Name        string             `bson:"name"`
		Description string             `bson:"description"`
		Groups      *[]GreenSkillGroup `bson:"groups"`
	}

	GreenSkillGroup struct {
		ID   string `bson:"_id"`
		Name string `bson:"name"`
	}

	Skill struct {
		ID   string `bson:"_id"`
		Name string `bson:"name"`
	}
)

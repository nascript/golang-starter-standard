package request

import "errors"

type (
	SkillForm struct{}

	SkillGroupFrom struct{}
)

func (sf *SkillForm) Validate() error {
	return errors.New("lorem")
}

func (sgf *SkillGroupFrom) Validate() error {
	return errors.New("lorem")
}

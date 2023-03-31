// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "skilledin-green-skills-api/modules/skill/domain/entity"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

// ISkillRepository is an autogenerated mock type for the ISkillRepository type
type ISkillRepository struct {
	mock.Mock
}

// AggregateGreenSkill provides a mock function with given fields: ctx, pipeline
func (_m *ISkillRepository) AggregateGreenSkill(ctx context.Context, pipeline mongo.Pipeline) ([]*entity.GreenSkill, error) {
	ret := _m.Called(ctx, pipeline)

	var r0 []*entity.GreenSkill
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, mongo.Pipeline) ([]*entity.GreenSkill, error)); ok {
		return rf(ctx, pipeline)
	}
	if rf, ok := ret.Get(0).(func(context.Context, mongo.Pipeline) []*entity.GreenSkill); ok {
		r0 = rf(ctx, pipeline)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.GreenSkill)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, mongo.Pipeline) error); ok {
		r1 = rf(ctx, pipeline)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AggregateSkill provides a mock function with given fields: ctx, pipeline
func (_m *ISkillRepository) AggregateSkill(ctx context.Context, pipeline mongo.Pipeline) ([]*entity.Skill, error) {
	ret := _m.Called(ctx, pipeline)

	var r0 []*entity.Skill
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, mongo.Pipeline) ([]*entity.Skill, error)); ok {
		return rf(ctx, pipeline)
	}
	if rf, ok := ret.Get(0).(func(context.Context, mongo.Pipeline) []*entity.Skill); ok {
		r0 = rf(ctx, pipeline)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Skill)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, mongo.Pipeline) error); ok {
		r1 = rf(ctx, pipeline)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountGreenSkill provides a mock function with given fields: ctx, pipeline
func (_m *ISkillRepository) CountGreenSkill(ctx context.Context, pipeline mongo.Pipeline) int64 {
	ret := _m.Called(ctx, pipeline)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, mongo.Pipeline) int64); ok {
		r0 = rf(ctx, pipeline)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// FindOrInsertSkills provides a mock function with given fields: ctx, skills
func (_m *ISkillRepository) FindOrInsertSkills(ctx context.Context, skills []string) error {
	ret := _m.Called(ctx, skills)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, skills)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewISkillRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewISkillRepository creates a new instance of ISkillRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISkillRepository(t mockConstructorTestingTNewISkillRepository) *ISkillRepository {
	mock := &ISkillRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
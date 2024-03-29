// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	response "skilledin-green-skills-api/modules/skill/domain/response"

	wrapper "skilledin-green-skills-api/pkg/http/wrapper"
)

// ISkillService is an autogenerated mock type for the ISkillService type
type ISkillService struct {
	mock.Mock
}

// GreenSkillDetail provides a mock function with given fields: ctx, id
func (_m *ISkillService) GreenSkillDetail(ctx context.Context, id string) (*response.GreenSkillResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 *response.GreenSkillResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*response.GreenSkillResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *response.GreenSkillResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.GreenSkillResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GreenSkillList provides a mock function with given fields: ctx, limit, offset, search
func (_m *ISkillService) GreenSkillList(ctx context.Context, limit int, offset int, search string) (*wrapper.PaginationResponse[response.GreenSkillResponse], error) {
	ret := _m.Called(ctx, limit, offset, search)

	var r0 *wrapper.PaginationResponse[response.GreenSkillResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) (*wrapper.PaginationResponse[response.GreenSkillResponse], error)); ok {
		return rf(ctx, limit, offset, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) *wrapper.PaginationResponse[response.GreenSkillResponse]); ok {
		r0 = rf(ctx, limit, offset, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrapper.PaginationResponse[response.GreenSkillResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, string) error); ok {
		r1 = rf(ctx, limit, offset, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewISkillService interface {
	mock.TestingT
	Cleanup(func())
}

// NewISkillService creates a new instance of ISkillService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISkillService(t mockConstructorTestingTNewISkillService) *ISkillService {
	mock := &ISkillService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	ids "github.com/Toinounet21/avalanchego-trafficked-v1.7.4/ids"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Calculator is an autogenerated mock type for the Calculator type
type Calculator struct {
	mock.Mock
}

// CalculateUptime provides a mock function with given fields: nodeID
func (_m *Calculator) CalculateUptime(nodeID ids.ShortID) (time.Duration, time.Time, error) {
	ret := _m.Called(nodeID)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func(ids.ShortID) time.Duration); ok {
		r0 = rf(nodeID)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	var r1 time.Time
	if rf, ok := ret.Get(1).(func(ids.ShortID) time.Time); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(ids.ShortID) error); ok {
		r2 = rf(nodeID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CalculateUptimePercent provides a mock function with given fields: nodeID
func (_m *Calculator) CalculateUptimePercent(nodeID ids.ShortID) (float64, error) {
	ret := _m.Called(nodeID)

	var r0 float64
	if rf, ok := ret.Get(0).(func(ids.ShortID) float64); ok {
		r0 = rf(nodeID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ids.ShortID) error); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateUptimePercentFrom provides a mock function with given fields: nodeID, startTime
func (_m *Calculator) CalculateUptimePercentFrom(nodeID ids.ShortID, startTime time.Time) (float64, error) {
	ret := _m.Called(nodeID, startTime)

	var r0 float64
	if rf, ok := ret.Get(0).(func(ids.ShortID, time.Time) float64); ok {
		r0 = rf(nodeID, startTime)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ids.ShortID, time.Time) error); ok {
		r1 = rf(nodeID, startTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

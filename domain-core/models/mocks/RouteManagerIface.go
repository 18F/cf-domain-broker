package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/18F/cf-cdn-service-broker/models"

// RouteManagerIface is an autogenerated mock type for the RouteManagerIface type
type RouteManagerIface struct {
	mock.Mock
}

// Create provides a mock function with given fields: instanceId, domain, origin, path, insecureOrigin, forwardedHeaders, forwardCookies, tags
func (_m *RouteManagerIface) Create(instanceId string, domain string, origin string, path string, insecureOrigin bool, forwardedHeaders []string, forwardCookies bool, tags map[string]string) (*models.Route, error) {
	ret := _m.Called(instanceId, domain, origin, path, insecureOrigin, forwardedHeaders, forwardCookies, tags)

	var r0 *models.Route
	if rf, ok := ret.Get(0).(func(string, string, string, string, bool, []string, bool, map[string]string) *models.Route); ok {
		r0 = rf(instanceId, domain, origin, path, insecureOrigin, forwardedHeaders, forwardCookies, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, bool, []string, bool, map[string]string) error); ok {
		r1 = rf(instanceId, domain, origin, path, insecureOrigin, forwardedHeaders, forwardCookies, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOrphanedCerts provides a mock function with given fields:
func (_m *RouteManagerIface) DeleteOrphanedCerts() {
	_m.Called()
}

// Disable provides a mock function with given fields: route
func (_m *RouteManagerIface) Disable(route *models.Route) error {
	ret := _m.Called(route)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Route) error); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: instanceId
func (_m *RouteManagerIface) Get(instanceId string) (*models.Route, error) {
	ret := _m.Called(instanceId)

	var r0 *models.Route
	if rf, ok := ret.Get(0).(func(string) *models.Route); ok {
		r0 = rf(instanceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(instanceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDNSInstructions provides a mock function with given fields: _a0
func (_m *RouteManagerIface) GetDNSInstructions(_a0 []byte) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func([]byte) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Poll provides a mock function with given fields: route
func (_m *RouteManagerIface) Poll(route *models.Route) error {
	ret := _m.Called(route)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Route) error); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Renew provides a mock function with given fields: route
func (_m *RouteManagerIface) Renew(route *models.Route) error {
	ret := _m.Called(route)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Route) error); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenewAll provides a mock function with given fields:
func (_m *RouteManagerIface) RenewAll() {
	_m.Called()
}

// Update provides a mock function with given fields: instanceId, domain, origin, path, insecureOrigin, forwardedHeaders, forwardCookies
func (_m *RouteManagerIface) Update(instanceId string, domain string, origin string, path string, insecureOrigin bool, forwardedHeaders []string, forwardCookies bool) error {
	ret := _m.Called(instanceId, domain, origin, path, insecureOrigin, forwardedHeaders, forwardCookies)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, bool, []string, bool) error); ok {
		r0 = rf(instanceId, domain, origin, path, insecureOrigin, forwardedHeaders, forwardCookies)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ models.RouteManagerIface = (*RouteManagerIface)(nil)

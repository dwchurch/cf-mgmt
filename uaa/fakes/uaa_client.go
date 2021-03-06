// This file was generated by counterfeiter
package fakes

import (
	"sync"

	go_uaa "github.com/cloudfoundry-community/go-uaa"
)

type FakeUaa struct {
	CreateUserStub        func(user go_uaa.User) (*go_uaa.User, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		user go_uaa.User
	}
	createUserReturns struct {
		result1 *go_uaa.User
		result2 error
	}
	ListAllUsersStub        func(filter string, sortBy string, attributes string, sortOrder go_uaa.SortOrder) ([]go_uaa.User, error)
	listAllUsersMutex       sync.RWMutex
	listAllUsersArgsForCall []struct {
		filter     string
		sortBy     string
		attributes string
		sortOrder  go_uaa.SortOrder
	}
	listAllUsersReturns struct {
		result1 []go_uaa.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUaa) CreateUser(user go_uaa.User) (*go_uaa.User, error) {
	fake.createUserMutex.Lock()
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		user go_uaa.User
	}{user})
	fake.recordInvocation("CreateUser", []interface{}{user})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(user)
	} else {
		return fake.createUserReturns.result1, fake.createUserReturns.result2
	}
}

func (fake *FakeUaa) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeUaa) CreateUserArgsForCall(i int) go_uaa.User {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return fake.createUserArgsForCall[i].user
}

func (fake *FakeUaa) CreateUserReturns(result1 *go_uaa.User, result2 error) {
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 *go_uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUaa) ListAllUsers(filter string, sortBy string, attributes string, sortOrder go_uaa.SortOrder) ([]go_uaa.User, error) {
	fake.listAllUsersMutex.Lock()
	fake.listAllUsersArgsForCall = append(fake.listAllUsersArgsForCall, struct {
		filter     string
		sortBy     string
		attributes string
		sortOrder  go_uaa.SortOrder
	}{filter, sortBy, attributes, sortOrder})
	fake.recordInvocation("ListAllUsers", []interface{}{filter, sortBy, attributes, sortOrder})
	fake.listAllUsersMutex.Unlock()
	if fake.ListAllUsersStub != nil {
		return fake.ListAllUsersStub(filter, sortBy, attributes, sortOrder)
	} else {
		return fake.listAllUsersReturns.result1, fake.listAllUsersReturns.result2
	}
}

func (fake *FakeUaa) ListAllUsersCallCount() int {
	fake.listAllUsersMutex.RLock()
	defer fake.listAllUsersMutex.RUnlock()
	return len(fake.listAllUsersArgsForCall)
}

func (fake *FakeUaa) ListAllUsersArgsForCall(i int) (string, string, string, go_uaa.SortOrder) {
	fake.listAllUsersMutex.RLock()
	defer fake.listAllUsersMutex.RUnlock()
	return fake.listAllUsersArgsForCall[i].filter, fake.listAllUsersArgsForCall[i].sortBy, fake.listAllUsersArgsForCall[i].attributes, fake.listAllUsersArgsForCall[i].sortOrder
}

func (fake *FakeUaa) ListAllUsersReturns(result1 []go_uaa.User, result2 error) {
	fake.ListAllUsersStub = nil
	fake.listAllUsersReturns = struct {
		result1 []go_uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUaa) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.listAllUsersMutex.RLock()
	defer fake.listAllUsersMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeUaa) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

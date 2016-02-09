package main

import (
	"errors"
	"log"
)

type InMemStorage struct {
	apiUsers map[string]APIUser
}

func (this InMemStorage) Create(data interface{}) (interface{}, error) {
	if user, ok := data.(APIUser); ok {
		apiMutex.Lock()
		defer apiMutex.Unlock()
		this.apiUsers[user.Token] = user
		return user, nil
	} else {
		log.Print("Can't create new record: uncoprehensible input")
		return nil, errors.New("Can't create new record: uncoprehensible input")
	}
}

func (this InMemStorage) Read(data interface{}) (interface{}, error) {
	if user, ok := data.(APIUser); ok {
		apiMutex.Lock()
		defer apiMutex.Unlock()
		if retUser, userOk := this.apiUsers[user.Token]; userOk {
			return retUser, nil
		} else {
			return nil, nil
		}
	} else {
		log.Print("Can't find user: uncoprehensible input")
		return nil, errors.New("Can't find user: uncoprehensible input")
	}
}

func (this InMemStorage) Update(data interface{}) (interface{}, error) {
	if user, ok := data.(APIUser); ok {
		apiMutex.Lock()
		defer apiMutex.Unlock()
		retUser, userOk := this.apiUsers[user.Token]
		this.apiUsers[user.Token] = user
		if userOk {
			return retUser, nil
		} else {
			return nil, nil
		}
	} else {
		log.Print("Can't update user: uncoprehensible input")
		return nil, errors.New("Can't update user: uncoprehensible input")
	}
}

func (this InMemStorage) Delete(data interface{}) (interface{}, error) {
	if user, ok := data.(APIUser); ok {
		apiMutex.Lock()
		defer apiMutex.Unlock()
		if retUser, userOk := this.apiUsers[user.Token]; userOk {
			apiMutex.Lock()
			defer apiMutex.Unlock()
			delete(this.apiUsers, user.Token)
			return retUser, nil
		} else {
			return nil, nil
		}
	} else {
		log.Print("Can't delete user: uncoprehensible input")
		return nil, errors.New("Cant't delete user: uncoprehensible input")
	}
}

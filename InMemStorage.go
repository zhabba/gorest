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
			log.Print("User not found in DB")
			return nil, errors.New("User not found in DB")
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
		if _, userOk := this.apiUsers[user.Token]; userOk {
			return nil, errors.New("Updating of nonexistent user")
		} else {
			apiMutex.Lock()
			defer apiMutex.Unlock()
			this.apiUsers[user.Token] = user
			return user, nil
		}
	} else {
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
			return nil, errors.New("Deleting of nonexistent user")
		}
	} else {
		return nil, errors.New("Cant't delete user: uncoprehensible input")
	}
}

package main

//type DALStruct struct {
//	IDAL
//}

type DAL struct {
	stor IStorage
}

func New(stor IStorage) IDAL {
	return &DAL{stor}
}

func (this DAL) FindUser(user APIUser) (APIUser, error) {
	if apiUser, err := this.stor.Read(user); err == nil {
		return (apiUser).(APIUser), err
	} else {
		return APIUser{}, err
	}
}

func (this DAL) CreateUser(user APIUser) (APIUser, error) {
	if apiUser, err := this.stor.Create(user); err == nil {
		return (apiUser).(APIUser), err
	} else {
		return APIUser{}, err
	}
}

func (this DAL) UpdateUser(user APIUser) (APIUser, error) {
	if apiUser, err := this.stor.Update(user); err == nil {
		return (apiUser).(APIUser), err
	} else {
		return APIUser{}, err
	}
}

func (this DAL) DeleteUser(user APIUser) (APIUser, error) {
	if apiUser, err := this.stor.Delete(user); err == nil {
		return (apiUser).(APIUser), err
	} else {
		return APIUser{}, err
	}
}

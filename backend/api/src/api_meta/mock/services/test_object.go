package services

import (
	"api_meta/models"
	"api_meta/types"
)

type TestObjectRepositoryMock struct {
	Objects map[string][]models.TestObject
}

func (m *TestObjectRepositoryMock) Reset() {
	m.Objects = map[string][]models.TestObject{
		PredefinedAccountHash: {PredefinedTestObject1, PredefinedTestObject2},
	}
}

func (m *TestObjectRepositoryMock) Create(
	accountHash string,
	entity map[string]interface{},
) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for _, object := range m.Objects[accountHash] {
		if object.Name == entity["name"].(string) {
			return types.UniqueEntityAlreadyExists{}
		}
	}

	m.Objects[accountHash] = append(m.Objects[accountHash], models.TestObject{
		Name: entity["name"].(string),
		Hash: entity["hash"].(string),
	})

	return nil
}

func (m *TestObjectRepositoryMock) GetAll(accountHash string, dest interface{}) error {
	if accountHash == BadAccountHash {
		return someError
	}

	destPtr := dest.(*[]models.TestObject)
	*destPtr = append(*destPtr, m.Objects[accountHash]...)
	return nil
}

func (m *TestObjectRepositoryMock) Get(accountHash, entityHash string, dest interface{}) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for _, object := range m.Objects[accountHash] {
		if object.Hash == entityHash {
			dest.(*models.TestObject).Name = object.Name
			dest.(*models.TestObject).Hash = object.Hash
			break
		}
	}

	return nil
}

func (m *TestObjectRepositoryMock) Update(
	accountHash string,
	entities []models.UpdateModel,
) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for objectIndex, object := range m.Objects[accountHash] {
		for _, entity := range entities {
			if object.Hash == entity.Hash {
				switch entity.FieldName {
				case "name":
					m.Objects[accountHash][objectIndex].Name = entity.NewValue.(string)
				case "hash":
					m.Objects[accountHash][objectIndex].Hash = entity.NewValue.(string)
				}
			}
		}
	}

	return nil
}

func (m *TestObjectRepositoryMock) Delete(accountHash, entityHash string) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for objectIndex, object := range m.Objects[accountHash] {
		if object.Hash == entityHash {
			m.Objects[accountHash] = append(
				m.Objects[accountHash][:objectIndex],
				m.Objects[accountHash][objectIndex+1:]...,
			)
		}
	}

	return nil
}

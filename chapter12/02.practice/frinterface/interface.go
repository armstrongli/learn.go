package frinterface

import (
	"learn.go/pkg/apis"
)

type ServeInterface interface {
	RegisterPersonalInformation(pi *apis.PersonalInformation) error
	UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error)
	GetFatRate(name string) (*apis.PersonalRank, error)
	GetTop() ([]*apis.PersonalRank, error)
}

type RankInitInterface interface {
	Init() error
}

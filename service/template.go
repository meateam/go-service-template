package service

import (
	pb "github.com/meateam/go-service-template/proto"
)

// Template is an interface of a template object.
type Template interface {
	GetFirstName() string
	SetFirstName(firstName string) error

	GetLastName() string
	SetLastName(lastName string) error

	GetFullName() string

	MarshalProto(user *pb.User) error
}

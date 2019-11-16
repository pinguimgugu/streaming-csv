package contract

import "github.com/pinguimgugu/streaming-csv/domain/entity"

// UserRepository interface
type UserRepository interface {
	StreamingAllUsers() chan *entity.User
}

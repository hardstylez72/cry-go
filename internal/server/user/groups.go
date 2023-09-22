package user

import (
	"database/sql"

	"github.com/hardstylez72/cry/internal/lib"
)

const GroupSupport = "support"
const GroupWorker = "worker"
const GroupPublisher = "publisher"

var support = map[string]string{
	"korotkovcv77@gmail.com": GroupSupport,
}

var publisher = map[string]string{
	"korotkovcv77@gmail.com":     GroupPublisher,
	"cryptodoubleplus@gmail.com": GroupPublisher,
	"fred.khamzin@gmail.com":     GroupPublisher,
}

type User struct {
	Id           string
	Email        string
	ControlledBy sql.NullString
}

func Groups(u *User) (*lib.Set[string], error) {

	out := lib.NewSet[string]()
	_, ok := support[u.Email]
	if ok {
		out.Set(GroupSupport)
	}

	_, ok = publisher[u.Email]
	if ok {
		out.Set(GroupPublisher)
	}

	if u.ControlledBy.Valid {
		out.Set(GroupWorker)
	}

	return out, nil
}

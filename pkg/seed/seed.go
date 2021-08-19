package seed

import (
	"base/pkg/model"
)

func All() {
	CreateAdmin(model.DB, "admin", "admin")
}

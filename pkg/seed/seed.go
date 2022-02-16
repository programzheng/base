package seed

import (
	"github.com/programzheng/base/pkg/model"
)

func All() {
	CreateAdmin(model.DB, "admin", "admin")
}

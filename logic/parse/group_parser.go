package parse

import (
	"github.com/MayugeStudio/easy_task/domain"
	"github.com/MayugeStudio/easy_task/utils"
)

func toGroup(str string) *domain.Group {
	l := utils.NewLine(str).
		TrimSpace().
		TrimPrefix("-").
		TrimSpace()
	g := domain.NewGroup(string(l))
	return g
}

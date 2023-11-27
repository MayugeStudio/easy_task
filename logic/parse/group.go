package parse

import (
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
)

func toGroup(str string) *domain.Group {
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	g := domain.NewGroup(str)
	return g
}

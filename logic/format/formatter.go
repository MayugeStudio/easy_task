package format

import (
	"strings"

	"github.com/MayugeStudio/easy_task/logic/internal/share"
)

func ToFormattedStrings(taskStrings []string) ([]string, []error) {
	result := make([]string, 0)
	errs := make([]error, 0)
	inGroup := false
	var formattedString string
	var err error
	for _, str := range taskStrings {
		if share.IsGroupTitle(str) { // TODO: rename share -> share/XXX. 'share.IsGroupTitle' is difficult to read.
			formattedString, err = toFormattedGroupTitle(str)
			inGroup = true
			goto success
		}
		if share.IsTaskString(str) {
			if inGroup {
				formattedString, err = toFormattedGroupTaskString(str)
			} else {
				formattedString, err = toFormattedTaskString(str)
				inGroup = false
			}
			goto success
		}
		if share.IsItemModificationString(str) {
			formattedString, err = toFormattedModificationString(str)
			goto success
		}

		if !strings.HasPrefix(str, "  ") {
			inGroup = false
		}
		continue

	success:
		if err != nil {
			errs = append(errs, err)
			continue
		}
		result = append(result, formattedString)
	}
	return result, errs
}

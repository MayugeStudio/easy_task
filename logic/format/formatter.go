package format

import (
	"github.com/MayugeStudio/easy_task/logic/internal/share"
)

func ToFormattedStrings(taskStrings []string) ([]string, []error) {
	result := make([]string, 0)
	errs := make([]error, 0)
	inGroup := false
	for _, str := range taskStrings {
		var fStr string
		var err error
		if share.IsGroupTitle(str) { // TODO: rename share -> share/XXX. 'share.IsGroupTitle' is difficult to read.
			fStr, err = toFormattedGroupTitle(str)
			inGroup = true
		}
		if share.IsTaskString(str) {
			if inGroup {
				fStr, err = toFormattedGroupTaskString(str)
			} else {
				fStr, err = toFormattedTaskString(str)
				inGroup = false
			}
		}
		if share.IsItemModificationString(str) {
			fStr, err = toFormattedModificationString(str)
		}

		if err != nil {
			errs = append(errs, err)
			continue
		}
		if fStr != "" {
			result = append(result, fStr)
		}
	}
	return result, errs
}

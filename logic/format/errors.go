package format

import (
	"errors"
	"fmt"
)

// FIXME: Make the error message more detailed.
var (
	errSyntax                       = errors.New("format error")
	errNoDash                       = fmt.Errorf("%w: no dash", errSyntax)
	errNoBracketStart               = fmt.Errorf("%w: no bracket start", errSyntax)
	errNoBracketEnd                 = fmt.Errorf("%w: no bracket end", errSyntax)
	errNoColon                      = fmt.Errorf("%w: no colon", errSyntax)
	errInvalidIndent                = fmt.Errorf("%w: no valid indent", errSyntax) // FIXME: fix error message.
	errInvalidModification          = fmt.Errorf("%w: invalid modification", errSyntax)
	errInvalidModificationAttribute = fmt.Errorf("%w: invalid modification attribute", errSyntax)
)

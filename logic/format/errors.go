package format

import (
	"errors"
	"fmt"
)

var (
	errSyntax                   = errors.New("format error")
	errNoDash                   = fmt.Errorf("%w: no dash", errSyntax)
	errNoBracketStart           = fmt.Errorf("%w: no bracket start", errSyntax)
	errNoBracketEnd             = fmt.Errorf("%w: no bracket end", errSyntax)
	errNoColon                  = fmt.Errorf("%w: no colon", errSyntax)
	errInvalidIndent            = fmt.Errorf("%w: invalid indent", errSyntax)
	errInvalidModifier          = fmt.Errorf("%w: invalid modifier", errSyntax)
	errInvalidModifierAttribute = fmt.Errorf("%w: invalid modifier attribute", errSyntax)
	errInvalidGroupTitle        = fmt.Errorf("%w: invalid group title", errSyntax)
)

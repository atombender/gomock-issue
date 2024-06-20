package bar

import "github.com/atombender/gomock-issue/foo"

type Starter interface {
	Start(thing *foo.Thing)
}

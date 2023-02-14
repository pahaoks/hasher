package servers

import "context"

type Server interface {
	Start(context.Context)
	Stop()
}

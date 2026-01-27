package di

type Container struct {
	LinkContainer   *LinkContainer
	ClientContainer *ClientContainer
}

var GlobalContainer *Container

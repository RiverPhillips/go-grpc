package message

import (
	"context"
	"fmt"
)

type Server struct{}

func (s Server) HelloWorld(_ context.Context, input *HelloWorldInput) (*HelloWorldOutput, error) {
	name := input.Name
	// todo - save name in database
	return &HelloWorldOutput{Greeting: fmt.Sprintf("Hello %s!", name)}, nil
}

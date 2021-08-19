package function

import (
	"context"
	"fmt"
	"os"

	event "github.com/cloudevents/sdk-go/v2"
)

func Handle(ctx context.Context, event event.Event) error {
	if err := event.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "invalid event received. %v", err)
		return err
	}
	fmt.Printf("FUNCTIONS EVENT TEST: %v\n", event)
	return nil
}

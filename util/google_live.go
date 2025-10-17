package util

// import (
// 	livestream "cloud.google.com/go/video/livestream/apiv1"
// 	"cloud.google.com/go/video/livestream/apiv1/livestreampb"
// 	"context"
// 	"fmt"
// 	"io"
// )

// // createInput creates an input endpoint. You send an input video stream to this
// // endpoint.
// func createInput(w io.Writer, projectID, location, inputID string) error {
// 	// projectID := "my-project-id"
// 	// location := "us-central1"
// 	// inputID := "my-input"
// 	ctx := context.Background()
// 	client, err := livestream.NewClient(ctx)
// 	if err != nil {
// 		return fmt.Errorf("NewClient: %w", err)
// 	}
// 	defer func(client *livestream.Client) {
// 		err := client.Close()
// 		if err != nil {

// 		}
// 	}(client)

// 	req := &livestreampb.CreateInputRequest{
// 		Parent:  fmt.Sprintf("projects/%s/locations/%s", projectID, location),
// 		InputId: inputID,
// 		Input: &livestreampb.Input{
// 			Type: livestreampb.Input_RTMP_PUSH,
// 		},
// 	}
// 	// Creates the input.
// 	op, err := client.CreateInput(ctx, req)
// 	if err != nil {
// 		return fmt.Errorf("CreateInput: %w", err)
// 	}
// 	response, err := op.Wait(ctx)
// 	if err != nil {
// 		return fmt.Errorf("Wait: %w", err)
// 	}

// 	fmt.Fprintf(w, "Input: %v", response.Name)
// 	return nil
// }

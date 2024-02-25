package main

import (
	"context"
	"flag"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func main() {
	projectPtr := flag.String("project", "", "Project ID")
	topicPtr := flag.String("topic", "", "Topic ID")
	messagePtr := flag.String("message", "", "Message to publish")
	flag.Parse()

	messageId, err := publish(*projectPtr, *topicPtr, *messagePtr)
	if err != nil {
		fmt.Printf("[error] %v\n", err)
	} else {
		fmt.Printf("[success] Message published: messageID = %v\n", messageId)
	}
}

// https://cloud.google.com/pubsub/docs/publish-receive-messages-client-library?hl=ja#publish_messages
func publish(projectID string, topicID string, msg string) (string, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("pubsub: NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("pubsub: result.Get: %w", err)
	}

	return id, nil
}

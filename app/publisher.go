package main

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// Ref.
// - https://medium.com/@sugamon.g/cloud-pubsub%E3%81%AE%E5%8D%98%E4%BD%93%E3%83%86%E3%82%B9%E3%83%88%E3%81%AE%E3%82%84%E3%82%8A%E6%96%B9-2705c567aca5
// - https://iruca21.hateblo.jp/entry/2022/03/29/210000

type publisher struct {
	client *pubsub.Client
}

func newPublisher(ctx context.Context, projectID string, opts ...option.ClientOption) (*publisher, error) {
	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &publisher{client: client}, nil
}

func (p *publisher) close() {
	p.client.Close()
}

func (p *publisher) publish(ctx context.Context, topicID string, msg string) (string, error) {
	t := p.client.Topic(topicID)
	id, err := t.Publish(ctx, &pubsub.Message{Data: []byte(msg)}).Get(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}

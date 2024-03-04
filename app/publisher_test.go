// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------
// https://github.com/googleapis/google-cloud-go/blob/main/pubsub/pstest/examples_test.go

package main

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// Ref.
// - https://medium.com/@sugamon.g/cloud-pubsub%E3%81%AE%E5%8D%98%E4%BD%93%E3%83%86%E3%82%B9%E3%83%88%E3%81%AE%E3%82%84%E3%82%8A%E6%96%B9-2705c567aca5
// - https://iruca21.hateblo.jp/entry/2022/03/29/210000

func TestPublisher(t *testing.T) {
	const TARGET_TOPIC = "target-topic"

	cases := map[string]struct {
		projectID                string
		topicID                  string
		message                  string
		expectedId               string
		expectedPublishedMessage string
	}{
		"success": {
			projectID:                "my-project",
			topicID:                  TARGET_TOPIC,
			message:                  "Hello, World!",
			expectedId:               "m0",
			expectedPublishedMessage: "Hello, World!",
		},
		"failure": {
			projectID:                "my-project",
			topicID:                  "not-existed-topic",
			message:                  "Hello, World!",
			expectedId:               "",
			expectedPublishedMessage: "",
		},
	}

	for k, c := range cases {
		t.Run(k, func(t *testing.T) {
			// Start a fake Pub/Sub server
			srv := pstest.NewServer()
			defer srv.Close()
			conn, _ := grpc.Dial(srv.Addr, grpc.WithInsecure())
			defer conn.Close()

			ctx := context.Background()

			p, _ := newPublisher(ctx, c.projectID, option.WithGRPCConn(conn))
			defer p.close()

			// Create a topic before testing
			p.client.CreateTopic(ctx, TARGET_TOPIC)

			id, err := p.publish(ctx, c.topicID, c.message)
			if c.topicID != TARGET_TOPIC && err == nil {
				t.Errorf("publish(%s, %s, %s): expectedId error, actual nil", c.projectID, c.topicID, c.message)
			}
			if id != c.expectedId {
				t.Errorf("publish(%s, %s, %s): expectedId %s, actual %s", c.projectID, c.topicID, c.message, c.expectedId, id)
			}

			// Check the published message
			// https://github.com/googleapis/google-cloud-go/blob/e5d0c2fc2182174b9307363b48c0a0e4056cb3f4/pubsub/pstest/fake.go#L228
			if id != "" {
				published := srv.Message(id)
				if string(published.Data) != c.expectedPublishedMessage {
					t.Errorf("publish(%s, %s, %s): expectedPublishedMessage %s, actual %s", c.projectID, c.topicID, c.message, c.expectedPublishedMessage, published.Data)
				}
			}
		})
	}
}

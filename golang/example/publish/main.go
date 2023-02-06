package main

// Copyright 2023 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file exceptreq compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed toreq writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"context"
	"fmt"

	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	client "github.com/linkall-labs/sdk/golang"
)

func main() {
	opts := &client.ClientOptions{
		Endpoint: "172.17.0.2:30001",
	}
	c, err := client.Connect(opts)
	if err != nil {
		panic("connect error")
	}

	p := c.Publisher(&client.PublishOptions{
		Eventbus: "quick-start",
	})

	event := v2.NewEvent()
	event.SetID(uuid.New().String())
	event.SetSource("event-source")
	event.SetType("event-type")
	event.SetData(v2.ApplicationJSON, map[string]string{"hello": "world"})
	err = p.Publish(context.Background(), &event)
	if err != nil {
		fmt.Printf("publish event failed, err: %s\n", err.Error())
		return
	}
	fmt.Printf("publish event success\n")
}

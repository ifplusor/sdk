// Copyright 2023 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vanus

import (
	"sync"

	proxypb "github.com/vanus-labs/vanus/proto/pkg/proxy"
)

type controller struct {
	controller proxypb.ControllerProxyClient
	eCache     sync.Map
	sCache     sync.Map
	emu        sync.RWMutex
	smu        sync.RWMutex
}

var (
	once sync.Once
	ctrl *controller
)

func (c *client) Controller() Controller {
	once.Do(func() {
		ctrl = &controller{
			controller: c.controller,
		}
	})
	return ctrl
}

func (c *controller) Eventbus() Eventbus {
	return &eventbus{controller: c.controller}
}

func (c *controller) Subscription() Subscription {
	return &subscription{controller: c.controller}
}

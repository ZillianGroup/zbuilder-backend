// Copyright 2022 The ILLA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ws

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/illacloud/builder-backend/internal/util/builderoperation"
)

// message protocol from client in json:
//
// {
//     "signal":number,
//     "option":number(work as int32 bit),
//     "payload":string
// }

// for message

const OPTION_BROADCAST_ROOM = 1 // 00000000000000000000000000000001; // use as signed int32 in typescript

// for broadcast rewrite
const BROADCAST_TYPE_SUFFIX = "/remote"
const BROADCAST_TYPE_ENTER = "enter"
const BROADCAST_TYPE_ATTACH_COMPONENT = "attachComponent"

type Broadcast struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type Message struct {
	ClientID      uuid.UUID     `json:"clientID"`
	Signal        int           `json:"signal"`
	APPID         int           `json:"appID"` // also as APP ID
	Option        int           `json:"option"`
	Payload       []interface{} `json:"payload"`
	Target        int           `json:"target"`
	Broadcast     *Broadcast    `json:"broadcast"`
	NeedBroadcast bool
}

func NewMessage(clientID uuid.UUID, appID int, rawMessage []byte) (*Message, error) {
	// init Action
	var message Message
	if err := json.Unmarshal(rawMessage, &message); err != nil {
		return nil, err
	}
	message.ClientID = clientID
	message.APPID = appID
	if message.Broadcast == nil {
		message.NeedBroadcast = false
	} else {
		message.NeedBroadcast = true
	}
	return &message, nil
}

func (m *Message) SetSignal(s int) {
	m.Signal = builderoperation.SIGNAL_COOPERATE_ATTACH
}

func (m *Message) SetBroadcastType(t string) {
	if m.Broadcast != nil {
		m.Broadcast.Type = t
	}
}

func (m *Message) SetBroadcastPayload(any interface{}) {
	if m.Broadcast != nil {
		m.Broadcast.Payload = any
	}
}

func (m *Message) RewriteBroadcast() {
	if m.NeedBroadcast {
		m.Broadcast.Type = m.Broadcast.Type + BROADCAST_TYPE_SUFFIX
	}
}

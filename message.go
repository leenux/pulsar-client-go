// Copyright 2018 Comcast Cable Communications Management, LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pulsar

import (
	"bytes"

	"github.com/golang/protobuf/proto"
	"github.com/leenux/pulsar-client-go/api"
)

// Message represents a received MESSAGE from the Pulsar server.
type Message struct {
	Topic      string
	consumerID uint64

	Msg     *api.CommandMessage
	Meta    *api.MessageMetadata
	Payload []byte
}

// Equal returns true if the provided other Message
// is equal to the receiver Message.
func (m *Message) Equal(other *Message) bool {
	return m.consumerID == other.consumerID &&
		proto.Equal(m.Msg, other.Msg) &&
		proto.Equal(m.Meta, other.Meta) &&
		bytes.Equal(m.Payload, other.Payload)
}

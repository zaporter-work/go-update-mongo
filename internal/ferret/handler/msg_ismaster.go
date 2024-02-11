// Copyright 2021 FerretDB Inc.
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

package handler

import (
	"context"

	"github.com/zaporter-work/go-update-mongo/internal/ferret/handler/common"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/util/lazyerrors"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/util/must"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/wire"
)

// MsgIsMaster implements `isMaster` command.
func (h *Handler) MsgIsMaster(ctx context.Context, msg *wire.OpMsg) (*wire.OpMsg, error) {
	doc, err := msg.Document()
	if err != nil {
		return nil, err
	}

	if err := common.CheckClientMetadata(ctx, doc); err != nil {
		return nil, lazyerrors.Error(err)
	}

	var reply wire.OpMsg
	must.NoError(reply.SetSections(wire.MakeOpMsgSection(
		common.IsMasterDocument(h.TCPHost, h.ReplSetName),
	)))

	return &reply, nil
}

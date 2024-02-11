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
	"errors"
	"strconv"
	"strings"

	"github.com/zaporter-work/go-update-mongo/internal/ferret/handler/common"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/handler/handlererrors"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/types"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/util/lazyerrors"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/util/must"
	"github.com/zaporter-work/go-update-mongo/internal/ferret/wire"
)

// MsgDebugError implements `debugError` command.
func (h *Handler) MsgDebugError(ctx context.Context, msg *wire.OpMsg) (*wire.OpMsg, error) {
	document, err := msg.Document()
	if err != nil {
		return nil, lazyerrors.Error(err)
	}

	expected, err := common.GetRequiredParam[string](document, document.Command())
	if err != nil {
		return nil, err
	}

	// check if parameter is an error code
	if n, err := strconv.ParseInt(expected, 10, 32); err == nil {
		errCode := handlererrors.ErrorCode(n)
		return nil, errors.New(errCode.String())
	}

	switch {
	case expected == "ok":
		var reply wire.OpMsg

		replyDoc := must.NotFail(types.NewDocument(
			"ok", float64(1),
		))

		must.NoError(reply.SetSections(wire.MakeOpMsgSection(
			replyDoc,
		)))

		return &reply, nil

	case strings.HasPrefix(expected, "panic"):
		panic("debugError " + expected)

	case strings.HasPrefix(expected, "lazy"):
		return nil, lazyerrors.New(expected)

	default:
		return nil, errors.New(expected)
	}
}

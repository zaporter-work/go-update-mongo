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

package stages

import (
	"context"

	"github.com/zaporter/go-update-mongo/internal/ferret/handler/common"
	"github.com/zaporter/go-update-mongo/internal/ferret/handler/common/aggregations"
	"github.com/zaporter/go-update-mongo/internal/ferret/handler/common/aggregations/stages/projection"
	"github.com/zaporter/go-update-mongo/internal/ferret/handler/handlererrors"
	"github.com/zaporter/go-update-mongo/internal/ferret/types"
	"github.com/zaporter/go-update-mongo/internal/ferret/util/iterator"
)

// project represents $project stage.
//
//	{
//	  $project: {
//	    <output field1>: <expression1>,
//	    ...
//	    <output fieldN>: <expressionN>
//	  }
type project struct {
	projection *types.Document
	inclusion  bool
}

// newProject validates projection document and creates a new $project stage.
func newProject(stage *types.Document) (aggregations.Stage, error) {
	fields, err := common.GetRequiredParam[*types.Document](stage, "$project")
	if err != nil {
		return nil, handlererrors.NewCommandErrorMsgWithArgument(
			handlererrors.ErrProjectBadExpression,
			"$project specification must be an object",
			"$project (stage)",
		)
	}

	validated, inclusion, err := projection.ValidateProjection(fields)
	if err != nil {
		return nil, err
	}

	return &project{
		projection: validated,
		inclusion:  inclusion,
	}, nil
}

// Process implements Stage interface.
//
//nolint:lll // for readability
func (p *project) Process(_ context.Context, iter types.DocumentsIterator, closer *iterator.MultiCloser) (types.DocumentsIterator, error) {
	return projection.ProjectionIterator(iter, closer, p.projection)
}

// check interfaces
var (
	_ aggregations.Stage = (*project)(nil)
)

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

package types

import "github.com/zaporter-work/go-update-mongo/internal/ferret/util/iterator"

// DocumentsIterator represents an iterator over documents (slice, query results, etc).
//
// Key/index is not used there because it is unclear what it should be in the filter/sort/limit/skip chain,
// and it is not used anyway.
type DocumentsIterator iterator.Interface[struct{}, *Document]

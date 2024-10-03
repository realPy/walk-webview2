//go:build windows
// +build windows

package webview2

import (
	"github.com/google/uuid"
)

// compatibility Ref C++ with GO garbage collector
var keepRefGO map[uuid.UUID]interface{} = map[uuid.UUID]interface{}{}

func setRefGO(uuid uuid.UUID, i interface{}) {
	keepRefGO[uuid] = i
}

func deleteRefGO(uuid uuid.UUID) {
	delete(keepRefGO, uuid)
}

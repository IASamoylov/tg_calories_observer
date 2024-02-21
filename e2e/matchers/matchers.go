package matchers

import (
	"encoding/json"

	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

// MatchValueAsJSON succeeds if actual is a string or stringer of JSON that matches
// the expected JSON. The JSONs are decoded and the resulting objects are compared via
// reflect.DeepEqual so things like key-ordering and whitespace shouldn't matter.
func MatchValueAsJSON(value interface{}) types.GomegaMatcher {
	bytes, _ := json.Marshal(value)

	return &matchers.MatchJSONMatcher{
		JSONToMatch: bytes,
	}
}

// Code generated by "stringer -type=ResourceType"; DO NOT EDIT.

package est

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SQLDBResource-1]
	_ = x[CronJobResource-2]
	_ = x[PubSubTopicResource-3]
	_ = x[PubSubSubscriptionResource-4]
}

const _ResourceType_name = "SQLDBResourceCronJobResourcePubSubTopicResourcePubSubSubscriptionResource"

var _ResourceType_index = [...]uint8{0, 13, 28, 47, 73}

func (i ResourceType) String() string {
	i -= 1
	if i < 0 || i >= ResourceType(len(_ResourceType_index)-1) {
		return "ResourceType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ResourceType_name[_ResourceType_index[i]:_ResourceType_index[i+1]]
}

// Code generated by "stringer -type opcode,closecode -output consts_generated.go"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opcodeDispatch-0]
	_ = x[opcodeHeartbeat-1]
	_ = x[opcodeIdentify-2]
	_ = x[opcodePresenceUpdate-3]
	_ = x[opcodeVoiceStateUpdate-4]
	_ = x[opcodeResume-6]
	_ = x[opcodeReconnect-7]
	_ = x[opcodeRequestGuildMembers-8]
	_ = x[opcodeInvalidSession-9]
	_ = x[opcodeHello-10]
	_ = x[opcodeHeartbeatACK-11]
}

const (
	_opcode_name_0 = "opcodeDispatchopcodeHeartbeatopcodeIdentifyopcodePresenceUpdateopcodeVoiceStateUpdate"
	_opcode_name_1 = "opcodeResumeopcodeReconnectopcodeRequestGuildMembersopcodeInvalidSessionopcodeHelloopcodeHeartbeatACK"
)

var (
	_opcode_index_0 = [...]uint8{0, 14, 29, 43, 63, 85}
	_opcode_index_1 = [...]uint8{0, 12, 27, 52, 72, 83, 101}
)

func (i opcode) String() string {
	switch {
	case 0 <= i && i <= 4:
		return _opcode_name_0[_opcode_index_0[i]:_opcode_index_0[i+1]]
	case 6 <= i && i <= 11:
		i -= 6
		return _opcode_name_1[_opcode_index_1[i]:_opcode_index_1[i+1]]
	default:
		return "opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[closecodeUnknownError-4000]
	_ = x[closecodeUnknownOpCode-4001]
	_ = x[closecodeDecodeError-4002]
	_ = x[closecodeNotAuthenticated-4003]
	_ = x[closecodeAuthenticationFailed-4004]
	_ = x[closecodeAlreadyAuthenticated-4005]
	_ = x[closecodeInvalidSquence-4007]
	_ = x[closecodeRateLimited-4008]
	_ = x[closecodeSessionTimedOut-4009]
	_ = x[closecodeInvalidShard-4010]
	_ = x[closecodeShardingRequired-4011]
	_ = x[closecodeInvalidAPIVersion-4012]
	_ = x[closecodeInvalidIntents-4013]
	_ = x[closecodeDisallowedIntents-4014]
}

const (
	_closecode_name_0 = "closecodeUnknownErrorclosecodeUnknownOpCodeclosecodeDecodeErrorclosecodeNotAuthenticatedclosecodeAuthenticationFailedclosecodeAlreadyAuthenticated"
	_closecode_name_1 = "closecodeInvalidSquenceclosecodeRateLimitedclosecodeSessionTimedOutclosecodeInvalidShardclosecodeShardingRequiredclosecodeInvalidAPIVersionclosecodeInvalidIntentsclosecodeDisallowedIntents"
)

var (
	_closecode_index_0 = [...]uint8{0, 21, 43, 63, 88, 117, 146}
	_closecode_index_1 = [...]uint8{0, 23, 43, 67, 88, 113, 139, 162, 188}
)

func (i closecode) String() string {
	switch {
	case 4000 <= i && i <= 4005:
		i -= 4000
		return _closecode_name_0[_closecode_index_0[i]:_closecode_index_0[i+1]]
	case 4007 <= i && i <= 4014:
		i -= 4007
		return _closecode_name_1[_closecode_index_1[i]:_closecode_index_1[i+1]]
	default:
		return "closecode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
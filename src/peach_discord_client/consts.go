package main

// Gateway opcodes, denote payload type, see https://discordapp.com/developers/docs/topics/opcodes-and-status-codes#gateway-opcodes
const (
	opCodeDispatch            = 0  // Receive      | An event was dispatched.
	opCodeHeartbeat           = 1  // Send/Receive | Fired periodically by the client to keep the connection alive.
	opCodeIdentify            = 2  // Send         | Starts a new session during the initial handshake.
	opCodePresenceUpdate      = 3  // Send         | Update the client's presence.
	opCodeVoiceStateUpdate    = 4  // Send         | Used to join/leave or move between voice channels.
	opCodeResume              = 6  // Send         | Resume a previous session that was disconnected.
	opCodeReconnect           = 7  // Receive      | You must reconnect with a new session immediately.
	opCodeRequestGuildMembers = 8  // Send         | Request information about offline guild members in a large guild.
	opCodeInvalidSession      = 9  // Receive      | The session has been invalidated. You should reconnect and identify/resume accordingly.
	opCodeHello               = 10 // Receive      | Sent immediately after connecting, contains the heartbeat_interval to use.
	opCodeHeartbeatACK        = 11 // Receive      | Sent in response to receiving a heartbeat to acknowledge that it has been received.
)

// Gateway Close Event Codes, denote reason for gateway closure, see https://discordapp.com/developers/docs/topics/opcodes-and-status-codes#gateway-opcodes
const (
	closeCodeUnknownError         = 4000 // Not sure what went wrong. Try reconnecting.
	closeCodeUnknownOpCode        = 4001 // Sent invalid opcode or invalid payload for opcode.
	closeCodeDecodeError          = 4002 // Sent invalid payload.
	closeCodeNotAuthenticated     = 4003 // Sent payload prior to identifying.
	closeCodeAuthenticationFailed = 4004 // Account token in identify payload is incorrect.
	closeCodeAlreadyAuthenticated = 4005 // Sent more than one identify payload.
	closeCodeInvalidSquence       = 4007 // Sent invalid sequence when resuming.
	closeCodeRateLimited          = 4008 // Sending payloads to quickly.
	closeCodeSessionTimedOut      = 4009 // Session timed out. Reconnect or start new session.
	closeCodeInvalidShard         = 4010 // Sent invalid shard in identify payload.
	closeCodeShardingRequired     = 4011 // Sharding required because bot is in too many guilds.
	closeCodeInvalidAPIVersion    = 4012 // Sent an invalid gateway version.
	closeCodeInvalidIntents       = 4013 // Sent invalid gateway intent.
	closeCodeDisallowedIntents    = 4014 // Sent intent the account isn't eligible for.
)

// Code generated by genevent. DO NOT EDIT.

package gateway

import "github.com/diamondburned/arikawa/v3/utils/ws"

func init() {
	OpUnmarshalers.Add(
		func() ws.Event { return new(HeartbeatCommand) },
		func() ws.Event { return new(HeartbeatAckEvent) },
		func() ws.Event { return new(ReconnectEvent) },
		func() ws.Event { return new(HelloEvent) },
		func() ws.Event { return new(ResumeCommand) },
		func() ws.Event { return new(InvalidSessionEvent) },
		func() ws.Event { return new(RequestGuildMembersCommand) },
		func() ws.Event { return new(UpdateVoiceStateCommand) },
		func() ws.Event { return new(UpdatePresenceCommand) },
		func() ws.Event { return new(GuildSubscribeCommand) },
		func() ws.Event { return new(ResumedEvent) },
		func() ws.Event { return new(ChannelCreateEvent) },
		func() ws.Event { return new(ChannelUpdateEvent) },
		func() ws.Event { return new(ChannelDeleteEvent) },
		func() ws.Event { return new(ChannelPinsUpdateEvent) },
		func() ws.Event { return new(ChannelUnreadUpdateEvent) },
		func() ws.Event { return new(ThreadCreateEvent) },
		func() ws.Event { return new(ThreadUpdateEvent) },
		func() ws.Event { return new(ThreadDeleteEvent) },
		func() ws.Event { return new(ThreadListSyncEvent) },
		func() ws.Event { return new(ThreadMemberUpdateEvent) },
		func() ws.Event { return new(ThreadMembersUpdateEvent) },
		func() ws.Event { return new(GuildCreateEvent) },
		func() ws.Event { return new(GuildUpdateEvent) },
		func() ws.Event { return new(GuildDeleteEvent) },
		func() ws.Event { return new(GuildBanAddEvent) },
		func() ws.Event { return new(GuildBanRemoveEvent) },
		func() ws.Event { return new(GuildEmojisUpdateEvent) },
		func() ws.Event { return new(GuildIntegrationsUpdateEvent) },
		func() ws.Event { return new(GuildMemberAddEvent) },
		func() ws.Event { return new(GuildMemberRemoveEvent) },
		func() ws.Event { return new(GuildMemberUpdateEvent) },
		func() ws.Event { return new(GuildMembersChunkEvent) },
		func() ws.Event { return new(GuildRoleCreateEvent) },
		func() ws.Event { return new(GuildRoleUpdateEvent) },
		func() ws.Event { return new(GuildRoleDeleteEvent) },
		func() ws.Event { return new(InviteCreateEvent) },
		func() ws.Event { return new(InviteDeleteEvent) },
		func() ws.Event { return new(MessageCreateEvent) },
		func() ws.Event { return new(MessageUpdateEvent) },
		func() ws.Event { return new(MessageDeleteEvent) },
		func() ws.Event { return new(MessageDeleteBulkEvent) },
		func() ws.Event { return new(MessageReactionAddEvent) },
		func() ws.Event { return new(MessageReactionRemoveEvent) },
		func() ws.Event { return new(MessageReactionRemoveAllEvent) },
		func() ws.Event { return new(MessageReactionRemoveEmojiEvent) },
		func() ws.Event { return new(MessageAckEvent) },
		func() ws.Event { return new(PresenceUpdateEvent) },
		func() ws.Event { return new(PresencesReplaceEvent) },
		func() ws.Event { return new(SessionsReplaceEvent) },
		func() ws.Event { return new(TypingStartEvent) },
		func() ws.Event { return new(UserUpdateEvent) },
		func() ws.Event { return new(VoiceStateUpdateEvent) },
		func() ws.Event { return new(VoiceServerUpdateEvent) },
		func() ws.Event { return new(WebhooksUpdateEvent) },
		func() ws.Event { return new(InteractionCreateEvent) },
		func() ws.Event { return new(UserGuildSettingsUpdateEvent) },
		func() ws.Event { return new(UserSettingsUpdateEvent) },
		func() ws.Event { return new(UserNoteUpdateEvent) },
		func() ws.Event { return new(RelationshipAddEvent) },
		func() ws.Event { return new(RelationshipRemoveEvent) },
		func() ws.Event { return new(ReadyEvent) },
		func() ws.Event { return new(ReadySupplementalEvent) },
		func() ws.Event { return new(IdentifyCommand) },
	)
}

// Op implements Event. It always returns Op 1.
func (*HeartbeatCommand) Op() ws.OpCode { return 1 }

// EventType implements Event.
func (*HeartbeatCommand) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 11.
func (*HeartbeatAckEvent) Op() ws.OpCode { return 11 }

// EventType implements Event.
func (*HeartbeatAckEvent) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 7.
func (*ReconnectEvent) Op() ws.OpCode { return 7 }

// EventType implements Event.
func (*ReconnectEvent) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 10.
func (*HelloEvent) Op() ws.OpCode { return 10 }

// EventType implements Event.
func (*HelloEvent) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 6.
func (*ResumeCommand) Op() ws.OpCode { return 6 }

// EventType implements Event.
func (*ResumeCommand) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 9.
func (*InvalidSessionEvent) Op() ws.OpCode { return 9 }

// EventType implements Event.
func (*InvalidSessionEvent) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 8.
func (*RequestGuildMembersCommand) Op() ws.OpCode { return 8 }

// EventType implements Event.
func (*RequestGuildMembersCommand) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 4.
func (*UpdateVoiceStateCommand) Op() ws.OpCode { return 4 }

// EventType implements Event.
func (*UpdateVoiceStateCommand) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 3.
func (*UpdatePresenceCommand) Op() ws.OpCode { return 3 }

// EventType implements Event.
func (*UpdatePresenceCommand) EventType() ws.EventType { return "" }

// Op implements Event. It always returns Op 14.
func (*GuildSubscribeCommand) Op() ws.OpCode { return 14 }

// EventType implements Event.
func (*GuildSubscribeCommand) EventType() ws.EventType { return "" }

// Op implements Event. It always returns 0.
func (*ResumedEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ResumedEvent) EventType() ws.EventType { return "RESUMED" }

// Op implements Event. It always returns 0.
func (*ChannelCreateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ChannelCreateEvent) EventType() ws.EventType { return "CHANNEL_CREATE" }

// Op implements Event. It always returns 0.
func (*ChannelUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ChannelUpdateEvent) EventType() ws.EventType { return "CHANNEL_UPDATE" }

// Op implements Event. It always returns 0.
func (*ChannelDeleteEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ChannelDeleteEvent) EventType() ws.EventType { return "CHANNEL_DELETE" }

// Op implements Event. It always returns 0.
func (*ChannelPinsUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ChannelPinsUpdateEvent) EventType() ws.EventType { return "CHANNEL_PINS_UPDATE" }

// Op implements Event. It always returns 0.
func (*ChannelUnreadUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ChannelUnreadUpdateEvent) EventType() ws.EventType { return "CHANNEL_UNREAD_UPDATE" }

// Op implements Event. It always returns 0.
func (*ThreadCreateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ThreadCreateEvent) EventType() ws.EventType { return "THREAD_CREATE" }

// Op implements Event. It always returns 0.
func (*ThreadUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ThreadUpdateEvent) EventType() ws.EventType { return "THREAD_UPDATE" }

// Op implements Event. It always returns 0.
func (*ThreadDeleteEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ThreadDeleteEvent) EventType() ws.EventType { return "THREAD_DELETE" }

// Op implements Event. It always returns 0.
func (*ThreadListSyncEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ThreadListSyncEvent) EventType() ws.EventType { return "THREAD_LIST_SYNC" }

// Op implements Event. It always returns 0.
func (*ThreadMemberUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ThreadMemberUpdateEvent) EventType() ws.EventType { return "THREAD_MEMBER_UPDATE" }

// Op implements Event. It always returns 0.
func (*ThreadMembersUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ThreadMembersUpdateEvent) EventType() ws.EventType { return "THREAD_MEMBERS_UPDATE" }

// Op implements Event. It always returns 0.
func (*GuildCreateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildCreateEvent) EventType() ws.EventType { return "GUILD_CREATE" }

// Op implements Event. It always returns 0.
func (*GuildUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildUpdateEvent) EventType() ws.EventType { return "GUILD_UPDATE" }

// Op implements Event. It always returns 0.
func (*GuildDeleteEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildDeleteEvent) EventType() ws.EventType { return "GUILD_DELETE" }

// Op implements Event. It always returns 0.
func (*GuildBanAddEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildBanAddEvent) EventType() ws.EventType { return "GUILD_BAN_ADD" }

// Op implements Event. It always returns 0.
func (*GuildBanRemoveEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildBanRemoveEvent) EventType() ws.EventType { return "GUILD_BAN_REMOVE" }

// Op implements Event. It always returns 0.
func (*GuildEmojisUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildEmojisUpdateEvent) EventType() ws.EventType { return "GUILD_EMOJIS_UPDATE" }

// Op implements Event. It always returns 0.
func (*GuildIntegrationsUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildIntegrationsUpdateEvent) EventType() ws.EventType { return "GUILD_INTEGRATIONS_UPDATE" }

// Op implements Event. It always returns 0.
func (*GuildMemberAddEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildMemberAddEvent) EventType() ws.EventType { return "GUILD_MEMBER_ADD" }

// Op implements Event. It always returns 0.
func (*GuildMemberRemoveEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildMemberRemoveEvent) EventType() ws.EventType { return "GUILD_MEMBER_REMOVE" }

// Op implements Event. It always returns 0.
func (*GuildMemberUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildMemberUpdateEvent) EventType() ws.EventType { return "GUILD_MEMBER_UPDATE" }

// Op implements Event. It always returns 0.
func (*GuildMembersChunkEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildMembersChunkEvent) EventType() ws.EventType { return "GUILD_MEMBERS_CHUNK" }

// Op implements Event. It always returns 0.
func (*GuildRoleCreateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildRoleCreateEvent) EventType() ws.EventType { return "GUILD_ROLE_CREATE" }

// Op implements Event. It always returns 0.
func (*GuildRoleUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildRoleUpdateEvent) EventType() ws.EventType { return "GUILD_ROLE_UPDATE" }

// Op implements Event. It always returns 0.
func (*GuildRoleDeleteEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*GuildRoleDeleteEvent) EventType() ws.EventType { return "GUILD_ROLE_DELETE" }

// Op implements Event. It always returns 0.
func (*InviteCreateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*InviteCreateEvent) EventType() ws.EventType { return "INVITE_CREATE" }

// Op implements Event. It always returns 0.
func (*InviteDeleteEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*InviteDeleteEvent) EventType() ws.EventType { return "INVITE_DELETE" }

// Op implements Event. It always returns 0.
func (*MessageCreateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageCreateEvent) EventType() ws.EventType { return "MESSAGE_CREATE" }

// Op implements Event. It always returns 0.
func (*MessageUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageUpdateEvent) EventType() ws.EventType { return "MESSAGE_UPDATE" }

// Op implements Event. It always returns 0.
func (*MessageDeleteEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageDeleteEvent) EventType() ws.EventType { return "MESSAGE_DELETE" }

// Op implements Event. It always returns 0.
func (*MessageDeleteBulkEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageDeleteBulkEvent) EventType() ws.EventType { return "MESSAGE_DELETE_BULK" }

// Op implements Event. It always returns 0.
func (*MessageReactionAddEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageReactionAddEvent) EventType() ws.EventType { return "MESSAGE_REACTION_ADD" }

// Op implements Event. It always returns 0.
func (*MessageReactionRemoveEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageReactionRemoveEvent) EventType() ws.EventType { return "MESSAGE_REACTION_REMOVE" }

// Op implements Event. It always returns 0.
func (*MessageReactionRemoveAllEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageReactionRemoveAllEvent) EventType() ws.EventType { return "MESSAGE_REACTION_REMOVE_ALL" }

// Op implements Event. It always returns 0.
func (*MessageReactionRemoveEmojiEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageReactionRemoveEmojiEvent) EventType() ws.EventType {
	return "MESSAGE_REACTION_REMOVE_EMOJI"
}

// Op implements Event. It always returns 0.
func (*MessageAckEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*MessageAckEvent) EventType() ws.EventType { return "MESSAGE_ACK" }

// Op implements Event. It always returns 0.
func (*PresenceUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*PresenceUpdateEvent) EventType() ws.EventType { return "PRESENCE_UPDATE" }

// Op implements Event. It always returns 0.
func (*PresencesReplaceEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*PresencesReplaceEvent) EventType() ws.EventType { return "PRESENCES_REPLACE" }

// Op implements Event. It always returns 0.
func (*SessionsReplaceEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*SessionsReplaceEvent) EventType() ws.EventType { return "SESSIONS_REPLACE" }

// Op implements Event. It always returns 0.
func (*TypingStartEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*TypingStartEvent) EventType() ws.EventType { return "TYPING_START" }

// Op implements Event. It always returns 0.
func (*UserUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*UserUpdateEvent) EventType() ws.EventType { return "USER_UPDATE" }

// Op implements Event. It always returns 0.
func (*VoiceStateUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*VoiceStateUpdateEvent) EventType() ws.EventType { return "VOICE_STATE_UPDATE" }

// Op implements Event. It always returns 0.
func (*VoiceServerUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*VoiceServerUpdateEvent) EventType() ws.EventType { return "VOICE_SERVER_UPDATE" }

// Op implements Event. It always returns 0.
func (*WebhooksUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*WebhooksUpdateEvent) EventType() ws.EventType { return "WEBHOOKS_UPDATE" }

// Op implements Event. It always returns 0.
func (*InteractionCreateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*InteractionCreateEvent) EventType() ws.EventType { return "INTERACTION_CREATE" }

// Op implements Event. It always returns 0.
func (*UserGuildSettingsUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*UserGuildSettingsUpdateEvent) EventType() ws.EventType { return "USER_GUILD_SETTINGS_UPDATE" }

// Op implements Event. It always returns 0.
func (*UserSettingsUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*UserSettingsUpdateEvent) EventType() ws.EventType { return "USER_SETTINGS_UPDATE" }

// Op implements Event. It always returns 0.
func (*UserNoteUpdateEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*UserNoteUpdateEvent) EventType() ws.EventType { return "USER_NOTE_UPDATE" }

// Op implements Event. It always returns 0.
func (*RelationshipAddEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*RelationshipAddEvent) EventType() ws.EventType { return "RELATIONSHIP_ADD" }

// Op implements Event. It always returns 0.
func (*RelationshipRemoveEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*RelationshipRemoveEvent) EventType() ws.EventType { return "RELATIONSHIP_REMOVE" }

// Op implements Event. It always returns 0.
func (*ReadyEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ReadyEvent) EventType() ws.EventType { return "READY" }

// Op implements Event. It always returns 0.
func (*ReadySupplementalEvent) Op() ws.OpCode { return dispatchOp }

// EventType implements Event.
func (*ReadySupplementalEvent) EventType() ws.EventType { return "READY_SUPPLEMENTAL" }

// Op implements Event. It always returns Op 2.
func (*IdentifyCommand) Op() ws.OpCode { return 2 }

// EventType implements Event.
func (*IdentifyCommand) EventType() ws.EventType { return "" }

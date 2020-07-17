package main

import (
	"fmt"
	"strings"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

func (c *Client) onChannelCreate(ctx *EventChannelCreate) error {

	c.ChannelCache.Set(ctx.ID, ctx.Channel, cache.DefaultExpiration)

	return nil
}

func (c *Client) onChannelDelete(ctx *EventChannelDelete) error {

	c.ChannelCache.Set(ctx.ID, ctx.Channel, 0)

	return nil
}

func (c *Client) onChannelPinsUpdate(ctx *EventChannelPinsUpdate) error {
	return nil
}

func (c *Client) onChannelUpdate(ctx *EventChannelUpdate) error {

	c.ChannelCache.Set(ctx.ID, ctx.Channel, cache.DefaultExpiration)

	return nil
}

func (c *Client) onGuildBanAdd(ctx *EventGuildBanAdd) error {
	return nil
}

func (c *Client) onGuildBanRemove(ctx *EventGuildBanRemove) error {
	return nil
}

func (c *Client) onGuildCreate(ctx *EventGuildCreate) error {

	c.GuildCache.Set(ctx.ID, ctx.Guild, cache.DefaultExpiration)

	return nil
}

func (c *Client) onGuildDelete(ctx *EventGuildDelete) error {

	guild, cached := c.GuildCache.Get(ctx.UnavailableGuild.ID)
	if cached {
		guild := guild.(Guild)
		guild.Unavailable = true
		c.GuildCache.Set(guild.ID, guild, cache.DefaultExpiration)
	}

	return nil
}

func (c *Client) onGuildEmojisUpdate(ctx *EventGuildEmojisUpdate) error {
	return nil
}

func (c *Client) onGuildIntegrationsUpdate(ctx *EventGuildIntegrationsUpdate) error {
	return nil
}

func (c *Client) onGuildMemberAdd(ctx *EventGuildMemberAdd) error {
	return nil
}

func (c *Client) onGuildMemberRemove(ctx *EventGuildMemberRemove) error {
	return nil
}

func (c *Client) onGuildMemberUpdate(ctx *EventGuildMemberUpdate) error {
	return nil
}

func (c *Client) onGuildRoleCreate(ctx *EventGuildRoleCreate) error {
	return nil
}

func (c *Client) onGuildRoleDelete(ctx *EventGuildRoleDelete) error {
	return nil
}

func (c *Client) onGuildRoleUpdate(ctx *EventGuildRoleUpdate) error {
	return nil
}

func (c *Client) onGuildUpdate(ctx *EventGuildUpdate) error {

	c.GuildCache.Set(ctx.ID, ctx.Guild, cache.DefaultExpiration)

	return nil
}

func (c *Client) onHello(ctx *EventHello) error {
	return nil
}

func (c *Client) onInviteCreate(ctx *EventInviteCreate) error {
	return nil
}

func (c *Client) onInviteDelete(ctx *EventInviteDelete) error {
	return nil
}

func (c *Client) onMessageCreate(ctx *EventMessageCreate) error {

	var err error

	if ctx.Author.ID != c.User.ID {
		c.Log.WithFields(logrus.Fields{
			"author":   ctx.Author.Username,
			"message":  ctx.Content,
			"serverid": ctx.GuildID,
		}).Debug("Websocket: received message")
	}

	if strings.HasPrefix(ctx.Content, "!") {
		noPrefix := ctx.Content[1:]
		command := strings.Fields(noPrefix)[0]
		args := strings.Fields(noPrefix)[1:]
		switch command {
		case "clear":
			err = c.extClearOnMessage(ctx, args)
		default:
			err = nil
		}
		if err != nil {
			return fmt.Errorf("Couldn't execute command: %s", err)
		}
	}

	return nil
}

func (c *Client) onMessageDelete(ctx *EventMessageDelete) error {
	return nil
}

func (c *Client) onMessageDeleteBulk(ctx *EventMessageDeleteBulk) error {
	return nil
}

func (c *Client) onMessageReactionAdd(ctx *EventMessageReactionAdd) error {
	return nil
}

func (c *Client) onMessageReactionRemove(ctx *EventMessageReactionRemove) error {
	return nil
}

func (c *Client) onMessageReactionRemoveAll(ctx *EventMessageReactionRemoveAll) error {
	return nil
}

func (c *Client) onMessageReactionRemoveEmoji(ctx *EventMessageReactionRemoveEmoji) error {
	return nil
}

func (c *Client) onMessageUpdate(ctx *EventMessageUpdate) error {
	return nil
}

func (c *Client) onPresenceUpdate(ctx *EventPresenceUpdate) error {
	return nil
}

func (c *Client) onReady(ctx *EventReady) error {
	c.User = ctx.User
	c.SessionID = ctx.SessionID
	if c.Sharded {
		err := CCReady(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) onReconnect(ctx *EventReconnect) error {
	return nil
}

func (c *Client) onResumed(ctx *EventResumed) error {
	return nil
}

func (c *Client) onTypingStart(ctx *EventTypingStart) error {
	return nil
}

func (c *Client) onUserUpdate(ctx *EventUserUpdate) error {

	c.User = ctx.User

	return nil
}

func (c *Client) onVoiceServerUpdate(ctx *EventVoiceServerUpdate) error {
	return nil
}

func (c *Client) onVoiceStateUpdate(ctx *EventVoiceStateUpdate) error {
	return nil
}

func (c *Client) onWebhooksUpdate(ctx *EventWebhooksUpdate) error {
	return nil
}

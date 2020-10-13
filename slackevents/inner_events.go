// inner_events.go provides EventsAPI particular inner events

package slackevents

import (
	"encoding/json"

	"github.com/slack-go/slack"
)

// EventsAPIInnerEvent the inner event of a EventsAPI event_callback Event.
type EventsAPIInnerEvent struct {
	Type string `json:"type"`
	Data interface{}
}

// AppMentionEvent is an (inner) EventsAPI subscribable event.
type AppMentionEvent struct {
	Type            string      `json:"type"`
	User            string      `json:"user"`
	Text            string      `json:"text"`
	TimeStamp       string      `json:"ts"`
	ThreadTimeStamp string      `json:"thread_ts"`
	Channel         string      `json:"channel"`
	EventTimeStamp  json.Number `json:"event_ts"`

	// When Message comes from a channel that is shared between workspaces
	UserTeam   string `json:"user_team,omitempty"`
	SourceTeam string `json:"source_team,omitempty"`

	// BotID is filled out when a bot triggers the app_mention event
	BotID string `json:"bot_id,omitempty"`
}

// AppHomeOpenedEvent Your Slack app home was opened.
type AppHomeOpenedEvent struct {
	Type           string      `json:"type"`
	User           string      `json:"user"`
	Channel        string      `json:"channel"`
	EventTimeStamp json.Number `json:"event_ts"`
	Tab            string      `json:"tab"`
	View           slack.View  `json:"view"`
}

// AppUninstalledEvent Your Slack app was uninstalled.
type AppUninstalledEvent struct {
	Type string `json:"type"`
}

// GridMigrationFinishedEvent An enterprise grid migration has finished on this workspace.
type GridMigrationFinishedEvent struct {
	Type         string `json:"type"`
	EnterpriseID string `json:"enterprise_id"`
}

// GridMigrationStartedEvent An enterprise grid migration has started on this workspace.
type GridMigrationStartedEvent struct {
	Type         string `json:"type"`
	EnterpriseID string `json:"enterprise_id"`
}

// LinkSharedEvent A message was posted containing one or more links relevant to your application
type LinkSharedEvent struct {
	Type             string        `json:"type"`
	User             string        `json:"user"`
	TimeStamp        string        `json:"ts"`
	Channel          string        `json:"channel"`
	MessageTimeStamp json.Number   `json:"message_ts"`
	ThreadTimeStamp  string        `json:"thread_ts"`
	Links            []sharedLinks `json:"links"`
}

type sharedLinks struct {
	Domain string `json:"domain"`
	URL    string `json:"url"`
}

// MessageEvent occurs when a variety of types of messages has been posted.
// Parse ChannelType to see which
// if ChannelType = "group", this is a private channel message
// if ChannelType = "channel", this message was sent to a channel
// if ChannelType = "im", this is a private message
// if ChannelType = "mim", A message was posted in a multiparty direct message channel
// TODO: Improve this so that it is not required to manually parse ChannelType
type MessageEvent struct {
	// Basic Message Event - https://api.slack.com/events/message
	ClientMsgID     string      `json:"client_msg_id"`
	Type            string      `json:"type"`
	User            string      `json:"user"`
	Text            string      `json:"text"`
	ThreadTimeStamp string      `json:"thread_ts"`
	TimeStamp       string      `json:"ts"`
	Channel         string      `json:"channel"`
	ChannelType     string      `json:"channel_type"`
	EventTimeStamp  json.Number `json:"event_ts"`

	// When Message comes from a channel that is shared between workspaces
	UserTeam   string `json:"user_team,omitempty"`
	SourceTeam string `json:"source_team,omitempty"`

	// Edited Message
	Message         *MessageEvent `json:"message,omitempty"`
	PreviousMessage *MessageEvent `json:"previous_message,omitempty"`
	Edited          *Edited       `json:"edited,omitempty"`

	// Message Subtypes
	SubType string `json:"subtype,omitempty"`

	// bot_message (https://api.slack.com/events/message/bot_message)
	BotID    string `json:"bot_id,omitempty"`
	Username string `json:"username,omitempty"`
	Icons    *Icon  `json:"icons,omitempty"`

	Upload bool   `json:"upload"`
	Files  []File `json:"files"`

	Attachments []slack.Attachment `json:"attachments,omitempty"`

	// Root is the message that was broadcast to the channel when the SubType is
	// thread_broadcast. If this is not a thread_broadcast message event, this
	// value is nil.
	Root *MessageEvent `json:"root"`
}

// MemberJoinedChannelEvent A member joined a public or private channel
type MemberJoinedChannelEvent struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channel_type"`
	Team        string `json:"team"`
	Inviter     string `json:"inviter"`
}

// MemberLeftChannelEvent A member left a public or private channel
type MemberLeftChannelEvent struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channel_type"`
	Team        string `json:"team"`
}

// Both PinAddedEvent and PinRemovedEvent have this same underlying structure,
// but this is not an Event Slack will ever send you itself
type PinEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Item           Item   `json:"item"`
	Channel        string `json:"channel_id"`
	EventTimestamp string `json:"event_ts"`
	HasPins        bool   `json:"has_pins,omitempty"`
}

// Both ReactionAddedEvent and ReactionRemovedEvent have this same underlying structure,
// but this is not an Event Slack will ever send you itself
type ReactionEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Reaction       string `json:"reaction"`
	ItemUser       string `json:"item_user"`
	Item           Item   `json:"item"`
	EventTimestamp string `json:"event_ts"`
}

// ReactionAddedEvent An reaction was added to a message - https://api.slack.com/events/reaction_added
type ReactionAddedEvent ReactionEvent

// ReactionRemovedEvent An reaction was removed from a message - https://api.slack.com/events/reaction_removed
type ReactionRemovedEvent ReactionEvent

// PinAddedEvent An item was pinned to a channel - https://api.slack.com/events/pin_added
type PinAddedEvent PinEvent

// PinRemovedEvent An item was unpinned from a channel - https://api.slack.com/events/pin_removed
type PinRemovedEvent PinEvent

type tokens struct {
	Oauth []string `json:"oauth"`
	Bot   []string `json:"bot"`
}

// TokensRevokedEvent APP's API tokes are revoked - https://api.slack.com/events/tokens_revoked
type TokensRevokedEvent struct {
	Type   string `json:"type"`
	Tokens tokens `json:"tokens"`
}

// JSONTime exists so that we can have a String method converting the date
type JSONTime int64

// Comment contains all the information relative to a comment
type Comment struct {
	ID        string   `json:"id,omitempty"`
	Created   JSONTime `json:"created,omitempty"`
	Timestamp JSONTime `json:"timestamp,omitempty"`
	User      string   `json:"user,omitempty"`
	Comment   string   `json:"comment,omitempty"`
}

// File is a file upload
type File struct {
	ID                 string `json:"id"`
	Created            int    `json:"created"`
	Timestamp          int    `json:"timestamp"`
	Name               string `json:"name"`
	Title              string `json:"title"`
	Mimetype           string `json:"mimetype"`
	Filetype           string `json:"filetype"`
	PrettyType         string `json:"pretty_type"`
	User               string `json:"user"`
	Editable           bool   `json:"editable"`
	Size               int    `json:"size"`
	Mode               string `json:"mode"`
	IsExternal         bool   `json:"is_external"`
	ExternalType       string `json:"external_type"`
	IsPublic           bool   `json:"is_public"`
	PublicURLShared    bool   `json:"public_url_shared"`
	DisplayAsBot       bool   `json:"display_as_bot"`
	Username           string `json:"username"`
	URLPrivate         string `json:"url_private"`
	URLPrivateDownload string `json:"url_private_download"`
	Thumb64            string `json:"thumb_64"`
	Thumb80            string `json:"thumb_80"`
	Thumb360           string `json:"thumb_360"`
	Thumb360W          int    `json:"thumb_360_w"`
	Thumb360H          int    `json:"thumb_360_h"`
	Thumb480           string `json:"thumb_480"`
	Thumb480W          int    `json:"thumb_480_w"`
	Thumb480H          int    `json:"thumb_480_h"`
	Thumb160           string `json:"thumb_160"`
	Thumb720           string `json:"thumb_720"`
	Thumb720W          int    `json:"thumb_720_w"`
	Thumb720H          int    `json:"thumb_720_h"`
	Thumb800           string `json:"thumb_800"`
	Thumb800W          int    `json:"thumb_800_w"`
	Thumb800H          int    `json:"thumb_800_h"`
	Thumb960           string `json:"thumb_960"`
	Thumb960W          int    `json:"thumb_960_w"`
	Thumb960H          int    `json:"thumb_960_h"`
	Thumb1024          string `json:"thumb_1024"`
	Thumb1024W         int    `json:"thumb_1024_w"`
	Thumb1024H         int    `json:"thumb_1024_h"`
	ImageExifRotation  int    `json:"image_exif_rotation"`
	OriginalW          int    `json:"original_w"`
	OriginalH          int    `json:"original_h"`
	Permalink          string `json:"permalink"`
	PermalinkPublic    string `json:"permalink_public"`
}

// Edited is included when a Message is edited
type Edited struct {
	User      string `json:"user"`
	TimeStamp string `json:"ts"`
}

// Icon is used for bot messages
type Icon struct {
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

// Item is any type of slack message - message, file, or file comment.
type Item struct {
	Type      string       `json:"type"`
	Channel   string       `json:"channel,omitempty"`
	Message   *ItemMessage `json:"message,omitempty"`
	File      *File        `json:"file,omitempty"`
	Comment   *Comment     `json:"comment,omitempty"`
	Timestamp string       `json:"ts,omitempty"`
}

// ItemMessage is the event message
type ItemMessage struct {
	Type            string   `json:"type"`
	User            string   `json:"user"`
	Text            string   `json:"text"`
	Timestamp       string   `json:"ts"`
	PinnedTo        []string `json:"pinned_to"`
	ReplaceOriginal bool     `json:"replace_original"`
	DeleteOriginal  bool     `json:"delete_original"`
}

// IsEdited checks if the MessageEvent is caused by an edit
func (e MessageEvent) IsEdited() bool {
	return e.Message != nil &&
		e.Message.Edited != nil
}

const (
	// AppMention is an Events API subscribable event
	AppMention = "app_mention"
	// AppHomeOpened Your Slack app home was opened
	AppHomeOpened = "app_home_opened"
	// AppUninstalled Your Slack app was uninstalled.
	AppUninstalled = "app_uninstalled"
	// GridMigrationFinished An enterprise grid migration has finished on this workspace.
	GridMigrationFinished = "grid_migration_finished"
	// GridMigrationStarted An enterprise grid migration has started on this workspace.
	GridMigrationStarted = "grid_migration_started"
	// LinkShared A message was posted containing one or more links relevant to your application
	LinkShared = "link_shared"
	// Message A message was posted to a channel, private channel (group), im, or mim
	Message = "message"
	// Member Joined Channel
	MemberJoinedChannel = "member_joined_channel"
	// Member Left Channel
	MemberLeftChannel = "member_left_channel"
	// PinAdded An item was pinned to a channel
	PinAdded = "pin_added"
	// PinRemoved An item was unpinned from a channel
	PinRemoved = "pin_removed"
	// ReactionAdded An reaction was added to a message
	ReactionAdded = "reaction_added"
	// ReactionRemoved An reaction was removed from a message
	ReactionRemoved = "reaction_removed"
	// TokensRevoked APP's API tokes are revoked
	TokensRevoked = "tokens_revoked"
)

// EventsAPIInnerEventMapping maps INNER Event API events to their corresponding struct
// implementations. The structs should be instances of the unmarshalling
// target for the matching event type.
var EventsAPIInnerEventMapping = map[string]interface{}{
	AppMention:            AppMentionEvent{},
	AppHomeOpened:         AppHomeOpenedEvent{},
	AppUninstalled:        AppUninstalledEvent{},
	GridMigrationFinished: GridMigrationFinishedEvent{},
	GridMigrationStarted:  GridMigrationStartedEvent{},
	LinkShared:            LinkSharedEvent{},
	Message:               MessageEvent{},
	MemberJoinedChannel:   MemberJoinedChannelEvent{},
	MemberLeftChannel:     MemberLeftChannelEvent{},
	PinAdded:              PinAddedEvent{},
	PinRemoved:            PinRemovedEvent{},
	ReactionAdded:         ReactionAddedEvent{},
	ReactionRemoved:       ReactionRemovedEvent{},
	TokensRevoked:         TokensRevokedEvent{},
}

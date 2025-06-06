// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.notification.getUnreadCount

import (
	"context"

	"github.com/bluesky-social/indigo/lex/util"
)

// NotificationGetUnreadCount_Output is the output of a app.bsky.notification.getUnreadCount call.
type NotificationGetUnreadCount_Output struct {
	Count int64 `json:"count" cborgen:"count"`
}

// NotificationGetUnreadCount calls the XRPC method "app.bsky.notification.getUnreadCount".
func NotificationGetUnreadCount(ctx context.Context, c util.LexClient, priority bool, seenAt string) (*NotificationGetUnreadCount_Output, error) {
	var out NotificationGetUnreadCount_Output

	params := map[string]interface{}{}
	if priority {
		params["priority"] = priority
	}
	if seenAt != "" {
		params["seenAt"] = seenAt
	}
	if err := c.LexDo(ctx, util.Query, "", "app.bsky.notification.getUnreadCount", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

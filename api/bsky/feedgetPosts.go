// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.getPosts

import (
	"context"

	"github.com/bluesky-social/indigo/lex/util"
)

// FeedGetPosts_Output is the output of a app.bsky.feed.getPosts call.
type FeedGetPosts_Output struct {
	Posts []*FeedDefs_PostView `json:"posts" cborgen:"posts"`
}

// FeedGetPosts calls the XRPC method "app.bsky.feed.getPosts".
//
// uris: List of post AT-URIs to return hydrated views for.
func FeedGetPosts(ctx context.Context, c util.LexClient, uris []string) (*FeedGetPosts_Output, error) {
	var out FeedGetPosts_Output

	params := map[string]interface{}{}
	params["uris"] = uris
	if err := c.LexDo(ctx, util.Query, "", "app.bsky.feed.getPosts", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.actor.searchActorsTypeahead

import (
	"context"

	"github.com/bluesky-social/indigo/lex/util"
)

// ActorSearchActorsTypeahead_Output is the output of a app.bsky.actor.searchActorsTypeahead call.
type ActorSearchActorsTypeahead_Output struct {
	Actors []*ActorDefs_ProfileViewBasic `json:"actors" cborgen:"actors"`
}

// ActorSearchActorsTypeahead calls the XRPC method "app.bsky.actor.searchActorsTypeahead".
//
// q: Search query prefix; not a full query string.
// term: DEPRECATED: use 'q' instead.
func ActorSearchActorsTypeahead(ctx context.Context, c util.LexClient, limit int64, q string, term string) (*ActorSearchActorsTypeahead_Output, error) {
	var out ActorSearchActorsTypeahead_Output

	params := map[string]interface{}{}
	if limit != 0 {
		params["limit"] = limit
	}
	if q != "" {
		params["q"] = q
	}
	if term != "" {
		params["term"] = term
	}
	if err := c.LexDo(ctx, util.Query, "", "app.bsky.actor.searchActorsTypeahead", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

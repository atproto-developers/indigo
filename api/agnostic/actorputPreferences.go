// Copied from indigo:api/bsky/actorputPreferences.go

package agnostic

// schema: app.bsky.actor.putPreferences

import (
	"context"

	"github.com/bluesky-social/indigo/lex/util"
)

// ActorPutPreferences_Input is the input argument to a app.bsky.actor.putPreferences call.
type ActorPutPreferences_Input struct {
	Preferences []map[string]any `json:"preferences" cborgen:"preferences"`
}

// ActorPutPreferences calls the XRPC method "app.bsky.actor.putPreferences".
func ActorPutPreferences(ctx context.Context, c util.LexClient, input *ActorPutPreferences_Input) error {
	if err := c.LexDo(ctx, util.Procedure, "application/json", "app.bsky.actor.putPreferences", nil, input, nil); err != nil {
		return err
	}

	return nil
}

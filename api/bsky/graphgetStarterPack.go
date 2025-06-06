// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.getStarterPack

import (
	"context"

	"github.com/bluesky-social/indigo/lex/util"
)

// GraphGetStarterPack_Output is the output of a app.bsky.graph.getStarterPack call.
type GraphGetStarterPack_Output struct {
	StarterPack *GraphDefs_StarterPackView `json:"starterPack" cborgen:"starterPack"`
}

// GraphGetStarterPack calls the XRPC method "app.bsky.graph.getStarterPack".
//
// starterPack: Reference (AT-URI) of the starter pack record.
func GraphGetStarterPack(ctx context.Context, c util.LexClient, starterPack string) (*GraphGetStarterPack_Output, error) {
	var out GraphGetStarterPack_Output

	params := map[string]interface{}{}
	params["starterPack"] = starterPack
	if err := c.LexDo(ctx, util.Query, "", "app.bsky.graph.getStarterPack", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

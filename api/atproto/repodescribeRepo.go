// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.repo.describeRepo

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// RepoDescribeRepo_Output is the output of a com.atproto.repo.describeRepo call.
type RepoDescribeRepo_Output struct {
	Collections     []string    `json:"collections" cborgen:"collections"`
	Did             string      `json:"did" cborgen:"did"`
	DidDoc          interface{} `json:"didDoc" cborgen:"didDoc"`
	Handle          string      `json:"handle" cborgen:"handle"`
	HandleIsCorrect bool        `json:"handleIsCorrect" cborgen:"handleIsCorrect"`
}

// RepoDescribeRepo calls the XRPC method "com.atproto.repo.describeRepo".
//
// repo: The handle or DID of the repo.
func RepoDescribeRepo(ctx context.Context, c *xrpc.Client, repo string) (*RepoDescribeRepo_Output, error) {
	var out RepoDescribeRepo_Output

	params := map[string]interface{}{
		"repo": repo,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.repo.describeRepo", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

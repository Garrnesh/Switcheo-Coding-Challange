package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "CC/testutil/keeper"
	"CC/x/cc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CcKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

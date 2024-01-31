package cc_test

import (
	"testing"

	keepertest "CC/testutil/keeper"
	"CC/testutil/nullify"
	"CC/x/cc/module"
	"CC/x/cc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CcKeeper(t)
	cc.InitGenesis(ctx, k, genesisState)
	got := cc.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

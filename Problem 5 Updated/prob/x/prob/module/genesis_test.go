package prob_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "prob/testutil/keeper"
	"prob/testutil/nullify"
	"prob/x/prob/module"
	"prob/x/prob/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProbKeeper(t)
	prob.InitGenesis(ctx, k, genesisState)
	got := prob.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

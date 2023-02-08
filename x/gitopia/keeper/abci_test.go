package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"
	"github.com/gitopia/gitopia/app/params"
	"github.com/gitopia/gitopia/testutil/simapp"
	gitopiatypes "github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/stretchr/testify/assert"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
)


func TestTokenDistributionSucessWithNoDistributionProportion(t *testing.T){
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "gitopia-1", Time: time.Now().UTC()})
	gParams := gitopiatypes.Params{
		DistributionProportions: []gitopiatypes.DistributionProportion{},
	}
	minter := gitopiatypes.MinterAccountName
	feeCollector :=	authtypes.FeeCollectorName
	gitopiaKeeper := app.GitopiaKeeper
	bankKeeper := app.BankKeeper
	accountKeeper := app.AccountKeeper

	// setup
	err := testutil.FundModuleAccount(bankKeeper, ctx, minter, sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50))))
	assert.NoError(t, err)
	gitopiaKeeper.SetParams(ctx, gParams)

	gitopiaKeeper.BeginBlocker(ctx)
	
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50)), bankKeeper.GetBalance(ctx, accountKeeper.GetModuleAddress(feeCollector), "utlore"))
} 


package horizon

import (
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/services/horizon/internal/actions"
	"github.com/stellar/go/services/horizon/internal/ledger"
	"github.com/stellar/go/services/horizon/internal/resourceadapter"
	"github.com/stellar/go/support/render/hal"
)

// Interface verification
var _ actions.JSONer = (*RootAction)(nil)

// RootAction provides a summary of the horizon instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() error {
	var res horizon.Root
	templates := map[string]string{
		"accounts":           actions.AccountsQuery{}.URITemplate(),
		"offers":             actions.OffersQuery{}.URITemplate(),
		"strictReceivePaths": StrictReceivePathsQuery{}.URITemplate(),
		"strictSendPaths":    FindFixedPathsQuery{}.URITemplate(),
	}
	coreInfo := action.App.coreSettings.get()
	resourceadapter.PopulateRoot(
		action.R.Context(),
		&res,
		ledger.CurrentState(),
		action.App.horizonVersion,
		coreInfo.coreVersion,
		action.App.config.NetworkPassphrase,
		coreInfo.currentProtocolVersion,
		coreInfo.coreSupportedProtocolVersion,
		action.App.config.FriendbotURL,
		templates,
	)

	hal.Render(action.W, res)
	return action.Err
}

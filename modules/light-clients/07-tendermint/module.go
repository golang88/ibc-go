package tendermint

import (
	"github.com/golang88/ibc-go/modules/light-clients/07-tendermint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}

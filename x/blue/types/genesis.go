package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PrefPoolList: []PrefPool{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in prefPool
	prefPoolIdMap := make(map[uint64]bool)
	prefPoolCount := gs.GetPrefPoolCount()
	for _, elem := range gs.PrefPoolList {
		if _, ok := prefPoolIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for prefPool")
		}
		if elem.Id >= prefPoolCount {
			return fmt.Errorf("prefPool id should be lower or equal than the last id")
		}
		prefPoolIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

package contract

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/baking-bad/bcdhub/internal/bcd"
	astContract "github.com/baking-bad/bcdhub/internal/bcd/contract"
	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/parsers"
	"github.com/pkg/errors"
)

// Jakarta -
type Jakarta struct {
	ctx *config.Context
}

// NewJakarta -
func NewJakarta(ctx *config.Context) *Jakarta {
	return &Jakarta{ctx: ctx}
}

// Parse -
func (p *Jakarta) Parse(operation *operation.Operation, store parsers.Store) error {
	if !operation.IsOrigination() {
		return errors.Errorf("invalid operation kind in computeContractMetrics: %s", operation.Kind)
	}

	contract := contract.Contract{
		Level:      operation.Level,
		Timestamp:  operation.Timestamp,
		Manager:    operation.Source,
		Account:    operation.Destination,
		Delegate:   operation.Delegate,
		LastAction: operation.Timestamp,
	}

	if err := p.computeMetrics(operation, &contract); err != nil {
		return err
	}

	store.AddContracts(&contract)
	return nil
}

func (p *Jakarta) computeMetrics(operation *operation.Operation, c *contract.Contract) error {
	script, err := astContract.NewParser(operation.Script)
	if err != nil {
		return errors.Wrap(err, "astContract.NewParser")
	}
	operation.AST = script.Code

	contractScript, err := p.ctx.Scripts.ByHash(script.Hash)
	if err != nil {
		if !p.ctx.Storage.IsRecordNotFound(err) {
			return err
		}
		var s bcd.RawScript
		if err := json.Unmarshal(script.CodeRaw, &s); err != nil {
			return err
		}

		constants, err := script.FindConstants()
		if err != nil {
			return errors.Wrap(err, "script.FindConstants")
		}

		if len(constants) > 0 {
			globalConstants, err := p.ctx.GlobalConstants.All(constants...)
			if err != nil {
				return err
			}
			contractScript.Constants = globalConstants
			replaceConstants(&contractScript, operation)

			script, err = astContract.NewParser(operation.Script)
			if err != nil {
				return errors.Wrap(err, "astContract.NewParser")
			}
			operation.AST = script.Code
		}

		if err := script.Parse(); err != nil {
			return err
		}

		params, err := script.Code.Parameter.ToTypedAST()
		if err != nil {
			return err
		}

		operation.Script = script.CodeRaw
		contractScript.Hash = script.Hash
		contractScript.Code = s.Code
		contractScript.Parameter = s.Parameter
		contractScript.Storage = s.Storage
		contractScript.Views = s.Views
		contractScript.FailStrings = script.FailStrings.Values()
		contractScript.Annotations = script.Annotations.Values()
		contractScript.Tags = types.NewTags(script.Tags.Values())
		contractScript.Hardcoded = script.HardcodedAddresses.Values()
		contractScript.Entrypoints = params.GetEntrypoints()

		c.Jakarta = contractScript
	} else {
		c.JakartaID = contractScript.ID
		c.Jakarta = contractScript
	}

	c.Tags = contractScript.Tags

	return nil
}

func replaceConstants(c *contract.Script, operation *operation.Operation) {
	pattern := `{"prim":"constant","args":[{"string":"%s"}]}`
	for i := range c.Constants {
		operation.Script = bytes.ReplaceAll(
			operation.Script,
			[]byte(fmt.Sprintf(pattern, c.Constants[i].Address)),
			c.Constants[i].Value,
		)
	}
}
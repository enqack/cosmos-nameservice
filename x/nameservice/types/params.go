package types

import (
	"errors"
	"fmt"

	yaml "gopkg.in/yaml.v2"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultCreateWhoisPrice string = "10trycoin"
	DefaultUpdateWhoisPrice string = "5trycoin"
	DefaultDeleteWhoisPrice string = "1trycoin"
)

// Parameter keys
var (
	KeyCreateWhoisPrice = []byte("CreateWhoisPrice")
	KeyUpdateWhoisPrice = []byte("UpdateWhoisPrice")
	KeyDeleteWhoisPrice = []byte("DeleteWhoisPrice")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Params return all of the whois params
type Params struct {
	CreateWhoisPrice string `json:"minimum_create_whois_price" yaml:"minimum_create_whois_price"`
	UpdateWhoisPrice string `json:"update_whois_price" yaml:"update_whois_price"`
	DeleteWhoisPrice string `json:"delete_whois_price" yaml:"delete_whois_price"`
}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(createWhoisPrice string, updateWhoisPrice string, deleteWhoisPrice string) Params {
	return Params{
		CreateWhoisPrice: createWhoisPrice,
		UpdateWhoisPrice: updateWhoisPrice,
		DeleteWhoisPrice: deleteWhoisPrice,
	}
}

// DefaultParams returns default whois parameters
func DefaultParams() Params {
	return NewParams(
		DefaultCreateWhoisPrice,
		DefaultUpdateWhoisPrice,
		DefaultDeleteWhoisPrice,
	)
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCreateWhoisPrice, &p.CreateWhoisPrice, validateCreateWhoisPrice),
		paramtypes.NewParamSetPair(KeyUpdateWhoisPrice, &p.UpdateWhoisPrice, validateUpdateWhoisPrice),
		paramtypes.NewParamSetPair(KeyDeleteWhoisPrice, &p.DeleteWhoisPrice, validateDeleteWhoisPrice),
	}
}

// validate a set of params
func (p Params) Validate() error {
	if err := validateCreateWhoisPrice(p.CreateWhoisPrice); err != nil {
		return err
	}

	if err := validateUpdateWhoisPrice(p.UpdateWhoisPrice); err != nil {
		return err
	}

	if err := validateDeleteWhoisPrice(p.DeleteWhoisPrice); err != nil {
		return err
	}

	return nil
}

func validateCreateWhoisPrice(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return errors.New("create whois price cannot be empty")
	}

	return nil
}

func validateUpdateWhoisPrice(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return errors.New("update whois price cannot be empty")
	}

	return nil
}

func validateDeleteWhoisPrice(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return errors.New("delete whois price cannot be empty")
	}

	return nil
}

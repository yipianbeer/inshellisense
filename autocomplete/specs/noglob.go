// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["noglob"] = model.Subcommand{
		Name:        []string{"noglob"},
		Description: `ZSH pre-command modifier that disables glob expansion`,
		Args: []model.Arg{{
			IsCommand: true,
		}},
	}
}
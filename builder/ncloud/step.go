// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ncloud

import (
	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func processStepResult(err error, sayError func(error), state multistep.StateBag) multistep.StepAction {
	if err != nil {
		state.Put("error", err)
		sayError(err)

		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

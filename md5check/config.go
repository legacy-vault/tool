//============================================================================//
//
// Copyright © 2018 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2018-12-01.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// config.go.

// Application's Configuration.

package main

import (
	"errors"
	"fmt"
	"os"
)

// Command Line Arguments' Parameters.
const CLACountMax = 5
const CLAOptionPrefix = '-' // Options Marker.
const OptionHelp1 = "-h"
const OptionHelp2 = "--help"
const OptionIgnore1 = "-i"
const OptionIgnore2 = "--ignore"
const OptionShowProgress1 = "-p"
const OptionShowProgress2 = "--show-progress"
const OptionVerboseMode1 = "-v"
const OptionVerboseMode2 = "--verbose"

// Error Messages and Formats.
const ErrInputFilePathIsNotSet = "Input File Path is not set"
const ErrCLACount = "Command Line Arguments Count Error"
const ErrCLALength = "Command Line Argument Length Error"
const ErrFormatCLAUnknownOption = "Unknown Command Line Option '%s'"
const ErrCLAFilePathCount = "File Path is set multiple Times"

type Configuration struct {
	InputFilePath           string
	HelpMode                bool
	IgnoreNotAcessibleFiles bool
	ShowProgress            bool
	VerboseMode             bool
}

// Cached Errors.
var errInputFilePathIsNotSet = errors.New(ErrInputFilePathIsNotSet)
var errCLACount = errors.New(ErrCLACount)
var errCLALength = errors.New(ErrCLALength)
var errCLAFilePathCount = errors.New(ErrCLAFilePathCount)

// Configures the Application.
func configureApp(cfg *Configuration) error {

	var claCount int
	var err error
	var inputFilePathIsSet bool = false
	var osArgsCount int

	// Clear the Configuration.
	cfg.HelpMode = false
	cfg.IgnoreNotAcessibleFiles = false
	cfg.VerboseMode = false

	// Get Command Line Arguments...

	// 1. Check Arguments' Count.
	osArgsCount = len(os.Args)
	claCount = osArgsCount - 1
	if claCount > CLACountMax {
		return errCLACount
	}
	if claCount == 0 {
		cfg.HelpMode = true
		cfg.VerboseMode = true
		return nil
	}

	// 2. Fill Configuration Object.
	for i := 1; i < osArgsCount; i++ {
		arg := os.Args[i]
		switch arg {

		case OptionHelp1:
			cfg.HelpMode = true

		case OptionHelp2:
			cfg.HelpMode = true

		case OptionIgnore1:
			cfg.IgnoreNotAcessibleFiles = true

		case OptionIgnore2:
			cfg.IgnoreNotAcessibleFiles = true

		case OptionShowProgress1:
			cfg.ShowProgress = true

		case OptionShowProgress2:
			cfg.ShowProgress = true

		case OptionVerboseMode1:
			cfg.VerboseMode = true

		case OptionVerboseMode2:
			cfg.VerboseMode = true

		default:
			// Argument is not a known Option.
			// Ensure that it is really not an Option.
			if len(arg) < 1 {
				return errCLALength
			}
			if arg[0] == CLAOptionPrefix {
				err = fmt.Errorf(ErrFormatCLAUnknownOption, arg)
				return err
			}

			// Normal Argument.
			if inputFilePathIsSet {
				return errCLAFilePathCount
			}
			cfg.InputFilePath = arg
			inputFilePathIsSet = true
		}
	}

	// 3. Check that required Parameters have been configured.
	if !inputFilePathIsSet {
		return errInputFilePathIsNotSet
	}

	return nil
}

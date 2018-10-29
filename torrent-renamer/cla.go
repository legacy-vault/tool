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
// Creation Date:	2018-10-29.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// cla.go.

// Command Line Arguments Handling.

// Last Update Time: 2018-10-30.

package main

import (
	"flag"
	"fmt"
)

// Names of Command Line Arguments (Keys).
const CLANameInputFolder = "in"
const CLANameOutputFolder = "out"
const CLANameSubLevels = "sl"
const CLANameUpperCaseOut = "uc"

// Default values of Command Line Arguments.
const CLADefaultValueInputFolder = ""
const CLADefaultValueOutputFolder = ""
const CLADefaultValueSubLevels = true
const CLADefaultValueUpperCaseOut = true

// Hint Texts of Command Line Arguments.
const CLAHintInputFolder = "Path to Input Folder"
const CLAHintOutputFolder = "Path to Output Folder"
const CLAHintSubLevels = "Read Sub-Level Directories?"
const CLAHintUpperCaseOut = "Convert Output File Names to Upper Case?"

// Initializes the Command Line Arguments.
func initCLA() error {

	var claInputFolder *string
	var claOutputFolder *string
	var claSubLevels *bool
	var claOutUpperCase *bool

	// Prepare C.L.A. Parameters.
	claInputFolder = flag.String(
		CLANameInputFolder,
		CLADefaultValueInputFolder,
		CLAHintInputFolder,
	)
	claOutputFolder = flag.String(
		CLANameOutputFolder,
		CLADefaultValueOutputFolder,
		CLAHintOutputFolder,
	)
	claSubLevels = flag.Bool(
		CLANameSubLevels,
		CLADefaultValueSubLevels,
		CLAHintSubLevels,
	)
	claOutUpperCase = flag.Bool(
		CLANameUpperCaseOut,
		CLADefaultValueUpperCaseOut,
		CLAHintUpperCaseOut,
	)

	// Read Flags.
	flag.Parse()
	appCfg.FolderPathInput = *claInputFolder
	appCfg.FolderPathOutput = *claOutputFolder
	appCfg.ReadSubLevels = *claSubLevels
	appCfg.OutputFileNameToUpperCase = *claOutUpperCase

	// Check Values.
	if !pathExists(appCfg.FolderPathInput) {
		return fmt.Errorf(
			ErrPathExistence,
			appCfg.FolderPathInput,
		)
	}
	if !pathExists(appCfg.FolderPathOutput) {
		return fmt.Errorf(
			ErrPathExistence,
			appCfg.FolderPathOutput,
		)
	}

	return nil
}

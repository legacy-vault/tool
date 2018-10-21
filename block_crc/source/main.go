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
// Creation Date:	2018-10-21.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// main.go.

// Block CRC :: Main File.

package main

import (
	"fmt"
	"os"
)

// This is a Test Tool to create and verify CRC32 CheckSums of a File.
//
// In 'CheckSum File Creation' Mode, this Tools reads the Data File by Blocks
// of the specified Size and then writes CheckSums to the CheckSums File. For
// each Block of the Data File, two CheckSums are calculated and written:
//
//	1. CheckSum of the current Block;
//	2. CheckSum of the "sliding" ("running") Block.
//
// Sliding Block is a Concatenation of all previously read Blocks (before the
// current Block) into a big single Block. The first sliding CheckSum equals to
// the CheckSum of the first current Block. The last sliding CheckSum equals to
// the CheckSum of the whole Data File. Sliding Block CheckSums help in
// inspecting Error Types, they may be very useful in discovering the Source
// of Errors. Not in 100% Cases, but very often, the Place and Reason of
// Corruption may be learned by the Inspection of Error List.
//
// In 'CheckSum Verification' Mode, this Tools reads the CheckSum File (as well
// as the Data File) and verifies the Integrity of both Data File and CheckSums
// File.
//
// Unfortunately, this Tool can only discover Errors and the Place where they
// occur. This Tool can not recover corrupted Data, while this Operation is very
// Time consuming due to the Nature of Hash Sums.
//
// If the Path to CheckSums File is set to empty, the Tool then copies it from
// the Data File Path Parameter, with the 'bcrc' Extension appended.
//
// If the Input File has Size which is not multiple of the specified Block Size,
// then the Data received from the last File's Part is appended with a Zeroed
// Postfix. To make it simple, if the Block Size is 4 Letters, and we have a
// File with the "ABCDE" Contents, then we calculate Check Sums of two Blocks:
// "ABCD" and "E000".

// Default Configuration Values (before Configuration has been read).
const VerboseModeDefault = false

var verboseMode bool

// Program's Entry Point.
func main() {

	var err error

	// Set Verbose Mode before it has been set by Comman Line Arguments.
	verboseMode = VerboseModeDefault

	// Initialize Application Stages.
	err = stagesInit()
	if err != nil {
		panic(err)
	}

	// Initialize Everything.
	runStage(StageInitialization)
	defer FinDeferred()

	// Read O.S. Command Line Arguments.
	runStage(StageCommandLineArguments)

	// Run Action.
	runStage(actionStage)
}

// Initializations.
func Init() error {

	var err error

	// Self-Check.
	err = selfCheck()
	if err != nil {
		return err
	}

	return nil
}

// Finalizations.
func Fin() error {

	var err error

	// Do Something.
	err = nil // A Plug.
	if err != nil {
		return err
	}

	os.Exit(0)

	return nil
}

// Deferred Finalizations.
func FinDeferred() {

	runStage(StageFinalization)
}

// Does Nothing.
func void() error {

	return nil
}

// Runs the Stage.
func runStage(stage int) {

	var err error
	var f StageFunction
	var msg string

	// Check Stage Index.
	if (stage < StageIndexMin) || (stage > StageIndexMax) {
		stage = StageUnknown
	}

	// Get Stage Function.
	f = stageFucntions[stage]

	// Run Stage Function.
	if verboseMode {
		//msg = stageStatusText(stage, StatusStarted)
		//fmt.Println(msg)
	}
	err = f()
	if err != nil {
		if verboseMode {
			msg = stageStatusText(stage, StatusFailed)
			fmt.Println(msg)
		}
		panic(err)
	}
	if verboseMode {
		msg = stageStatusText(stage, StatusFinished)
		fmt.Println(msg)
	}
}

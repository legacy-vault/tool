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

// stage.go.

// Block CRC :: Application Stages.

package main

// Known Stages List.
const StagesCount = 5
const StageIndexMin = 1
const StageIndexMax = 5

// Stage Name Indices.
const StageUnknown = 0
const StageInitialization = 1
const StageCommandLineArguments = 2
const StageFinalization = 3
const StageBlockCRCCreation = 4
const StageBlockCRCCheck = 5

// Stage Names.
const StageNameUnknown = "Unknown"                             // 0.
const StageNameInitialization = "Program Initialization"       // 1.
const StageNameCommandLineArguments = "Command Line Arguments" // 2.
const StageNameFinalization = "Program Finalization"           // 3.
const StageNameBlockCRCCreation = "Block CRC Creation"         // 4.
const StageNameBlockCRCCheck = "Block CRC Check"               // 5.

// Known Statuses List.
const StatusesCount = 3
const StatusIndexMin = 1
const StatusIndexMax = 3

// Status Indices.
const StatusUnknown = 0
const StatusStarted = 1
const StatusFinished = 2
const StatusFailed = 3

// Status Names.
//const StatusNameUnknown = "has unknown Status" // 0.
//const StatusNameStarted = "has started"        // 1.
//const StatusNameFinished = "has finished"      // 2.
//const StatusNameFailed = "has failed"          // 3.
const StatusNameUnknown = "???"    // 0.
const StatusNameStarted = "..."    // 1.
const StatusNameFinished = "OK"    // 2.
const StatusNameFailed = "Failure" // 3.

// Delimiters.
const StageDelimiterNamePrefix = "["
const StageDelimiterNamePostfix = "]"
const StageDelimiterStatusPrefix = " "
const StageDelimiterStatusPostfix = "."

type StageFunction func() error

var stageFucntions []StageFunction
var stageNames []string
var stageStatusNames []string

// Initializes Stages Names.
func stagesInit() error {

	// Cache Stage Names.
	stageNames = make([]string, StagesCount+1)

	stageNames[StageUnknown] = StageNameUnknown
	stageNames[StageInitialization] = StageNameInitialization
	stageNames[StageCommandLineArguments] = StageNameCommandLineArguments
	stageNames[StageFinalization] = StageNameFinalization
	stageNames[StageBlockCRCCreation] = StageNameBlockCRCCreation
	stageNames[StageBlockCRCCheck] = StageNameBlockCRCCheck

	// Cache Stage Status Names.
	stageStatusNames = make([]string, StatusesCount+1)

	stageStatusNames[StatusUnknown] = StatusNameUnknown
	stageStatusNames[StatusStarted] = StatusNameStarted
	stageStatusNames[StatusFinished] = StatusNameFinished
	stageStatusNames[StatusFailed] = StatusNameFailed

	// Cache Stage Functions.
	stageFucntions = make([]StageFunction, StagesCount+1)

	stageFucntions[StageUnknown] = void
	stageFucntions[StageInitialization] = Init
	stageFucntions[StageCommandLineArguments] = claRead
	stageFucntions[StageFinalization] = Fin
	stageFucntions[StageBlockCRCCreation] = bCRCCreate
	stageFucntions[StageBlockCRCCheck] = bCRCCheck

	return nil
}

// Returns the Text containing the Stage Name and its Status.
func stageStatusText(stage int, status int) string {

	var s string

	// Check Stage Index.
	if (stage < StageIndexMin) || (stage > StageIndexMax) {
		stage = StageUnknown
	}

	// Check Stage Status.
	if (status < StatusIndexMin) || (status > StatusIndexMax) {
		status = StatusUnknown
	}

	// Create a String.
	s = StageDelimiterNamePrefix +
		stageNames[stage] +
		StageDelimiterNamePostfix +
		StageDelimiterStatusPrefix +
		stageStatusNames[status] +
		StageDelimiterStatusPostfix

	return s
}

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

// result.go.

// Application's Results.

package main

import "fmt"

type Result struct {
	AllClear          bool
	MatchedFilesCount int
	Mismatches        []Mismatch
	SkippedFilesCount int
}

type Mismatch struct {
	FilePath string
}

// Messages and Formats.
const MsgChecksumMismatch = "[Checksum Mismatch]"
const MsgFormatChecksumMatchCount = "Checksum Matches Count: %d.\r\n"
const MsgFormatChecksumMismatchCount = "Checksum Mismatches Count: %d.\r\n"
const MsgFormatSkippedFilesCount = "Skipped Files Count: %d.\r\n"
const MsgAllClear = "All Clear"

// Shows the Results.
func showResults(
	appConfig Configuration,
	appResult Result,
) {

	if appResult.AllClear {
		fmt.Println(MsgAllClear)
		return
	}

	switch appConfig.VerboseMode {

	case true:
		for _, mismatch := range appResult.Mismatches {
			fmt.Println(MsgChecksumMismatch + " " + mismatch.FilePath)
		}
		fmt.Printf(MsgFormatSkippedFilesCount, appResult.SkippedFilesCount)
		fmt.Printf(MsgFormatChecksumMatchCount, appResult.MatchedFilesCount)
		fmt.Printf(MsgFormatChecksumMismatchCount, len(appResult.Mismatches))

	case false:
		fmt.Printf(MsgFormatSkippedFilesCount, appResult.SkippedFilesCount)
		fmt.Printf(MsgFormatChecksumMatchCount, appResult.MatchedFilesCount)
		fmt.Printf(MsgFormatChecksumMismatchCount, len(appResult.Mismatches))
	}

	return
}

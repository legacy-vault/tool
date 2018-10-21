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

// argument.go.

// Block CRC :: O.S. Command Line Arguments Handling.

package main

import (
	"flag"
	"fmt"
)

// Help Hint Descriptions.
const CLAHint_Action = "⚙ Action ('check' or 'create') " +
	"performed with the Block CRC Sum"
const CLAHint_BlockSize = "⚙ Block Size (Bytes) for Check Sum Creation. " +
	"This Block Size is automatically detected when performing a Check, " +
	"as this Block Size is stored inside the CheckSum File created by this" +
	" Tool"
const CLAHint_DataFile = "⚙ Data File Path"
const CLAHint_SumFile = "⚙ Blocked CheckSum File Path " +
	"(if empty or omitted => set to Data File Path plus the '" +
	FileExtensionSeparator + FileOutputExt +
	"' Postfix)"
const CLAHint_VerboseMode = "⚙ Verbose Mode"

// Parameter Names.
const CLAParamName_Action = "action"
const CLAParamName_BlockSize = "block_size"
const CLAParamName_DataFile = "data_file"
const CLAParamName_SumFile = "sum_file"
const CLAParamName_VerboseMode = "verbose"

// Default Values.
const CLADefaultValue_Action = ActionNameCreateBlockCRCSum
const CLADefaultValue_BlockSize = 4096
const CLADefaultValue_DataFile = "input.txt"
const CLADefaultValue_SumFile = ""
const CLADefaultValue_VerboseMode = false

const ReportFormatA = "<%s> = [%v]." + NL

var claAction *string
var claBlockSize *uint64
var claDataFile *string
var claSumFile *string
var claSumFile2 string
var claVerboseMode *bool

// Reads O.S. Command Line Arguments.
func claRead() error {

	var reportFormat string

	// Read Verbose Mode.
	claVerboseMode = flag.Bool(
		CLAParamName_VerboseMode,
		CLADefaultValue_VerboseMode,
		CLAHint_VerboseMode,
	)

	// Read Input File's Name.
	claDataFile = flag.String(
		CLAParamName_DataFile,
		CLADefaultValue_DataFile,
		CLAHint_DataFile,
	)

	// Read Action.
	claAction = flag.String(
		CLAParamName_Action,
		CLADefaultValue_Action,
		CLAHint_Action,
	)

	// Read Block Size.
	claBlockSize = flag.Uint64(
		CLAParamName_BlockSize,
		CLADefaultValue_BlockSize,
		CLAHint_BlockSize,
	)

	// Read Output File's Name.
	claSumFile = flag.String(
		CLAParamName_SumFile,
		CLADefaultValue_SumFile,
		CLAHint_SumFile,
	)

	// Parse Flags.
	flag.Parse()

	// Set Values according to Flag States...

	// 1. Verbose Mode.
	verboseMode = *claVerboseMode

	// 2. Output File Name.
	if *claSumFile == CLADefaultValue_SumFile {
		claSumFile2 = *claDataFile +
			FileExtensionSeparator + FileOutputExt
	} else {
		claSumFile2 = *claSumFile
	}

	// 3. Action.
	parseAction(*claAction)

	// Report.
	if verboseMode {

		reportFormat = ReportFormatA

		// Action.
		fmt.Printf(
			reportFormat,
			CLAHint_Action,
			*claAction,
		)

		// Input File.
		fmt.Printf(
			reportFormat,
			CLAHint_DataFile,
			*claDataFile,
		)

		// Block Size.
		fmt.Printf(
			reportFormat,
			CLAHint_BlockSize,
			*claBlockSize,
		)

		// Output File.
		fmt.Printf(
			reportFormat,
			CLAHint_SumFile,
			claSumFile2,
		)
	}

	return nil
}

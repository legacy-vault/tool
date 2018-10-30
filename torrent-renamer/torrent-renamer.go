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

// torrent-renamer.go.

// Torrent File Renamer.

// Last Update Time: 2018-10-30.

package main

import (
	"fmt"
	"log"
	"os"
)

// Program's Entry Point.
func main() {

	var duplicatesList map[string][]string
	var duplicatesListLen int
	var err error
	var inputFilePaths []string
	var inputFilePathsCount int
	var userFeedback bool

	// Initialize Command Line Arguments.
	err = initCLA()
	checkInitError(err)

	// List the Files from Input Folder.
	inputFilePaths = listInputFiles()
	inputFilePathsCount = len(inputFilePaths)
	if inputFilePathsCount == 0 {
		fmt.Println(ErrNoFiles)
		os.Exit(ExitCodeNoFilesToProcess)
	}

	// Show Work Summary.
	if appCfg.ReadSubLevels {
		fmt.Println(SubLevelsOn)
	} else {
		fmt.Println(SubLevelsOff)
	}
	if appCfg.OutputFileNameToUpperCase {
		fmt.Println(OutFileUpperCaseOn)
	} else {
		fmt.Println(OutFileUpperCaseOff)
	}
	fmt.Println(FolderIn, appCfg.FolderPathInput)
	fmt.Println(FolderOut, appCfg.FolderPathOutput)

	// Pre-process Files.
	// Check that they can be decoded and their BTIH Sums do not collide.
	// BTIH Collisions are not really important,
	// they are checked to compare the Number of Input and Output Files.
	duplicatesList, err = preProcessFiles(inputFilePaths)
	if err != nil {
		log.Println(err)
		os.Exit(ExitCodeFileProcessingError)
	}

	// If Files have no Syntax Errors, ask the User what to do...

	// 1. Duplicates Count Report.
	duplicatesListLen = len(duplicatesList)
	if duplicatesListLen == 1 {
		fmt.Printf(Duplicate, 1)
	} else {
		fmt.Printf(Duplicates, duplicatesListLen)
	}

	// 2. Show duplicate BTIH Sums?
	if duplicatesListLen > 0 {
		// Get Feedback.
		userFeedback = getUserFeedbackFromKeyboard(QuestionShowDetails)
		if userFeedback {
			for btih, dups := range duplicatesList {
				fmt.Println("BTIH:", btih)
				printStrings(dups)
			}
		}
	}

	// Confirm Processing.
	fmt.Printf(FilesToProcess, inputFilePathsCount)
	userFeedback = getUserFeedbackFromKeyboard("")
	if !userFeedback {
		os.Exit(ExitCodeNormal)
	}

	// Process Files.
	_, err = processFiles(
		inputFilePaths,
		appCfg.FolderPathOutput,
		appCfg.OutputFileNameToUpperCase,
	)
	if err != nil {
		log.Println(err)
		os.Exit(ExitCodeFileProcessingError)
	}
}

// Checker of initial Errors.
func checkInitError(e error) {

	if e != nil {
		fmt.Println(UsageHint)
		os.Exit(ExitCodeInitializationError)
	}
}

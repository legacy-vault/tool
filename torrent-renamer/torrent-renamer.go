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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Program's Entry Point.
func main() {

	var err error
	var inputFilePaths []string
	var inputFilePathsCount int
	var keyboardInput string
	var reader *bufio.Reader

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

	// Confirm Proceeding.
	fmt.Printf(MessagesToProcess, inputFilePathsCount)
	reader = bufio.NewReader(os.Stdin)
	keyboardInput, err = reader.ReadString('\n')
	if err != nil {
		os.Exit(ExitCodeKeyboardInputError)
	}
	keyboardInput = strings.ToLower(keyboardInput)
	keyboardInput = strings.TrimSpace(keyboardInput)
	if keyboardInput != KeyboardInputYes {
		os.Exit(ExitCodeNormal)
	}

	// Process Files.
	err = processFiles(
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

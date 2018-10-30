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

// file.go.

// File Functions.

// Last Update Time: 2018-10-30.

package main

import (
	"fmt"
	"github.com/legacy-vault/library/go/bencode"
	"github.com/legacy-vault/library/go/file"
	"io/ioutil"
	"path"
	"strings"
)

const ExtWantedNormal = "torrent"
const ExtWantedFull = "." + ExtWantedNormal
const OutputFilePermission = 0755

// Lists the Input Files.
func listInputFiles() []string {

	var list []string

	// Get the List with Extensions filtered.
	list = file.ListFilesExtAllowed(
		appCfg.FolderPathInput,
		appCfg.ReadSubLevels,
		[]string{ExtWantedNormal},
	)

	return list
}

// Processes the Files.
// Returns the Number of duplicate Items.
func processFiles(
	inputFilePaths []string,
	outputFolder string,
	outputToUpperCase bool,
) (int, error) {

	var do *bencode.DecodedObject
	var duplicates int
	var err error
	var filePath string
	var outputFileName string
	var outputFilePath string

	for _, filePath = range inputFilePaths {

		// Decode the File to get its BTIH.
		do, err = bencode.ParseFile(filePath)
		if err != nil {
			err = fmt.Errorf(
				ErrFileDecoding,
				filePath,
				err.Error(),
			)
			return duplicates, err
		}

		// Create a Name for Output File.
		outputFileName = do.BTIH.Text
		if outputToUpperCase {
			outputFileName = strings.ToUpper(outputFileName)
		}
		outputFileName = outputFileName + ExtWantedFull
		outputFilePath = path.Join(outputFolder, outputFileName)

		// Check whether an Output File already exists.
		if pathExists(outputFilePath) {
			duplicates++
		}

		// Write the Output File.
		err = ioutil.WriteFile(
			outputFilePath,
			do.SourceData,
			OutputFilePermission,
		)
		if err != nil {
			err = fmt.Errorf(
				ErrFileSaving,
				outputFilePath,
				err.Error(),
			)
			return duplicates, err
		}
	}

	return duplicates, nil
}

// Tries to processes all Files without writing the Results.
// Returns the List (Map) of BTIH Sums with their duplicate Items.
func preProcessFiles(
	inputFilePaths []string,
) (map[string][]string, error) {

	var btih string
	var do *bencode.DecodedObject
	var err error
	var exists bool
	var filePath string
	var list []string
	var outputFileNames map[string][]string

	// Prepare Data.
	// Key = BTIH, Value = List of File Names.
	outputFileNames = make(map[string][]string)

	// Try to process without writing to Disk.
	for _, filePath = range inputFilePaths {

		// Decode the File to get its BTIH.
		do, err = bencode.ParseFile(filePath)
		if err != nil {
			err = fmt.Errorf(
				ErrFileDecoding,
				filePath,
				err.Error(),
			)
			return nil, err
		}
		btih = strings.ToUpper(do.BTIH.Text)

		// Save File Name to List.
		_, exists = outputFileNames[btih]
		if !exists {
			outputFileNames[btih] = []string{}
		}
		outputFileNames[btih] = append(outputFileNames[btih], filePath)
	}

	// Delete non-duplicate Items from the List (Map).
	for btih, list = range outputFileNames {
		if len(list) < 2 {
			delete(outputFileNames, btih)
		}
	}

	return outputFileNames, nil
}

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

package main

import (
	"fmt"
	"github.com/legacy-vault/library/go/bencode"
	"github.com/legacy-vault/library/go/file"
	"io/ioutil"
	"path"
	"strings"
)

const ExtWanted = ".torrent"
const OutputFilePermission = 0755

// Lists the Input Files.
func listInputFiles() []string {

	var cleanList []string
	var rawList []string

	// Get raw List.
	rawList = file.ListFiles(
		appCfg.FolderPathInput,
		appCfg.ReadSubLevels,
	)

	// Filter out the unwanted Files.
	cleanList = []string{}
	for _, file := range rawList {
		if path.Ext(file) == ExtWanted {
			cleanList = append(cleanList, file)
		}
	}

	return cleanList
}

// Processes the Files.
func processFiles(
	inputFilePaths []string,
	outputFolder string,
	outputToUpperCase bool,
) error {

	var do *bencode.DecodedObject
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
			return err
		}

		// Create a Name for Output File.
		outputFileName = do.BTIH.Text
		if outputToUpperCase {
			outputFileName = strings.ToUpper(outputFileName)
		}
		outputFileName = outputFileName + ExtWanted
		outputFilePath = path.Join(outputFolder, outputFileName)

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
			return err
		}
	}

	return nil
}

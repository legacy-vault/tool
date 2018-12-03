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

// check.go.

// Checksum File Checker.

package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const ASCII_LF = '\n'

const PathDot = '.'

const ReportSkipped = "[SKIPPED]"
const ReportMismatch = "[MISMATCH]"
const ReportGood = "[GOOD]"

// Error Messages and Formats.
const ErrFormatAtLine = "Error at Line %d: %s"
const ErrFormatFileIsNotAvailable = "File is not accessible: %s"

// Checks the File.
func checkFile(appConfig Configuration) (Result, error) {

	var checkSumReal []byte
	var checkSumSample []byte
	var err error
	var file *os.File
	var fileIsAccessible bool
	var filePath string
	var line string
	var lineNumber int
	var mismatch Mismatch
	var reader *bufio.Reader
	var result Result

	result.AllClear = false
	result.MatchedFilesCount = 0
	result.Mismatches = make([]Mismatch, 0)
	result.SkippedFilesCount = 0

	// Open the File.
	file, err = os.Open(appConfig.InputFilePath)
	if err != nil {
		return Result{}, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// Read the File, Line by Line...
	reader = bufio.NewReader(file)
	// Read the first Line.
	lineNumber = 1
	line, err = reader.ReadString(ASCII_LF)
	for err == nil {
		// Parse the Line.
		checkSumSample, filePath, err = parseLine(line)
		if err != nil {
			err = fmt.Errorf(
				ErrFormatAtLine,
				lineNumber,
				err.Error(),
			)
			return Result{}, err
		}

		// Process the received File Path...

		// Prepare an absolute Path.
		if pathIsRelative(filePath) {
			filePath = filepath.Join(
				filepath.Dir(appConfig.InputFilePath),
				filePath,
			)
		}

		// 2. Check Availability.
		fileIsAccessible = fileIsAvailable(filePath)
		if !fileIsAccessible {
			if appConfig.IgnoreNotAcessibleFiles {
				if appConfig.ShowProgress {
					fmt.Println(filePath, ReportSkipped)
				}
				result.SkippedFilesCount++
				// Read the next Line.
				line, err = reader.ReadString(ASCII_LF)
				lineNumber++
				continue
			} else {
				err = fmt.Errorf(
					ErrFormatFileIsNotAvailable,
					filePath,
				)
				return Result{}, err
			}
		}

		// 3. Verify the Check Sum.
		checkSumReal, err = calculateFilesMD5(filePath)
		if err != nil {
			return Result{}, err
		}
		if bytes.Compare(checkSumReal, checkSumSample) == 0 {
			// Checksum Match.
			if appConfig.ShowProgress {
				fmt.Println(filePath, ReportGood)
			}
			result.MatchedFilesCount++
		} else {
			// Checksum Mismatch!
			result.AllClear = false
			mismatch.FilePath = filePath
			result.Mismatches = append(result.Mismatches, mismatch)
			if appConfig.ShowProgress {
				fmt.Println(filePath, ReportMismatch)
			}
		}

		// Read the next Line.
		line, err = reader.ReadString(ASCII_LF)
		lineNumber++
		continue
	}
	if (err != nil) && (err != io.EOF) {
		return Result{}, err
	}

	if len(result.Mismatches) == 0 {
		result.AllClear = true
	}

	return result, nil
}

// Parses the Text Line.
// Returns a Checksum and File Path.
func parseLine(
	line string,
) ([]byte, string, error) {

	var err error
	var filePath string
	var hashHexStr string
	var hashHex []byte
	var lineLetters []rune

	// Trim Spaces at Edges.
	line = strings.TrimSpace(line)

	// Get MD5 Hexadecimal String.
	lineLetters = []rune(line)
	hashHexStr = string(lineLetters[0:32])
	hashHex, err = hex.DecodeString(hashHexStr)
	if err != nil {
		return []byte{}, "", err
	}

	// Get File Path.
	filePath = strings.TrimSpace(string(lineLetters[32:]))

	return hashHex, filePath, nil
}

// Checks whether the specified File is available.
func fileIsAvailable(filePath string) bool {

	var err error

	_, err = os.Stat(filePath)
	if err != nil {
		return false
	}

	return true
}

// Checks whether the Path is absolute or relative.
func pathIsRelative(path string) bool {

	if len(path) < 1 {
		return false
	}

	if ([]rune(path))[0] == PathDot {
		return true
	}

	return false
}

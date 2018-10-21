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

// create.go.

// Block CRC :: File Creation.

package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

const MsgBytesReceived = "Total Bytes Received:"
const MsgBytesWritten = "Total Bytes Written:"

// Creates a File with Block CRC Data.
// Reads Data from the Input Data File.
// Writes Check-Sum to the Output Sum File.
func bCRCCreateFile(
	dataFilePath string,
	sumFilePath string,
	blockSize uint64,
) error {

	var blockSizeBA []byte
	var bytesRead int
	var bytesReadTotal uint64
	var bytesWritten int
	var bytesWrittenTotal uint64
	var currentBlock []byte
	var crcTable *crc32.Table
	var dataFile *os.File
	var err error
	var i uint64
	var slidingTotalBlock []byte // "Running Total" Block.
	var sumFile *os.File
	var sumOfCurrentBlock uint32
	var sumOfCurrentBlockBA []byte
	var sumOfSlidingBlock uint32
	var sumOfSlidingBlockBA []byte

	// Open the Input (Data) File.
	dataFile, err = os.Open(dataFilePath)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	// Create the Output (Sum) File.
	sumFile, err = os.Create(sumFilePath)
	if err != nil {
		return err
	}
	defer sumFile.Close()

	// Write Output (Sum) File Header.
	blockSizeBA = make([]byte, 8)
	binary.BigEndian.PutUint64(blockSizeBA, blockSize)
	bytesWritten, err = sumFile.Write(blockSizeBA)
	if err != nil {
		return err
	}
	bytesWrittenTotal = bytesWrittenTotal + uint64(bytesWritten)

	// Preparations.

	// 1. Buffer for current Block.
	currentBlock = make([]byte, blockSize)
	slidingTotalBlock = make([]byte, 0)

	// 2. CRC Table and CRC Sum Holders.
	crcTable = crc32.IEEETable
	sumOfCurrentBlockBA = make([]byte, 4)
	sumOfSlidingBlockBA = make([]byte, 4)

	// Read all Blocks from Input (Data) File.
	for true {

		bytesRead, err = dataFile.Read(currentBlock)
		bytesReadTotal = bytesReadTotal + uint64(bytesRead)

		// End of File?
		if err == io.EOF {
			break
		}
		if err != nil {
			// An Error has occured, and it is not EOF!
			return err
		}

		// Update sliding Total Block.
		slidingTotalBlock = append(slidingTotalBlock, currentBlock...)

		// No Errors. Process a Block.
		sumOfCurrentBlock = crc32.Checksum(currentBlock, crcTable)
		binary.BigEndian.PutUint32(sumOfCurrentBlockBA, sumOfCurrentBlock)
		sumOfSlidingBlock = crc32.Checksum(slidingTotalBlock, crcTable)
		binary.BigEndian.PutUint32(sumOfSlidingBlockBA, sumOfSlidingBlock)

		// Save Checksums to the Output (Sum) File.
		bytesWritten, err = sumFile.Write(sumOfCurrentBlockBA)
		if err != nil {
			return err
		}
		bytesWrittenTotal = bytesWrittenTotal + uint64(bytesWritten)
		bytesWritten, err = sumFile.Write(sumOfSlidingBlockBA)
		if err != nil {
			return err
		}
		bytesWrittenTotal = bytesWrittenTotal + uint64(bytesWritten)

		// Clear the Buffer before new Data arrives.
		for i = 0; i < blockSize; i++ {
			currentBlock[i] = 0
		}
	}

	// Summary Report.
	fmt.Println(MsgBytesReceived, bytesReadTotal)
	fmt.Println(MsgBytesWritten, bytesWrittenTotal)

	return nil
}

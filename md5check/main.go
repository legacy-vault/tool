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

// main.go.

// Application's Entry Point.

package main

import (
	"fmt"
	"log"
)

const AppName = "md5check"
const AppVersion = "1.0"
const AppDescription = "This Tool compares Checksums in the specified File " +
	"with the real Checksums of the Files mentioned in the specified File. " +
	"The Format of each Line of the specified File must be the following:\r\n" +
	"\t <MD5 CHECKSUM> <TABULATOR OR SPACE> <FILE PATH>\r\n"

var AppShortInfo = fmt.Sprintf("%s v%s", AppName, AppVersion)
var MsgHelp = fmt.Sprintf(
	"Usage:\r\n"+
		"\t%s [OPTIONS] <FILE>\r\n"+
		"\t%s <FILE> [OPTIONS]\r\n"+
		"\t%s [OPTIONS] <FILE> [OPTIONS]\r\n"+
		"\r\n"+
		"Possible Options:\r\n"+

		"\t"+
		OptionHelp1+", "+
		OptionHelp2+
		" = Help Mode.\r\n"+

		"\t"+
		OptionIgnore1+", "+
		OptionIgnore2+
		" = Ignore Files which are not accessible.\r\n"+

		"\t"+
		OptionShowProgress1+", "+
		OptionShowProgress2+
		" = Show Result for each File.\r\n"+

		"\t"+
		OptionVerboseMode1+", "+
		OptionVerboseMode2+
		" = Verbose Mode.\r\n",
	AppName,
	AppName,
	AppName,
)

// Applications Entry point.
func main() {

	var appConfig Configuration
	var appResult Result
	var err error

	// Configure the Application.
	err = configureApp(&appConfig)
	if err != nil {
		log.Fatal("Application Configuration Error: ", err)
	}

	// Run the Application.

	// Help Mode?
	if appConfig.HelpMode {
		if appConfig.VerboseMode {
			fmt.Println(AppShortInfo)
			fmt.Println()
			fmt.Println(AppDescription)
			fmt.Println()
		}
		fmt.Println(MsgHelp)
		return
	}

	// Check the File.
	appResult, err = checkFile(appConfig)
	if err != nil {
		log.Fatal("Application Error: ", err)
	}

	showResults(appConfig, appResult)
	return
}

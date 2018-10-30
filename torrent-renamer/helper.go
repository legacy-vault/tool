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

// helper.go.

// Helper-Functions.

// Last Update Time: 2018-10-30.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Checks whether the Path exists.
func pathExists(path string) bool {

	var err error

	_, err = os.Stat(path)
	if err != nil {
		return false
	}

	return true
}

// Gets User's Feedback from Keyboard.
func getUserFeedbackFromKeyboard(question string) bool {

	var err error
	var keyboardInput string
	var reader *bufio.Reader

	// Ask the Question.
	if len(question) > 0 {
		fmt.Println(question)
	}

	// Print the Hint.
	fmt.Print(Proceed)

	// Read from standard Input Stream (Keyboard).
	reader = bufio.NewReader(os.Stdin)
	keyboardInput, err = reader.ReadString('\n')
	if err != nil {
		os.Exit(ExitCodeKeyboardInputError)
	}
	keyboardInput = strings.ToLower(keyboardInput)
	keyboardInput = strings.TrimSpace(keyboardInput)

	// Check.
	if keyboardInput != KeyboardInputYes {
		return false
	}

	return true
}

// Prints the List of Strings.
func printStrings(list []string) {

	var s string

	for _, s = range list {
		fmt.Println(s)
	}
}

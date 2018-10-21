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

// self_check.go.

// Block CRC :: Program's Self Check

package main

import "errors"

// Error Messages.
const ErrMsgSelfCheckStageIndices = "Stage Indices Self Check Error"
const ErrMsgSelfCheckStatusIndices = "Stage Status Indices Self Check Error"

// Checks Application's Internals.
func selfCheck() error {

	var err error

	// Check Stage Indices.
	err = checkStageIndices()
	if err != nil {
		return err
	}

	return nil
}

// Checks Stage Indices.
func checkStageIndices() error {

	var err error

	// Stage Indices.
	if (StageIndexMax - StageIndexMin + 1) != StagesCount {
		err = errors.New(ErrMsgSelfCheckStageIndices)
		return err
	}

	// Stage Status Indices.
	if (StatusIndexMax - StatusIndexMin + 1) != StatusesCount {
		err = errors.New(ErrMsgSelfCheckStatusIndices)
		return err
	}

	return nil
}

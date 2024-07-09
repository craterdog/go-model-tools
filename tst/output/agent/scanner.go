/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	// Initialize class constants.
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	// Define class constants.
}

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	// Validate the arguments.
	switch {
	case isUndefined(source):
		panic("The source attribute is required for each Scanner.")
	case isUndefined(tokens):
		panic("The tokens attribute is required for each Scanner.")
	default:
		return &scanner_{
			// Initialize instance attributes.
			class_: c,
			source_: source,
			tokens_: tokens,
		}
	}
}

// Functions

func (c *scannerClass_) FormatToken(token TokenLike) string {
	var result_ string
	// TBA - Implement the function.
	return result_
}

func (c *scannerClass_) MatchToken(
	type_ TokenType,
	text string,
) col.ListLike[string] {
	var result_ col.ListLike[string]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	// Define instance attributes.
	class_ ScannerClassLike
	source_ string
	tokens_ col.QueueLike[TokenLike]
}

// Attributes

func (v *scanner_) GetClass() ScannerClassLike {
	return v.class_
}

// Private

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
	mod "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var tokenClass = &tokenClass_{
	// Initialize class constants.
}

// Function

func Token() TokenClassLike {
	return tokenClass
}

// CLASS METHODS

// Target

type tokenClass_ struct {
	// Define class constants.
}

// Constructors

func (c *tokenClass_) Make(
	line int,
	position int,
	type_ TokenType,
	value string,
) TokenLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(line):
		panic("The line attribute is required for each Token.")
	case mod.IsUndefined(position):
		panic("The position attribute is required for each Token.")
	case mod.IsUndefined(type_):
		panic("The type_ attribute is required for each Token.")
	case mod.IsUndefined(value):
		panic("The value attribute is required for each Token.")
	default:
		return &token_{
			// Initialize instance attributes.
			class_: c,
			line_: line,
			position_: position,
			type_: type_,
			value_: value,
		}
	}
}

// INSTANCE METHODS

// Target

type token_ struct {
	// Define instance attributes.
	class_ TokenClassLike
	line_ int
	position_ int
	type_ TokenType
	value_ string
}

// Attributes

func (v *token_) GetClass() TokenClassLike {
	return v.class_
}

func (v *token_) GetLine() int {
	return v.line_
}

func (v *token_) GetPosition() int {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}

// Private

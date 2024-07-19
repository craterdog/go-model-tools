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

package ast

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var characterClass = &characterClass_{
	// Initialize class constants.
}

// Function

func Character() CharacterClassLike {
	return characterClass
}

// CLASS METHODS

// Target

type characterClass_ struct {
	// Define class constants.
}

// Constructors

func (c *characterClass_) Make(any_ any) CharacterLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &character_{
			// Initialize instance attributes.
			class_: c,
			any_: any_,
		}
	}
}

// INSTANCE METHODS

// Target

type character_ struct {
	// Define instance attributes.
	class_ CharacterClassLike
	any_ any
}

// Attributes

func (v *character_) GetClass() CharacterClassLike {
	return v.class_
}

func (v *character_) GetAny() any {
	return v.any_
}

// Private

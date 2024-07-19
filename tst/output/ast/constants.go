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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var constantsClass = &constantsClass_{
	// Initialize class constants.
}

// Function

func Constants() ConstantsClassLike {
	return constantsClass
}

// CLASS METHODS

// Target

type constantsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constantsClass_) Make(
	note string,
	constants abs.Sequential[ConstantLike],
) ConstantsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(note):
		panic("The note attribute is required by this class.")
	case col.IsUndefined(constants):
		panic("The constants attribute is required by this class.")
	default:
		return &constants_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			constants_: constants,
		}
	}
}

// INSTANCE METHODS

// Target

type constants_ struct {
	// Define instance attributes.
	class_ ConstantsClassLike
	note_ string
	constants_ abs.Sequential[ConstantLike]
}

// Attributes

func (v *constants_) GetClass() ConstantsClassLike {
	return v.class_
}

func (v *constants_) GetNote() string {
	return v.note_
}

func (v *constants_) GetConstants() abs.Sequential[ConstantLike] {
	return v.constants_
}

// Private

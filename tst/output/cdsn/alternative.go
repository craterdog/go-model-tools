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
	col "github.com/craterdog/go-collection-framework/v4/collection"
	mod "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var alternativeClass = &alternativeClass_{
	// Initialize class constants.
}

// Function

func Alternative() AlternativeClassLike {
	return alternativeClass
}

// CLASS METHODS

// Target

type alternativeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *alternativeClass_) Make(parts col.Sequential[PartLike]) AlternativeLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(parts):
		panic("The parts attribute is required for each Alternative.")
	default:
		return &alternative_{
			// Initialize instance attributes.
			class_: c,
			parts_: parts,
		}
	}
}

// INSTANCE METHODS

// Target

type alternative_ struct {
	// Define instance attributes.
	class_ AlternativeClassLike
	parts_ col.Sequential[PartLike]
}

// Attributes

func (v *alternative_) GetClass() AlternativeClassLike {
	return v.class_
}

func (v *alternative_) GetParts() col.Sequential[PartLike] {
	return v.parts_
}

// Private

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

import ()

// CLASS ACCESS

// Reference

var cardinalityClass = &cardinalityClass_{
	// Initialize class constants.
}

// Function

func Cardinality() CardinalityClassLike {
	return cardinalityClass
}

// CLASS METHODS

// Target

type cardinalityClass_ struct {
	// Define class constants.
}

// Constructors

func (c *cardinalityClass_) MakeWithConstraint(constraint ConstraintLike) CardinalityLike {
	return &cardinality_{
		// Initialize instance attributes.
		class_: c,
		constraint_: constraint,
	}
}

// INSTANCE METHODS

// Target

type cardinality_ struct {
	// Define instance attributes.
	class_ CardinalityClassLike
	constraint_ ConstraintLike
}

// Attributes

func (v *cardinality_) GetClass() CardinalityClassLike {
	return v.class_
}

func (v *cardinality_) GetConstraint() ConstraintLike {
	return v.constraint_
}

// Private

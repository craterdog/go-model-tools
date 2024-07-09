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

package collection

// CLASS ACCESS

// Reference

var notationClass = &notationClass_{
	// Initialize class constants.
}

// Function

func Notation() NotationClassLike {
	return notationClass
}

// CLASS METHODS

// Target

type notationClass_ struct {
	// Define class constants.
}

// Constructors

func (c *notationClass_) Make() NotationLike {
	// Validate the arguments.
	switch {
	default:
		return &notation_{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// INSTANCE METHODS

// Target

type notation_ struct {
	// Define instance attributes.
	class_ NotationClassLike
}

// Attributes

func (v *notation_) GetClass() NotationClassLike {
	return v.class_
}

// Canonical

func (v *notation_) ParseSource(source string) (value any) {
	// TBA - Implement the method.
	return
}

func (v *notation_) FormatValue(value any) (source string) {
	// TBA - Implement the method.
	return
}

// Private

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
	ref "reflect"
)

// CLASS ACCESS

// Reference

var lineClass = &lineClass_{
	// Initialize class constants.
}

// Function

func Line() LineClassLike {
	return lineClass
}

// CLASS METHODS

// Target

type lineClass_ struct {
	// Define class constants.
}

// Constructors

func (c *lineClass_) Make(
	identifier IdentifierLike,
	note string,
) LineLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(identifier):
		panic("The identifier attribute is required for each Line.")
	case c.isUndefined(note):
		panic("The note attribute is required for each Line.")
	default:
		return &line_{
			// Initialize instance attributes.
			class_: c,
			identifier_: identifier,
			note_: note,
		}
	}
}

// Private

func (c *lineClass_) isUndefined(value any) bool {
	switch actual := value.(type) {
	case string:
		return len(actual) > 0
	default:
		var meta = ref.ValueOf(actual)
		return (meta.Kind() == ref.Ptr ||
			meta.Kind() == ref.Interface ||
			meta.Kind() == ref.Slice ||
			meta.Kind() == ref.Map ||
			meta.Kind() == ref.Chan ||
			meta.Kind() == ref.Func) && meta.IsNil()
	}
}

// INSTANCE METHODS

// Target

type line_ struct {
	// Define instance attributes.
	class_ LineClassLike
	identifier_ IdentifierLike
	note_ string
}

// Attributes

func (v *line_) GetClass() LineClassLike {
	return v.class_
}

func (v *line_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *line_) GetNote() string {
	return v.note_
}

// Private

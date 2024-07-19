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

var attributesClass = &attributesClass_{
	// Initialize class constants.
}

// Function

func Attributes() AttributesClassLike {
	return attributesClass
}

// CLASS METHODS

// Target

type attributesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *attributesClass_) Make(
	note string,
	attributes abs.Sequential[AttributeLike],
) AttributesLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(note):
		panic("The note attribute is required by this class.")
	case col.IsUndefined(attributes):
		panic("The attributes attribute is required by this class.")
	default:
		return &attributes_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			attributes_: attributes,
		}
	}
}

// INSTANCE METHODS

// Target

type attributes_ struct {
	// Define instance attributes.
	class_ AttributesClassLike
	note_ string
	attributes_ abs.Sequential[AttributeLike]
}

// Attributes

func (v *attributes_) GetClass() AttributesClassLike {
	return v.class_
}

func (v *attributes_) GetNote() string {
	return v.note_
}

func (v *attributes_) GetAttributes() abs.Sequential[AttributeLike] {
	return v.attributes_
}

// Private

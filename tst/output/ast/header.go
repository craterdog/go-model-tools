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

var headerClass = &headerClass_{
	// Initialize class constants.
}

// Function

func Header() HeaderClassLike {
	return headerClass
}

// CLASS METHODS

// Target

type headerClass_ struct {
	// Define class constants.
}

// Constructors

func (c *headerClass_) Make(
	comment string,
	name string,
) HeaderLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(comment):
		panic("The comment attribute is required by this class.")
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	default:
		return &header_{
			// Initialize instance attributes.
			class_: c,
			comment_: comment,
			name_: name,
		}
	}
}

// INSTANCE METHODS

// Target

type header_ struct {
	// Define instance attributes.
	class_ HeaderClassLike
	comment_ string
	name_ string
}

// Attributes

func (v *header_) GetClass() HeaderClassLike {
	return v.class_
}

func (v *header_) GetComment() string {
	return v.comment_
}

func (v *header_) GetName() string {
	return v.name_
}

// Private

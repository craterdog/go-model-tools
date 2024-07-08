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

func (c *headerClass_) Make(comment string) HeaderLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(comment):
		panic("The comment attribute is required for each Header.")
	default:
		return &header_{
			// Initialize instance attributes.
			class_: c,
			comment_: comment,
		}
	}
}

// Private

func (c *headerClass_) isUndefined(value any) bool {
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

type header_ struct {
	// Define instance attributes.
	class_ HeaderClassLike
	comment_ string
}

// Attributes

func (v *header_) GetClass() HeaderClassLike {
	return v.class_
}

func (v *header_) GetComment() string {
	return v.comment_
}

// Private

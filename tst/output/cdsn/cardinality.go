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

func (c *cardinalityClass_) Make(any_ any) CardinalityLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(any_):
		panic("The any_ attribute is required for each Cardinality.")
	default:
		return &cardinality_{
			// Initialize instance attributes.
			class_: c,
			any_: any_,
		}
	}
}

// Private

func (c *cardinalityClass_) isUndefined(value any) bool {
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

type cardinality_ struct {
	// Define instance attributes.
	class_ CardinalityClassLike
	any_ any
}

// Attributes

func (v *cardinality_) GetClass() CardinalityClassLike {
	return v.class_
}

func (v *cardinality_) GetAny() any {
	return v.any_
}

// Private

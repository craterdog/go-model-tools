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

var maximumClass = &maximumClass_{
	// Initialize class constants.
}

// Function

func Maximum() MaximumClassLike {
	return maximumClass
}

// CLASS METHODS

// Target

type maximumClass_ struct {
	// Define class constants.
}

// Constructors

func (c *maximumClass_) Make(number string) MaximumLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(number):
		panic("The number attribute is required for each Maximum.")
	default:
		return &maximum_{
			// Initialize instance attributes.
			class_: c,
			number_: number,
		}
	}
}

// Private

func (c *maximumClass_) isUndefined(value any) bool {
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

type maximum_ struct {
	// Define instance attributes.
	class_ MaximumClassLike
	number_ string
}

// Attributes

func (v *maximum_) GetClass() MaximumClassLike {
	return v.class_
}

func (v *maximum_) GetNumber() string {
	return v.number_
}

// Private

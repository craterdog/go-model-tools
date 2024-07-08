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

var prefixClass = &prefixClass_{
	// Initialize class constants.
}

// Function

func Prefix() PrefixClassLike {
	return prefixClass
}

// CLASS METHODS

// Target

type prefixClass_ struct {
	// Define class constants.
}

// Constructors

func (c *prefixClass_) Make(any any) PrefixLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(any):
		panic("The any attribute is required for each Prefix.")
	default:
		return &prefix_{
			// Initialize instance attributes.
			class_: c,
			any_: any,
		}
	}
}

// Private

func (c *prefixClass_) isUndefined(value any) bool {
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

type prefix_ struct {
	// Define instance attributes.
	class_ PrefixClassLike
	any_ any
}

// Attributes

func (v *prefix_) GetClass() PrefixClassLike {
	return v.class_
}

func (v *prefix_) GetAny() any {
	return v.any_
}

// Private

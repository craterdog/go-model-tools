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

var resultClass = &resultClass_{
	// Initialize class constants.
}

// Function

func Result() ResultClassLike {
	return resultClass
}

// CLASS METHODS

// Target

type resultClass_ struct {
	// Define class constants.
}

// Constructors

func (c *resultClass_) Make(any any) ResultLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(any):
		panic("The any attribute is required for each Result.")
	default:
		return &result_{
			// Initialize instance attributes.
			class_: c,
			any_: any,
		}
	}
}

// Private

func (c *resultClass_) isUndefined(value any) bool {
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

type result_ struct {
	// Define instance attributes.
	class_ ResultClassLike
	any_ any
}

// Attributes

func (v *result_) GetClass() ResultClassLike {
	return v.class_
}

func (v *result_) GetAny() any {
	return v.any_
}

// Private

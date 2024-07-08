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

var boundedClass = &boundedClass_{
	// Initialize class constants.
}

// Function

func Bounded() BoundedClassLike {
	return boundedClass
}

// CLASS METHODS

// Target

type boundedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *boundedClass_) Make(
	initial InitialLike,
	extent ExtentLike,
) BoundedLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(initial):
		panic("The initial attribute is required for each Bounded.")
	case c.isUndefined(extent):
		panic("The extent attribute is required for each Bounded.")
	default:
		return &bounded_{
			// Initialize instance attributes.
			class_: c,
			initial_: initial,
			extent_: extent,
		}
	}
}

// Private

func (c *boundedClass_) isUndefined(value any) bool {
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

type bounded_ struct {
	// Define instance attributes.
	class_ BoundedClassLike
	initial_ InitialLike
	extent_ ExtentLike
}

// Attributes

func (v *bounded_) GetClass() BoundedClassLike {
	return v.class_
}

func (v *bounded_) GetInitial() InitialLike {
	return v.initial_
}

func (v *bounded_) GetExtent() ExtentLike {
	return v.extent_
}

// Private

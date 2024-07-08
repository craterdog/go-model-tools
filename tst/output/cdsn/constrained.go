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

var constrainedClass = &constrainedClass_{
	// Initialize class constants.
}

// Function

func Constrained() ConstrainedClassLike {
	return constrainedClass
}

// CLASS METHODS

// Target

type constrainedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constrainedClass_) Make(
	minimum MinimumLike,
	maximum MaximumLike,
) ConstrainedLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(minimum):
		panic("The minimum attribute is required for each Constrained.")
	case c.isUndefined(maximum):
		panic("The maximum attribute is required for each Constrained.")
	default:
		return &constrained_{
			// Initialize instance attributes.
			class_: c,
			minimum_: minimum,
			maximum_: maximum,
		}
	}
}

// Private

func (c *constrainedClass_) isUndefined(value any) bool {
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

type constrained_ struct {
	// Define instance attributes.
	class_ ConstrainedClassLike
	minimum_ MinimumLike
	maximum_ MaximumLike
}

// Attributes

func (v *constrained_) GetClass() ConstrainedClassLike {
	return v.class_
}

func (v *constrained_) GetMinimum() MinimumLike {
	return v.minimum_
}

func (v *constrained_) GetMaximum() MaximumLike {
	return v.maximum_
}

// Private

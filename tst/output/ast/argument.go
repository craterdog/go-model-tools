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

var argumentClass = &argumentClass_{
	// Initialize class constants.
}

// Function

func Argument() ArgumentClassLike {
	return argumentClass
}

// CLASS METHODS

// Target

type argumentClass_ struct {
	// Define class constants.
}

// Constructors

func (c *argumentClass_) Make(abstraction AbstractionLike) ArgumentLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(abstraction):
		panic("The abstraction attribute is required for each Argument.")
	default:
		return &argument_{
			// Initialize instance attributes.
			class_: c,
			abstraction_: abstraction,
		}
	}
}

// Private

func (c *argumentClass_) isUndefined(value any) bool {
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

type argument_ struct {
	// Define instance attributes.
	class_ ArgumentClassLike
	abstraction_ AbstractionLike
}

// Attributes

func (v *argument_) GetClass() ArgumentClassLike {
	return v.class_
}

func (v *argument_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private

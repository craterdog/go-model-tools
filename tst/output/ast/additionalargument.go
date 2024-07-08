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

var additionalArgumentClass = &additionalArgumentClass_{
	// Initialize class constants.
}

// Function

func AdditionalArgument() AdditionalArgumentClassLike {
	return additionalArgumentClass
}

// CLASS METHODS

// Target

type additionalArgumentClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalArgumentClass_) Make(argument ArgumentLike) AdditionalArgumentLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(argument):
		panic("The argument attribute is required for each AdditionalArgument.")
	default:
		return &additionalArgument_{
			// Initialize instance attributes.
			class_: c,
			argument_: argument,
		}
	}
}

// Private

func (c *additionalArgumentClass_) isUndefined(value any) bool {
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

type additionalArgument_ struct {
	// Define instance attributes.
	class_ AdditionalArgumentClassLike
	argument_ ArgumentLike
}

// Attributes

func (v *additionalArgument_) GetClass() AdditionalArgumentClassLike {
	return v.class_
}

func (v *additionalArgument_) GetArgument() ArgumentLike {
	return v.argument_
}

// Private

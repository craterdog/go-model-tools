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

var genericArgumentsClass = &genericArgumentsClass_{
	// Initialize class constants.
}

// Function

func GenericArguments() GenericArgumentsClassLike {
	return genericArgumentsClass
}

// CLASS METHODS

// Target

type genericArgumentsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *genericArgumentsClass_) Make(arguments ArgumentsLike) GenericArgumentsLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(arguments):
		panic("The arguments attribute is required for each GenericArguments.")
	default:
		return &genericArguments_{
			// Initialize instance attributes.
			class_: c,
			arguments_: arguments,
		}
	}
}

// Private

func (c *genericArgumentsClass_) isUndefined(value any) bool {
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

type genericArguments_ struct {
	// Define instance attributes.
	class_ GenericArgumentsClassLike
	arguments_ ArgumentsLike
}

// Attributes

func (v *genericArguments_) GetClass() GenericArgumentsClassLike {
	return v.class_
}

func (v *genericArguments_) GetArguments() ArgumentsLike {
	return v.arguments_
}

// Private

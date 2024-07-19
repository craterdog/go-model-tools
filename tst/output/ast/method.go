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
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var methodClass = &methodClass_{
	// Initialize class constants.
}

// Function

func Method() MethodClassLike {
	return methodClass
}

// CLASS METHODS

// Target

type methodClass_ struct {
	// Define class constants.
}

// Constructors

func (c *methodClass_) Make(
	name string,
	optionalParameters ParametersLike,
	optionalResult ResultLike,
) MethodLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	default:
		return &method_{
			// Initialize instance attributes.
			class_: c,
			name_: name,
			optionalParameters_: optionalParameters,
			optionalResult_: optionalResult,
		}
	}
}

// INSTANCE METHODS

// Target

type method_ struct {
	// Define instance attributes.
	class_ MethodClassLike
	name_ string
	optionalParameters_ ParametersLike
	optionalResult_ ResultLike
}

// Attributes

func (v *method_) GetClass() MethodClassLike {
	return v.class_
}

func (v *method_) GetName() string {
	return v.name_
}

func (v *method_) GetOptionalParameters() ParametersLike {
	return v.optionalParameters_
}

func (v *method_) GetOptionalResult() ResultLike {
	return v.optionalResult_
}

// Private

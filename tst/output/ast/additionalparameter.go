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

var additionalParameterClass = &additionalParameterClass_{
	// Initialize class constants.
}

// Function

func AdditionalParameter() AdditionalParameterClassLike {
	return additionalParameterClass
}

// CLASS METHODS

// Target

type additionalParameterClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalParameterClass_) Make(parameter ParameterLike) AdditionalParameterLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(parameter):
		panic("The parameter attribute is required for each AdditionalParameter.")
	default:
		return &additionalParameter_{
			// Initialize instance attributes.
			class_: c,
			parameter_: parameter,
		}
	}
}

// Private

func (c *additionalParameterClass_) isUndefined(value any) bool {
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

type additionalParameter_ struct {
	// Define instance attributes.
	class_ AdditionalParameterClassLike
	parameter_ ParameterLike
}

// Attributes

func (v *additionalParameter_) GetClass() AdditionalParameterClassLike {
	return v.class_
}

func (v *additionalParameter_) GetParameter() ParameterLike {
	return v.parameter_
}

// Private

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
	parameters ParametersLike,
	result ResultLike,
) MethodLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(name):
		panic("The name attribute is required for each Method.")
	case c.isUndefined(parameters):
		panic("The parameters attribute is required for each Method.")
	case c.isUndefined(result):
		panic("The result attribute is required for each Method.")
	default:
		return &method_{
			// Initialize instance attributes.
			class_: c,
			name_: name,
			parameters_: parameters,
			result_: result,
		}
	}
}

// Private

func (c *methodClass_) isUndefined(value any) bool {
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

type method_ struct {
	// Define instance attributes.
	class_ MethodClassLike
	name_ string
	parameters_ ParametersLike
	result_ ResultLike
}

// Attributes

func (v *method_) GetClass() MethodClassLike {
	return v.class_
}

func (v *method_) GetName() string {
	return v.name_
}

func (v *method_) GetOptionalParameters() ParametersLike {
	return v.parameters_
}

func (v *method_) GetOptionalResult() ResultLike {
	return v.result_
}

// Private

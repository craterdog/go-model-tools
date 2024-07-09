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

// CLASS ACCESS

// Reference

var functionClass = &functionClass_{
	// Initialize class constants.
}

// Function

func Function() FunctionClassLike {
	return functionClass
}

// CLASS METHODS

// Target

type functionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *functionClass_) Make(
	name string,
	parameters ParametersLike,
	result ResultLike,
) FunctionLike {
	// Validate the arguments.
	switch {
	case isUndefined(name):
		panic("The name attribute is required for each Function.")
	case isUndefined(parameters):
		panic("The parameters attribute is required for each Function.")
	case isUndefined(result):
		panic("The result attribute is required for each Function.")
	default:
		return &function_{
			// Initialize instance attributes.
			class_: c,
			name_: name,
			parameters_: parameters,
			result_: result,
		}
	}
}

// INSTANCE METHODS

// Target

type function_ struct {
	// Define instance attributes.
	class_ FunctionClassLike
	name_ string
	optionalParameters_ ParametersLike
	result_ ResultLike
	parameters_ ParametersLike
}

// Attributes

func (v *function_) GetClass() FunctionClassLike {
	return v.class_
}

func (v *function_) GetName() string {
	return v.name_
}

func (v *function_) GetOptionalParameters() ParametersLike {
	return v.optionalParameters_
}

func (v *function_) GetResult() ResultLike {
	return v.result_
}

// Private

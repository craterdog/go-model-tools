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

package gcmn

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var functionClass = &functionClass_{
	// Any private class constants should be initialized here.
}

// Function

func Function() FunctionClassLike {
	return functionClass
}

// CLASS METHODS

// Target

type functionClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *functionClass_) MakeWithAttributes(
	identifier string,
	parameters col.ListLike[ParameterLike],
	result ResultLike,
) FunctionLike {
	return &function_{
		identifier_: identifier,
		parameters_: parameters,
		result_: result,
	}
}

// Functions

// INSTANCE METHODS

// Target

type function_ struct {
	class_ FunctionClassLike
	identifier_ string
	parameters_ col.ListLike[ParameterLike]
	result_ ResultLike
}

// Attributes

func (v *function_) GetClass() FunctionClassLike {
	return v.class_
}

func (v *function_) GetIdentifier() string {
	return v.identifier_
}

func (v *function_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

func (v *function_) GetResult() ResultLike {
	return v.result_
}

// Public

// Private
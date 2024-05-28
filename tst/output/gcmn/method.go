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
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var methodClass = &methodClass_{
	// Any private class constants should be initialized here.
}

// Function

func Method() MethodClassLike {
	return methodClass
}

// CLASS METHODS

// Target

type methodClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *methodClass_) MakeWithAttributes(
	identifier string,
	parameters col.ListLike[ParameterLike],
	result ResultLike,
) MethodLike {
	return &method_{
		identifier_: identifier,
		parameters_: parameters,
		result_: result,
	}
}

// Functions

// INSTANCE METHODS

// Target

type method_ struct {
	class_ MethodClassLike
	identifier_ string
	parameters_ col.ListLike[ParameterLike]
	result_ ResultLike
}

// Attributes

func (v *method_) GetClass() MethodClassLike {
	return v.class_
}

func (v *method_) GetIdentifier() string {
	return v.identifier_
}

func (v *method_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

func (v *method_) GetResult() ResultLike {
	return v.result_
}

// Public

// Private

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
	mod "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var argumentsClass = &argumentsClass_{
	// Initialize class constants.
}

// Function

func Arguments() ArgumentsClassLike {
	return argumentsClass
}

// CLASS METHODS

// Target

type argumentsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *argumentsClass_) Make(
	argument ArgumentLike,
	additionalArguments col.Sequential[AdditionalArgumentLike],
) ArgumentsLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(argument):
		panic("The argument attribute is required for each Arguments.")
	case mod.IsUndefined(additionalArguments):
		panic("The additionalArguments attribute is required for each Arguments.")
	default:
		return &arguments_{
			// Initialize instance attributes.
			class_: c,
			argument_: argument,
			additionalArguments_: additionalArguments,
		}
	}
}

// INSTANCE METHODS

// Target

type arguments_ struct {
	// Define instance attributes.
	class_ ArgumentsClassLike
	argument_ ArgumentLike
	additionalArguments_ col.Sequential[AdditionalArgumentLike]
}

// Attributes

func (v *arguments_) GetClass() ArgumentsClassLike {
	return v.class_
}

func (v *arguments_) GetArgument() ArgumentLike {
	return v.argument_
}

func (v *arguments_) GetAdditionalArguments() col.Sequential[AdditionalArgumentLike] {
	return v.additionalArguments_
}

// Private

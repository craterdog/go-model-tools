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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var functionalsClass = &functionalsClass_{
	// Initialize class constants.
}

// Function

func Functionals() FunctionalsClassLike {
	return functionalsClass
}

// CLASS METHODS

// Target

type functionalsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *functionalsClass_) Make(
	note string,
	functionals abs.Sequential[FunctionalLike],
) FunctionalsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(note):
		panic("The note attribute is required by this class.")
	case col.IsUndefined(functionals):
		panic("The functionals attribute is required by this class.")
	default:
		return &functionals_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			functionals_: functionals,
		}
	}
}

// INSTANCE METHODS

// Target

type functionals_ struct {
	// Define instance attributes.
	class_ FunctionalsClassLike
	note_ string
	functionals_ abs.Sequential[FunctionalLike]
}

// Attributes

func (v *functionals_) GetClass() FunctionalsClassLike {
	return v.class_
}

func (v *functionals_) GetNote() string {
	return v.note_
}

func (v *functionals_) GetFunctionals() abs.Sequential[FunctionalLike] {
	return v.functionals_
}

// Private

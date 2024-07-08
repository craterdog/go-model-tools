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

var declarationClass = &declarationClass_{
	// Initialize class constants.
}

// Function

func Declaration() DeclarationClassLike {
	return declarationClass
}

// CLASS METHODS

// Target

type declarationClass_ struct {
	// Define class constants.
}

// Constructors

func (c *declarationClass_) Make(
	comment string,
	name string,
	genericParameters GenericParametersLike,
) DeclarationLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(comment):
		panic("The comment attribute is required for each Declaration.")
	case c.isUndefined(name):
		panic("The name attribute is required for each Declaration.")
	case c.isUndefined(genericParameters):
		panic("The genericParameters attribute is required for each Declaration.")
	default:
		return &declaration_{
			// Initialize instance attributes.
			class_: c,
			comment_: comment,
			name_: name,
			genericParameters_: genericParameters,
		}
	}
}

// Private

func (c *declarationClass_) isUndefined(value any) bool {
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

type declaration_ struct {
	// Define instance attributes.
	class_ DeclarationClassLike
	comment_ string
	name_ string
	genericParameters_ GenericParametersLike
}

// Attributes

func (v *declaration_) GetClass() DeclarationClassLike {
	return v.class_
}

func (v *declaration_) GetComment() string {
	return v.comment_
}

func (v *declaration_) GetName() string {
	return v.name_
}

func (v *declaration_) GetOptionalGenericParameters() GenericParametersLike {
	return v.genericParameters_
}

// Private

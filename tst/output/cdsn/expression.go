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

var expressionClass = &expressionClass_{
	// Initialize class constants.
}

// Function

func Expression() ExpressionClassLike {
	return expressionClass
}

// CLASS METHODS

// Target

type expressionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *expressionClass_) Make(any_ any) ExpressionLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(any_):
		panic("The any_ attribute is required for each Expression.")
	default:
		return &expression_{
			// Initialize instance attributes.
			class_: c,
			any_: any_,
		}
	}
}

// Private

func (c *expressionClass_) isUndefined(value any) bool {
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

type expression_ struct {
	// Define instance attributes.
	class_ ExpressionClassLike
	any_ any
}

// Attributes

func (v *expression_) GetClass() ExpressionClassLike {
	return v.class_
}

func (v *expression_) GetAny() any {
	return v.any_
}

// Private

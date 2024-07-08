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
	ref "reflect"
)

// CLASS ACCESS

// Reference

var aspectClass = &aspectClass_{
	// Initialize class constants.
}

// Function

func Aspect() AspectClassLike {
	return aspectClass
}

// CLASS METHODS

// Target

type aspectClass_ struct {
	// Define class constants.
}

// Constructors

func (c *aspectClass_) Make(
	declaration DeclarationLike,
	methods col.Sequential[MethodLike],
) AspectLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(declaration):
		panic("The declaration attribute is required for each Aspect.")
	case c.isUndefined(methods):
		panic("The methods attribute is required for each Aspect.")
	default:
		return &aspect_{
			// Initialize instance attributes.
			class_: c,
			declaration_: declaration,
			methods_: methods,
		}
	}
}

// Private

func (c *aspectClass_) isUndefined(value any) bool {
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

type aspect_ struct {
	// Define instance attributes.
	class_ AspectClassLike
	declaration_ DeclarationLike
	methods_ col.Sequential[MethodLike]
}

// Attributes

func (v *aspect_) GetClass() AspectClassLike {
	return v.class_
}

func (v *aspect_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *aspect_) GetMethods() col.Sequential[MethodLike] {
	return v.methods_
}

// Private

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

var typeClass = &typeClass_{
	// Initialize class constants.
}

// Function

func Type() TypeClassLike {
	return typeClass
}

// CLASS METHODS

// Target

type typeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *typeClass_) Make(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	enumeration EnumerationLike,
) TypeLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(declaration):
		panic("The declaration attribute is required for each Type.")
	case c.isUndefined(abstraction):
		panic("The abstraction attribute is required for each Type.")
	case c.isUndefined(enumeration):
		panic("The enumeration attribute is required for each Type.")
	default:
		return &type_{
			// Initialize instance attributes.
			class_: c,
			declaration_: declaration,
			abstraction_: abstraction,
			enumeration_: enumeration,
		}
	}
}

// Private

func (c *typeClass_) isUndefined(value any) bool {
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

type type_ struct {
	// Define instance attributes.
	class_ TypeClassLike
	declaration_ DeclarationLike
	abstraction_ AbstractionLike
	enumeration_ EnumerationLike
}

// Attributes

func (v *type_) GetClass() TypeClassLike {
	return v.class_
}

func (v *type_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *type_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *type_) GetOptionalEnumeration() EnumerationLike {
	return v.enumeration_
}

// Private

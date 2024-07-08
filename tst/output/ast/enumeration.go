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

var enumerationClass = &enumerationClass_{
	// Initialize class constants.
}

// Function

func Enumeration() EnumerationClassLike {
	return enumerationClass
}

// CLASS METHODS

// Target

type enumerationClass_ struct {
	// Define class constants.
}

// Constructors

func (c *enumerationClass_) Make(values ValuesLike) EnumerationLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(values):
		panic("The values attribute is required for each Enumeration.")
	default:
		return &enumeration_{
			// Initialize instance attributes.
			class_: c,
			values_: values,
		}
	}
}

// Private

func (c *enumerationClass_) isUndefined(value any) bool {
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

type enumeration_ struct {
	// Define instance attributes.
	class_ EnumerationClassLike
	values_ ValuesLike
}

// Attributes

func (v *enumeration_) GetClass() EnumerationClassLike {
	return v.class_
}

func (v *enumeration_) GetValues() ValuesLike {
	return v.values_
}

// Private

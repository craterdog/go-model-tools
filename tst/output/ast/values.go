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

var valuesClass = &valuesClass_{
	// Initialize class constants.
}

// Function

func Values() ValuesClassLike {
	return valuesClass
}

// CLASS METHODS

// Target

type valuesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *valuesClass_) Make(
	value ValueLike,
	additionalValues col.Sequential[AdditionalValueLike],
) ValuesLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(value):
		panic("The value attribute is required for each Values.")
	case c.isUndefined(additionalValues):
		panic("The additionalValues attribute is required for each Values.")
	default:
		return &values_{
			// Initialize instance attributes.
			class_: c,
			value_: value,
			additionalValues_: additionalValues,
		}
	}
}

// Private

func (c *valuesClass_) isUndefined(value any) bool {
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

type values_ struct {
	// Define instance attributes.
	class_ ValuesClassLike
	value_ ValueLike
	additionalValues_ col.Sequential[AdditionalValueLike]
}

// Attributes

func (v *values_) GetClass() ValuesClassLike {
	return v.class_
}

func (v *values_) GetValue() ValueLike {
	return v.value_
}

func (v *values_) GetAdditionalValues() col.Sequential[AdditionalValueLike] {
	return v.additionalValues_
}

// Private

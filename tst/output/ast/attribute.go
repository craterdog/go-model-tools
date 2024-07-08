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

var attributeClass = &attributeClass_{
	// Initialize class constants.
}

// Function

func Attribute() AttributeClassLike {
	return attributeClass
}

// CLASS METHODS

// Target

type attributeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *attributeClass_) Make(
	name string,
	parameter ParameterLike,
	abstraction AbstractionLike,
) AttributeLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(name):
		panic("The name attribute is required for each Attribute.")
	case c.isUndefined(parameter):
		panic("The parameter attribute is required for each Attribute.")
	case c.isUndefined(abstraction):
		panic("The abstraction attribute is required for each Attribute.")
	default:
		return &attribute_{
			// Initialize instance attributes.
			class_: c,
			name_: name,
			parameter_: parameter,
			abstraction_: abstraction,
		}
	}
}

// Private

func (c *attributeClass_) isUndefined(value any) bool {
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

type attribute_ struct {
	// Define instance attributes.
	class_ AttributeClassLike
	name_ string
	parameter_ ParameterLike
	abstraction_ AbstractionLike
}

// Attributes

func (v *attribute_) GetClass() AttributeClassLike {
	return v.class_
}

func (v *attribute_) GetName() string {
	return v.name_
}

func (v *attribute_) GetOptionalParameter() ParameterLike {
	return v.parameter_
}

func (v *attribute_) GetOptionalAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private

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

var constructorClass = &constructorClass_{
	// Initialize class constants.
}

// Function

func Constructor() ConstructorClassLike {
	return constructorClass
}

// CLASS METHODS

// Target

type constructorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constructorClass_) Make(
	name string,
	parameters ParametersLike,
	abstraction AbstractionLike,
) ConstructorLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(name):
		panic("The name attribute is required for each Constructor.")
	case c.isUndefined(parameters):
		panic("The parameters attribute is required for each Constructor.")
	case c.isUndefined(abstraction):
		panic("The abstraction attribute is required for each Constructor.")
	default:
		return &constructor_{
			// Initialize instance attributes.
			class_: c,
			name_: name,
			parameters_: parameters,
			abstraction_: abstraction,
		}
	}
}

// Private

func (c *constructorClass_) isUndefined(value any) bool {
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

type constructor_ struct {
	// Define instance attributes.
	class_ ConstructorClassLike
	name_ string
	parameters_ ParametersLike
	abstraction_ AbstractionLike
}

// Attributes

func (v *constructor_) GetClass() ConstructorClassLike {
	return v.class_
}

func (v *constructor_) GetName() string {
	return v.name_
}

func (v *constructor_) GetOptionalParameters() ParametersLike {
	return v.parameters_
}

func (v *constructor_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private

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

var instanceClass = &instanceClass_{
	// Initialize class constants.
}

// Function

func Instance() InstanceClassLike {
	return instanceClass
}

// CLASS METHODS

// Target

type instanceClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instanceClass_) Make(
	declaration DeclarationLike,
	attributes AttributesLike,
	abstractions AbstractionsLike,
	methods MethodsLike,
) InstanceLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(declaration):
		panic("The declaration attribute is required for each Instance.")
	case c.isUndefined(attributes):
		panic("The attributes attribute is required for each Instance.")
	case c.isUndefined(abstractions):
		panic("The abstractions attribute is required for each Instance.")
	case c.isUndefined(methods):
		panic("The methods attribute is required for each Instance.")
	default:
		return &instance_{
			// Initialize instance attributes.
			class_: c,
			declaration_: declaration,
			attributes_: attributes,
			abstractions_: abstractions,
			methods_: methods,
		}
	}
}

// Private

func (c *instanceClass_) isUndefined(value any) bool {
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

type instance_ struct {
	// Define instance attributes.
	class_ InstanceClassLike
	declaration_ DeclarationLike
	attributes_ AttributesLike
	abstractions_ AbstractionsLike
	methods_ MethodsLike
}

// Attributes

func (v *instance_) GetClass() InstanceClassLike {
	return v.class_
}

func (v *instance_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instance_) GetAttributes() AttributesLike {
	return v.attributes_
}

func (v *instance_) GetOptionalAbstractions() AbstractionsLike {
	return v.abstractions_
}

func (v *instance_) GetOptionalMethods() MethodsLike {
	return v.methods_
}

// Private

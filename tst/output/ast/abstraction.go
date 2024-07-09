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

// CLASS ACCESS

// Reference

var abstractionClass = &abstractionClass_{
	// Initialize class constants.
}

// Function

func Abstraction() AbstractionClassLike {
	return abstractionClass
}

// CLASS METHODS

// Target

type abstractionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *abstractionClass_) Make(
	prefix PrefixLike,
	alias AliasLike,
	name string,
	genericArguments GenericArgumentsLike,
) AbstractionLike {
	// Validate the arguments.
	switch {
	case isUndefined(prefix):
		panic("The prefix attribute is required for each Abstraction.")
	case isUndefined(alias):
		panic("The alias attribute is required for each Abstraction.")
	case isUndefined(name):
		panic("The name attribute is required for each Abstraction.")
	case isUndefined(genericArguments):
		panic("The genericArguments attribute is required for each Abstraction.")
	default:
		return &abstraction_{
			// Initialize instance attributes.
			class_: c,
			prefix_: prefix,
			alias_: alias,
			name_: name,
			genericArguments_: genericArguments,
		}
	}
}

// INSTANCE METHODS

// Target

type abstraction_ struct {
	// Define instance attributes.
	class_ AbstractionClassLike
	optionalPrefix_ PrefixLike
	optionalAlias_ AliasLike
	name_ string
	optionalGenericArguments_ GenericArgumentsLike
	prefix_ PrefixLike
	alias_ AliasLike
	genericArguments_ GenericArgumentsLike
}

// Attributes

func (v *abstraction_) GetClass() AbstractionClassLike {
	return v.class_
}

func (v *abstraction_) GetOptionalPrefix() PrefixLike {
	return v.optionalPrefix_
}

func (v *abstraction_) GetOptionalAlias() AliasLike {
	return v.optionalAlias_
}

func (v *abstraction_) GetName() string {
	return v.name_
}

func (v *abstraction_) GetOptionalGenericArguments() GenericArgumentsLike {
	return v.optionalGenericArguments_
}

// Private

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

var aliasClass = &aliasClass_{
	// Initialize class constants.
}

// Function

func Alias() AliasClassLike {
	return aliasClass
}

// CLASS METHODS

// Target

type aliasClass_ struct {
	// Define class constants.
}

// Constructors

func (c *aliasClass_) Make(name string) AliasLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(name):
		panic("The name attribute is required for each Alias.")
	default:
		return &alias_{
			// Initialize instance attributes.
			class_: c,
			name_: name,
		}
	}
}

// Private

func (c *aliasClass_) isUndefined(value any) bool {
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

type alias_ struct {
	// Define instance attributes.
	class_ AliasClassLike
	name_ string
}

// Attributes

func (v *alias_) GetClass() AliasClassLike {
	return v.class_
}

func (v *alias_) GetName() string {
	return v.name_
}

// Private

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

var mapClass = &mapClass_{
	// Initialize class constants.
}

// Function

func Map() MapClassLike {
	return mapClass
}

// CLASS METHODS

// Target

type mapClass_ struct {
	// Define class constants.
}

// Constructors

func (c *mapClass_) Make(name string) MapLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(name):
		panic("The name attribute is required for each Map.")
	default:
		return &map_{
			// Initialize instance attributes.
			class_: c,
			name_: name,
		}
	}
}

// Private

func (c *mapClass_) isUndefined(value any) bool {
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

type map_ struct {
	// Define instance attributes.
	class_ MapClassLike
	name_ string
}

// Attributes

func (v *map_) GetClass() MapClassLike {
	return v.class_
}

func (v *map_) GetName() string {
	return v.name_
}

// Private

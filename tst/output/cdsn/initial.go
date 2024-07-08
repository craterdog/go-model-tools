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

var initialClass = &initialClass_{
	// Initialize class constants.
}

// Function

func Initial() InitialClassLike {
	return initialClass
}

// CLASS METHODS

// Target

type initialClass_ struct {
	// Define class constants.
}

// Constructors

func (c *initialClass_) Make(rune_ string) InitialLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(rune_):
		panic("The rune_ attribute is required for each Initial.")
	default:
		return &initial_{
			// Initialize instance attributes.
			class_: c,
			rune_: rune_,
		}
	}
}

// Private

func (c *initialClass_) isUndefined(value any) bool {
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

type initial_ struct {
	// Define instance attributes.
	class_ InitialClassLike
	rune_ string
}

// Attributes

func (v *initial_) GetClass() InitialClassLike {
	return v.class_
}

func (v *initial_) GetRune() string {
	return v.rune_
}

// Private

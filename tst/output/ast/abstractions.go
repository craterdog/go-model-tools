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

var abstractionsClass = &abstractionsClass_{
	// Initialize class constants.
}

// Function

func Abstractions() AbstractionsClassLike {
	return abstractionsClass
}

// CLASS METHODS

// Target

type abstractionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *abstractionsClass_) Make(
	note string,
	abstractions col.Sequential[AbstractionLike],
) AbstractionsLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(note):
		panic("The note attribute is required for each Abstractions.")
	case c.isUndefined(abstractions):
		panic("The abstractions attribute is required for each Abstractions.")
	default:
		return &abstractions_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			abstractions_: abstractions,
		}
	}
}

// Private

func (c *abstractionsClass_) isUndefined(value any) bool {
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

type abstractions_ struct {
	// Define instance attributes.
	class_ AbstractionsClassLike
	note_ string
	abstractions_ col.Sequential[AbstractionLike]
}

// Attributes

func (v *abstractions_) GetClass() AbstractionsClassLike {
	return v.class_
}

func (v *abstractions_) GetNote() string {
	return v.note_
}

func (v *abstractions_) GetAbstractions() col.Sequential[AbstractionLike] {
	return v.abstractions_
}

// Private

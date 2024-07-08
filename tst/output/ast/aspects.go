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

var aspectsClass = &aspectsClass_{
	// Initialize class constants.
}

// Function

func Aspects() AspectsClassLike {
	return aspectsClass
}

// CLASS METHODS

// Target

type aspectsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *aspectsClass_) Make(
	note string,
	aspects col.Sequential[AspectLike],
) AspectsLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(note):
		panic("The note attribute is required for each Aspects.")
	case c.isUndefined(aspects):
		panic("The aspects attribute is required for each Aspects.")
	default:
		return &aspects_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			aspects_: aspects,
		}
	}
}

// Private

func (c *aspectsClass_) isUndefined(value any) bool {
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

type aspects_ struct {
	// Define instance attributes.
	class_ AspectsClassLike
	note_ string
	aspects_ col.Sequential[AspectLike]
}

// Attributes

func (v *aspects_) GetClass() AspectsClassLike {
	return v.class_
}

func (v *aspects_) GetNote() string {
	return v.note_
}

func (v *aspects_) GetAspects() col.Sequential[AspectLike] {
	return v.aspects_
}

// Private

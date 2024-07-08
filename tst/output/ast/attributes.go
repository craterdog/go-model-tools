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

var attributesClass = &attributesClass_{
	// Initialize class constants.
}

// Function

func Attributes() AttributesClassLike {
	return attributesClass
}

// CLASS METHODS

// Target

type attributesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *attributesClass_) Make(
	note string,
	attributes col.Sequential[AttributeLike],
) AttributesLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(note):
		panic("The note attribute is required for each Attributes.")
	case c.isUndefined(attributes):
		panic("The attributes attribute is required for each Attributes.")
	default:
		return &attributes_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			attributes_: attributes,
		}
	}
}

// Private

func (c *attributesClass_) isUndefined(value any) bool {
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

type attributes_ struct {
	// Define instance attributes.
	class_ AttributesClassLike
	note_ string
	attributes_ col.Sequential[AttributeLike]
}

// Attributes

func (v *attributes_) GetClass() AttributesClassLike {
	return v.class_
}

func (v *attributes_) GetNote() string {
	return v.note_
}

func (v *attributes_) GetAttributes() col.Sequential[AttributeLike] {
	return v.attributes_
}

// Private

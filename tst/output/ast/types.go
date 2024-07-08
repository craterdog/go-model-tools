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

var typesClass = &typesClass_{
	// Initialize class constants.
}

// Function

func Types() TypesClassLike {
	return typesClass
}

// CLASS METHODS

// Target

type typesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *typesClass_) Make(
	note string,
	types col.Sequential[TypeLike],
) TypesLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(note):
		panic("The note attribute is required for each Types.")
	case c.isUndefined(types):
		panic("The types attribute is required for each Types.")
	default:
		return &types_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			types_: types,
		}
	}
}

// Private

func (c *typesClass_) isUndefined(value any) bool {
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

type types_ struct {
	// Define instance attributes.
	class_ TypesClassLike
	note_ string
	types_ col.Sequential[TypeLike]
}

// Attributes

func (v *types_) GetClass() TypesClassLike {
	return v.class_
}

func (v *types_) GetNote() string {
	return v.note_
}

func (v *types_) GetTypes() col.Sequential[TypeLike] {
	return v.types_
}

// Private

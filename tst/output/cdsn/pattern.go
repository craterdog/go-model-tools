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

var patternClass = &patternClass_{
	// Initialize class constants.
}

// Function

func Pattern() PatternClassLike {
	return patternClass
}

// CLASS METHODS

// Target

type patternClass_ struct {
	// Define class constants.
}

// Constructors

func (c *patternClass_) Make(
	parts col.ListLike[PartLike],
	alternatives col.ListLike[AlternativeLike],
) PatternLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(parts):
		panic("The parts attribute is required for each Pattern.")
	case c.isUndefined(alternatives):
		panic("The alternatives attribute is required for each Pattern.")
	default:
		return &pattern_{
			// Initialize instance attributes.
			class_: c,
			parts_: parts,
			alternatives_: alternatives,
		}
	}
}

// Private

func (c *patternClass_) isUndefined(value any) bool {
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

type pattern_ struct {
	// Define instance attributes.
	class_ PatternClassLike
	parts_ col.ListLike[PartLike]
	alternatives_ col.ListLike[AlternativeLike]
}

// Attributes

func (v *pattern_) GetClass() PatternClassLike {
	return v.class_
}

func (v *pattern_) GetParts() col.ListLike[PartLike] {
	return v.parts_
}

func (v *pattern_) GetAlternatives() col.ListLike[AlternativeLike] {
	return v.alternatives_
}

// Private

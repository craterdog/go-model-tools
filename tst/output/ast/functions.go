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

var functionsClass = &functionsClass_{
	// Initialize class constants.
}

// Function

func Functions() FunctionsClassLike {
	return functionsClass
}

// CLASS METHODS

// Target

type functionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *functionsClass_) Make(
	note string,
	functions col.Sequential[FunctionLike],
) FunctionsLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(note):
		panic("The note attribute is required for each Functions.")
	case c.isUndefined(functions):
		panic("The functions attribute is required for each Functions.")
	default:
		return &functions_{
			// Initialize instance attributes.
			class_: c,
			note_: note,
			functions_: functions,
		}
	}
}

// Private

func (c *functionsClass_) isUndefined(value any) bool {
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

type functions_ struct {
	// Define instance attributes.
	class_ FunctionsClassLike
	note_ string
	functions_ col.Sequential[FunctionLike]
}

// Attributes

func (v *functions_) GetClass() FunctionsClassLike {
	return v.class_
}

func (v *functions_) GetNote() string {
	return v.note_
}

func (v *functions_) GetFunctions() col.Sequential[FunctionLike] {
	return v.functions_
}

// Private

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

var classClass = &classClass_{
	// Initialize class constants.
}

// Function

func Class() ClassClassLike {
	return classClass
}

// CLASS METHODS

// Target

type classClass_ struct {
	// Define class constants.
}

// Constructors

func (c *classClass_) Make(
	declaration DeclarationLike,
	constructors ConstructorsLike,
	constants ConstantsLike,
	functions FunctionsLike,
) ClassLike {
	// Validate the arguments.
	switch {
	case c.isUndefined(declaration):
		panic("The declaration attribute is required for each Class.")
	case c.isUndefined(constructors):
		panic("The constructors attribute is required for each Class.")
	case c.isUndefined(constants):
		panic("The constants attribute is required for each Class.")
	case c.isUndefined(functions):
		panic("The functions attribute is required for each Class.")
	default:
		return &class_{
			// Initialize instance attributes.
			class_: c,
			declaration_: declaration,
			constructors_: constructors,
			constants_: constants,
			functions_: functions,
		}
	}
}

// Private

func (c *classClass_) isUndefined(value any) bool {
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

type class_ struct {
	// Define instance attributes.
	class_ ClassClassLike
	declaration_ DeclarationLike
	constructors_ ConstructorsLike
	constants_ ConstantsLike
	functions_ FunctionsLike
}

// Attributes

func (v *class_) GetClass() ClassClassLike {
	return v.class_
}

func (v *class_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *class_) GetConstructors() ConstructorsLike {
	return v.constructors_
}

func (v *class_) GetOptionalConstants() ConstantsLike {
	return v.constants_
}

func (v *class_) GetOptionalFunctions() FunctionsLike {
	return v.functions_
}

// Private

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

func (c *classClass_) MakeWithAttributes(
	declaration DeclarationLike,
	constants col.ListLike[ConstantLike],
	constructors col.ListLike[ConstructorLike],
	functions col.ListLike[FunctionLike],
) ClassLike {
	return &class_{
		// Initialize instance attributes.
		class_: c,
		declaration_: declaration,
		constants_: constants,
		constructors_: constructors,
		functions_: functions,
	}
}

// INSTANCE METHODS

// Target

type class_ struct {
	// Define instance attributes.
	class_ ClassClassLike
	declaration_ DeclarationLike
	constants_ col.ListLike[ConstantLike]
	constructors_ col.ListLike[ConstructorLike]
	functions_ col.ListLike[FunctionLike]
}

// Attributes

func (v *class_) GetClass() ClassClassLike {
	return v.class_
}

func (v *class_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *class_) GetConstants() col.ListLike[ConstantLike] {
	return v.constants_
}

func (v *class_) GetConstructors() col.ListLike[ConstructorLike] {
	return v.constructors_
}

func (v *class_) GetFunctions() col.ListLike[FunctionLike] {
	return v.functions_
}

// Private

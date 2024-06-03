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

import ()

// CLASS ACCESS

// Reference

var definitionClass = &definitionClass_{
	// Initialize class constants.
}

// Function

func Definition() DefinitionClassLike {
	return definitionClass
}

// CLASS METHODS

// Target

type definitionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *definitionClass_) MakeWithAttributes(
	comment string,
	name string,
	expression ExpressionLike,
) DefinitionLike {
	return &definition_{
		// Initialize instance attributes.
		class_: c,
		comment_: comment,
		name_: name,
		expression_: expression,
	}
}

// INSTANCE METHODS

// Target

type definition_ struct {
	// Define instance attributes.
	class_ DefinitionClassLike
	comment_ string
	name_ string
	expression_ ExpressionLike
}

// Attributes

func (v *definition_) GetClass() DefinitionClassLike {
	return v.class_
}

func (v *definition_) GetComment() string {
	return v.comment_
}

func (v *definition_) GetName() string {
	return v.name_
}

func (v *definition_) GetExpression() ExpressionLike {
	return v.expression_
}

// Private

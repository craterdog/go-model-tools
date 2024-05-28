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

package agent

import (
	ast "github.com/craterdog/go-model-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// Any private class constants should be initialized here.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{}
}

// Functions

// INSTANCE METHODS

// Target

type generator_ struct {
	class_ GeneratorClassLike
}

// Attributes

func (v *generator_) GetClass() GeneratorClassLike {
	return v.class_
}

// Public

func (v *generator_) CreateModel(
	name string,
	copyright string,
) ast.ModelLike {
	var result_ ast.ModelLike
	// TBA - Implement the method.
	return result_
}

func (v *generator_) GenerateClass(
	model ast.ModelLike,
	name string,
) string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Private

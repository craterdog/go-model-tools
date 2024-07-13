/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package complex

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var complexClass = &complexClass_{
	// Initialize class constants.
}

// Function

func Complex() ComplexClassLike {
	return complexClass
}

// CLASS METHODS

// Target

type complexClass_ struct {
	// Define class constants.
	zero_ ComplexLike
	infinity_ ComplexLike
}

// Constructors

func (c *complexClass_) Make(
	realPart float64,
	imaginaryPart float64,
	form Form,
) ComplexLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(realPart):
		panic("The realPart attribute is required for each Complex.")
	case col.IsUndefined(imaginaryPart):
		panic("The imaginaryPart attribute is required for each Complex.")
	case col.IsUndefined(form):
		panic("The form attribute is required for each Complex.")
	default:
		return &complex_{
			// Initialize instance attributes.
			class_: c,
			realPart_: realPart,
			imaginaryPart_: imaginaryPart,
			form_: form,
		}
	}
}

func (c *complexClass_) MakeFromValue(value complex128) ComplexLike {
	// Validate the arguments.
	switch {
	default:
		return &complex_{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// Constants

func (c *complexClass_) Zero() ComplexLike {
	return c.zero_
}

func (c *complexClass_) Infinity() ComplexLike {
	return c.infinity_
}

// Functions

func (c *complexClass_) Inverse(value ComplexLike) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Sum(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Difference(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Reciprocal(value ComplexLike) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Product(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Quotient(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Norm(
	function NormFunction[ComplexLike],
	value ComplexLike,
) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type complex_ struct {
	// Define instance attributes.
	class_ ComplexClassLike
	realPart_ float64
	imaginaryPart_ float64
	form_ Form
}

// Attributes

func (v *complex_) GetClass() ComplexClassLike {
	return v.class_
}

func (v *complex_) GetRealPart() float64 {
	return v.realPart_
}

func (v *complex_) GetImaginaryPart() float64 {
	return v.imaginaryPart_
}

func (v *complex_) GetForm() Form {
	return v.form_
}

func (v *complex_) SetForm(form Form) {
	if col.IsUndefined(form) {
		panic("The form attribute cannot be nil.")
	}
	v.form_ = form
}

// Continuous

func (v *complex_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *complex_) IsDiscrete() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *complex_) IsInfinity() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *complex_) IsReal() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *complex_) IsImaginary() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Private

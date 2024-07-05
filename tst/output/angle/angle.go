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

package angle

// CLASS ACCESS

// Reference

var angleClass = &angleClass_{
	// Initialize class constants.
}

// Function

func Angle() AngleClassLike {
	return angleClass
}

// CLASS METHODS

// Target

type angleClass_ struct {
	// Define class constants.
	pi_ AngleLike
	tau_ AngleLike
}

// Constructors

func (c *angleClass_) MakeFromValue(value float64) AngleLike {
	// TBA - Validate the value.
	return angle_(value)
}

func (c *angleClass_) MakeFromString(value string) AngleLike {
	var result_ AngleLike
	// TBA - Implement the method.
	return result_
}

// Constants

func (c *angleClass_) Pi() AngleLike {
	return c.pi_
}

func (c *angleClass_) Tau() AngleLike {
	return c.tau_
}

// Functions

func (c *angleClass_) Apply(
	function TrigonometricFunction,
	angle AngleLike,
) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

func (c *angleClass_) Sine(angle AngleLike) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

func (c *angleClass_) Cosine(angle AngleLike) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

func (c *angleClass_) Tangent(angle AngleLike) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type angle_ float64

// Attributes

func (v angle_) GetClass() AngleClassLike {
	return Angle()
}

// Angular

func (v angle_) AsNormalized() AngleLike {
	var result_ AngleLike
	// TBA - Implement the method.
	return result_
}

func (v angle_) InUnits(units Units) float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

// Public

func (v angle_) AsFloat() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v angle_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Private

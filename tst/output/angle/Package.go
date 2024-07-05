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

/*
Package "angle" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package angle

// Types

/*
Units is a constrained type representing the possible units for an angle.
*/
type Units uint8

const (
	Degrees Units = iota
	Radians
	Gradians
)

// Functionals

/*
TrigonometricFunction is a functional type that defines the signature for any
trigonometric function.
*/
type TrigonometricFunction func(angle AngleLike) float64

// Classes

/*
AngleClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each angle-like concrete
class.
*/
type AngleClassLike interface {
	// Constructors
	MakeFromValue(value float64) AngleLike
	MakeFromString(value string) AngleLike

	// Constants
	Pi() AngleLike
	Tau() AngleLike

	// Functions
	Apply(
		function TrigonometricFunction,
		angle AngleLike,
	) float64
	Sine(angle AngleLike) float64
	Cosine(angle AngleLike) float64
	Tangent(angle AngleLike) float64
}

// Instances

/*
AngleLike is an instance interface that defines the complete set of attributes,
abstractions and methods that must be supported by each instance of a concrete
angle-like class.
*/
type AngleLike interface {
	// Attributes
	GetClass() AngleClassLike

	// Abstractions
	Angular

	// Methods
	AsFloat() float64
	IsZero() bool
}

// Aspects

/*
Angular is an aspect interface that defines a set of method signatures that
must be supported by each instance of an angular concrete class.
*/
type Angular interface {
	AsNormalized() AngleLike
	InUnits(units Units) float64
}

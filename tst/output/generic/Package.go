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
Package "generic" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package generic

// Types

/*
Rank is a constrained type representing the possible rankings for two values.
*/
type Rank uint8

const (
	LesserRank Rank = iota
	EqualRank
	GreaterRank
)

// Functionals

/*
RankingFunction[V any] is a functional type that defines the signature for any
function that can determine the relative ranking of two values.
*/
type RankingFunction[V any] func(
	first V,
	second V,
) Rank

// Aspects

/*
Sequential[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of a sequential concrete class.
*/
type Sequential[V any] interface {
	// Methods
	AsArray() []V
	GetSize() int
	IsEmpty() bool
}

// Classes

/*
SetClassLike[V any] is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete set-like class.

The following functions are supported:

And() returns a new set containing the values that are both of the specified
sets.

Or() returns a new set containing the values that are in either of the specified
sets.

Sans() returns a new set containing the values that are in the first specified
set but not in the second specified set.

Xor() returns a new set containing the values that are in the first specified
set or the second specified set but not both.
*/
type SetClassLike[V any] interface {
	// Constants
	DefaultRanker() RankingFunction[V]

	// Constructors
	Make() SetLike[V]
	MakeFromArray(values []V) SetLike[V]
	MakeFromSequence(values Sequential[V]) SetLike[V]
	MakeWithRanker(ranker RankingFunction[V]) SetLike[V]

	// Functions
	And(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Or(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Sans(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Xor(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
}

// Instances

/*
SetLike[V any] is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete set-like class.  A set-like class maintains an ordered
sequence of values which can grow or shrink as needed.

This type is parameterized as follows:
  - V is any type of value.

The order of the values is determined by a configurable ranking function.
*/
type SetLike[V any] interface {
	// Attributes
	GetClass() SetClassLike[V]
	SetRanker(ranker RankingFunction[V])

	// Abstractions
	Sequential[V]

	// Methods
	AddValue(value V)
	RemoveValue(value V)
	RemoveAll()
}

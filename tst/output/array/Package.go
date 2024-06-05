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
Package "array" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package array

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
Accessible[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of an accessible concrete class.
The values in an accessible class are accessed using indices. The indices of an
accessible class are ORDINAL rather than ZERO based—which never really made
sense except for pointer offsets.

This approach allows for positive indices starting at the beginning of the
sequence as follows:

	    1           2           3             N
	[value 1] . [value 2] . [value 3] ... [value N]

Notice that because the indices are ordinal based, the index actually matches
the position in the sequence.
*/
type Accessible[V any] interface {
	// Methods
	GetValue(index uint) V
	GetValues(
		first uint,
		last uint,
	) Sequential[V]
}

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

/*
Updatable[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of an updatable concrete class.
*/
type Updatable[V any] interface {
	// Methods
	SetValue(
		index uint,
		value V,
	)
	SetValues(
		index uint,
		values Sequential[V],
	)
}

// Classes

/*
ArrayClassLike[V any] is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike[V any] interface {
	// Constants
	DefaultRanker() RankingFunction[V]

	// Constructors
	MakeFromValue(value []V) ArrayLike[V]
	MakeFromSequence(values Sequential[V]) ArrayLike[V]
	MakeFromSize(size int) ArrayLike[V]
}

// Instances

/*
ArrayLike[V any] is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete array-like class.

An array-like class maintains a fixed length indexed sequence of values.  Each
value is associated with an implicit positive integer index. An array-like class
uses ORDINAL based indexing rather than the more common—and nonsensical—ZERO
based indexing scheme.
*/
type ArrayLike[V any] interface {
	// Attributes
	GetClass() ArrayClassLike[V]

	// Abstractions
	Accessible[V]
	Sequential[V]
	Updatable[V]

	// Methods
	SortValues()
	SortValuesWithRanker(ranker RankingFunction[V])
}

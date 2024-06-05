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
Package "catalog" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package catalog

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
Associative[K comparable, V any] defines the set of method signatures that
must be supported by all sequences of key-value associations.
*/
type Associative[K comparable, V any] interface {
	// Methods
	GetKeys() Sequential[K]
	GetValue(key K) V
	RemoveValue(key K) V
	SetValue(
		key K,
		value V,
	)
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

// Classes

/*
AssociationClassLike[K comparable, V any] is a class interface that defines
the complete set of class constants, constructors and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike[K comparable, V any] interface {
	// Constructors
	MakeWithAttributes(
		key K,
		value V,
	) AssociationLike[K, V]
}

/*
CatalogClassLike[K comparable, V any] is a class interface that defines the
complete set of class constants, constructors and functions that must be
supported by each concrete catalog-like class.

The following functions are supported:

Extract() returns a new catalog containing only the associations that are in
the specified catalog that have the specified keys.  The associations in the
resulting catalog will be in the same order as the specified keys.

Merge() returns a new catalog containing all of the associations that are in
the specified catalogs in the order that they appear in each catalog.  If a
key is present in both catalogs, the value of the key from the second
catalog takes precedence.
*/
type CatalogClassLike[K comparable, V any] interface {
	// Constants
	DefaultRanker() RankingFunction[AssociationLike[K, V]]

	// Constructors
	Make() CatalogLike[K, V]
	MakeFromArray(associations []AssociationLike[K, V]) CatalogLike[K, V]
	MakeFromMap(associations map[K]V) CatalogLike[K, V]
	MakeFromSequence(associations Sequential[AssociationLike[K, V]]) CatalogLike[K, V]

	// Functions
	Extract(
		catalog CatalogLike[K, V],
		keys Sequential[K],
	) CatalogLike[K, V]
	Merge(
		first CatalogLike[K, V],
		second CatalogLike[K, V],
	) CatalogLike[K, V]
}

// Instances

/*
AssociationLike[K comparable, V any] is an instance interface that defines the
complete set of instance attributes, abstractions and methods that must be
supported by each instance of a concrete association-like class.
*/
type AssociationLike[K comparable, V any] interface {
	// Attributes
	GetClass() AssociationClassLike[K, V]
	GetKey() K
	GetValue() V
	SetValue(value V)
}

/*
CatalogLike[K comparable, V any] is an instance interface that defines the
complete set of instance attributes, abstractions and methods that must be
supported by each instance of a concrete catalog-like class.
*/
type CatalogLike[K comparable, V any] interface {
	// Attributes
	GetClass() CatalogClassLike[K, V]

	// Abstractions
	Associative[K, V]
	Sequential[AssociationLike[K, V]]

	// Methods
	SortValues()
	SortValuesWithRanker(ranker RankingFunction[AssociationLike[K, V]])
}

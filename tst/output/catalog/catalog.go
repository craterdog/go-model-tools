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

package catalog

import (
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var catalogClass = map[string]any{}
var catalogMutex syn.Mutex

// Function

func Catalog[
	K comparable,
	V any,
]() CatalogClassLike[K, V] {
	// Generate the name of the bound class type.
	var class *catalogClass_[K, V]
	var name = fmt.Sprintf("%T", class)

	// Check for existing bound class type.
	catalogMutex.Lock()
	var value = catalogClass[name]
	switch actual := value.(type) {
	case *catalogClass_[K, V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &catalogClass_[K, V]{
			// Initialize class constants.
		}
		catalogClass[name] = class
	}
	catalogMutex.Unlock()

	// Return a reference to the bound class type.
	return class
}

// CLASS METHODS

// Target

type catalogClass_[
	K comparable,
	V any,
] struct {
	// Define class constants.
	defaultRanker_ RankingFunction[AssociationLike[K, V]]
}

// Constructors

func (c *catalogClass_[K, V]) Make() CatalogLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &catalog_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *catalogClass_[K, V]) MakeFromArray(associations []AssociationLike[K, V]) CatalogLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &catalog_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *catalogClass_[K, V]) MakeFromMap(associations map[K]V) CatalogLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &catalog_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *catalogClass_[K, V]) MakeFromSequence(associations Sequential[AssociationLike[K, V]]) CatalogLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &catalog_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// Constants

func (c *catalogClass_[K, V]) DefaultRanker() RankingFunction[AssociationLike[K, V]] {
	return c.defaultRanker_
}

// Functions

func (c *catalogClass_[K, V]) Extract(
	catalog CatalogLike[K, V],
	keys Sequential[K],
) CatalogLike[K, V] {
	var result_ CatalogLike[K, V]
	// TBA - Implement the function.
	return result_
}

func (c *catalogClass_[K, V]) Merge(
	first CatalogLike[K, V],
	second CatalogLike[K, V],
) CatalogLike[K, V] {
	var result_ CatalogLike[K, V]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type catalog_[
	K comparable,
	V any,
] struct {
	// Define instance attributes.
	class_ CatalogClassLike[K, V]
}

// Attributes

func (v *catalog_[K, V]) GetClass() CatalogClassLike[K, V] {
	return v.class_
}

// Associative[K, V]

func (v *catalog_[K, V]) GetKeys() Sequential[K] {
	var result_ Sequential[K]
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) GetValue(key K) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) RemoveValue(key K) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) SetValue(
	key K,
	value V,
) {
	// TBA - Implement the method.
}

// Sequential[AssociationLike[K, V]]

func (v *catalog_[K, V]) AsArray() []AssociationLike[K, V] {
	var result_ []AssociationLike[K, V]
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *catalog_[K, V]) SortValues() {
	// TBA - Implement the method.
}

func (v *catalog_[K, V]) SortValuesWithRanker(ranker RankingFunction[AssociationLike[K, V]]) {
	// TBA - Implement the method.
}

// Private

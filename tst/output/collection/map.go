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

package collection

import (
	age "github.com/craterdog/go-collection-framework/v4/agent"
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var mapClass = map[string]any{}
var mapMutex syn.Mutex

// Function

func Map[
	K comparable,
	V any,
]() MapClassLike[K, V] {
	// Generate the name of the bound class type.
	var class *mapClass_[K, V]
	var name = fmt.Sprintf("%T", class)

	// Check for existing bound class type.
	mapMutex.Lock()
	var value = mapClass[name]
	switch actual := value.(type) {
	case *mapClass_[K, V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &mapClass_[K, V]{
			// Initialize class constants.
		}
		mapClass[name] = class
	}
	mapMutex.Unlock()

	// Return a reference to the bound class type.
	return class
}

// CLASS METHODS

// Target

type mapClass_[
	K comparable,
	V any,
] struct {
	// Define class constants.
	notation_ NotationLike
}

// Constructors

func (c *mapClass_[K, V]) Make() MapLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &map_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *mapClass_[K, V]) MakeFromArray(associations []AssociationLike[K, V]) MapLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &map_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *mapClass_[K, V]) MakeFromMap(associations map[K]V) MapLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &map_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *mapClass_[K, V]) MakeFromSequence(associations Sequential[AssociationLike[K, V]]) MapLike[K, V] {
	// Validate the arguments.
	switch {
	default:
		return &map_[K, V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// Constants

func (c *mapClass_[K, V]) Notation() NotationLike {
	return c.notation_
}

// INSTANCE METHODS

// Target

type map_[
	K comparable,
	V any,
] struct {
	// Define instance attributes.
	class_ MapClassLike[K, V]
}

// Attributes

func (v *map_[K, V]) GetClass() MapClassLike[K, V] {
	return v.class_
}

// Associative[K, V]

func (v *map_[K, V]) GetValue(key K) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) SetValue(
	key K,
	value V,
) {
	// TBA - Implement the method.
}

func (v *map_[K, V]) GetKeys() Sequential[K] {
	var result_ Sequential[K]
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) GetValues(keys Sequential[K]) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) RemoveValue(key K) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) RemoveValues(keys Sequential[K]) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) RemoveAll() {
	// TBA - Implement the method.
}

// Sequential[AssociationLike[K, V]]

func (v *map_[K, V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) AsArray() []AssociationLike[K, V] {
	var result_ []AssociationLike[K, V]
	// TBA - Implement the method.
	return result_
}

func (v *map_[K, V]) GetIterator() age.IteratorLike[AssociationLike[K, V]] {
	var result_ age.IteratorLike[AssociationLike[K, V]]
	// TBA - Implement the method.
	return result_
}

// Private

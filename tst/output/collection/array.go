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
	col "github.com/craterdog/go-collection-framework/v4"
	syn "sync"
)

// CLASS ACCESS

// Reference

var arrayClass = map[string]any{}
var arrayMutex syn.Mutex

// Function

func Array[V any]() ArrayClassLike[V] {
	// Generate the name of the bound class type.
	var class *arrayClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for existing bound class type.
	arrayMutex.Lock()
	var value = arrayClass[name]
	switch actual := value.(type) {
	case *arrayClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &arrayClass_[V]{
			// Initialize class constants.
		}
		arrayClass[name] = class
	}
	arrayMutex.Unlock()

	// Return a reference to the bound class type.
	return class
}

// CLASS METHODS

// Target

type arrayClass_[V any] struct {
	// Define class constants.
	notation_ NotationLike
}

// Constructors

func (c *arrayClass_[V]) MakeWithSize(size uint) ArrayLike[V] {
	// Validate the arguments.
	switch {
	case col.IsUndefined(size):
		panic("The size attribute is required for each Array.")
	default:
		return &array_[V]{
			// Initialize instance attributes.
			class_: c,
			size_: size,
		}
	}
}

func (c *arrayClass_[V]) MakeFromArray(values []V) ArrayLike[V] {
	// Validate the arguments.
	switch {
	default:
		return &array_[V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *arrayClass_[V]) MakeFromSequence(values Sequential[V]) ArrayLike[V] {
	// Validate the arguments.
	switch {
	default:
		return &array_[V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// Constants

func (c *arrayClass_[V]) Notation() NotationLike {
	return c.notation_
}

// INSTANCE METHODS

// Target

type array_[V any] struct {
	// Define instance attributes.
	class_ ArrayClassLike[V]
	size_ uint
}

// Attributes

func (v *array_[V]) GetClass() ArrayClassLike[V] {
	return v.class_
}

// Accessible[V]

func (v *array_[V]) GetValue(index int) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *array_[V]) GetValues(
	first int,
	last int,
) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

// Sequential[V]

func (v *array_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *array_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *array_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *array_[V]) GetIterator() age.IteratorLike[V] {
	var result_ age.IteratorLike[V]
	// TBA - Implement the method.
	return result_
}

// Sortable[V]

func (v *array_[V]) SortValues() {
	// TBA - Implement the method.
}

func (v *array_[V]) SortValuesWithRanker(ranker age.RankingFunction[V]) {
	// TBA - Implement the method.
}

func (v *array_[V]) ReverseValues() {
	// TBA - Implement the method.
}

func (v *array_[V]) ShuffleValues() {
	// TBA - Implement the method.
}

// Updatable[V]

func (v *array_[V]) SetValue(
	index int,
	value V,
) {
	// TBA - Implement the method.
}

func (v *array_[V]) SetValues(
	index int,
	values Sequential[V],
) {
	// TBA - Implement the method.
}

// Private

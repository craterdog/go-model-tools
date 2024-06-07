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

package array

import (
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var arrayClass = map[string]any{}
var arrayMutex syn.Mutex

// Function

func Array[V any]() ArrayClassLike[V] {
	// Generate the name of the bound class type.
	var result_ ArrayClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	arrayMutex.Lock()
	var value = arrayClass[name]
	switch actual := value.(type) {
	case *arrayClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &arrayClass_[V]{
			// Initialize class constants.
		}
		arrayClass[name] = result_
	}
	arrayMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type arrayClass_[V any] struct {
	// Define class constants.
	defaultRanker_ RankingFunction[V]
}

// Constants

func (c *arrayClass_[V]) DefaultRanker() RankingFunction[V] {
	return c.defaultRanker_
}

// Constructors

func (c *arrayClass_[V]) MakeWithSize(size uint) ArrayLike[V] {
	var result_ ArrayLike[V]
	// TBA - Implement the method.
	return result_
}

func (c *arrayClass_[V]) MakeFromValue(value []V) ArrayLike[V] {
	// TBA - Validate the value.
	return array_[V](value)
}

func (c *arrayClass_[V]) MakeFromSequence(values Sequential[V]) ArrayLike[V] {
	var result_ ArrayLike[V]
	// TBA - Implement the method.
	return result_
}

// INSTANCE METHODS

// Target

type array_[V any] []V

// Attributes

func (v array_[V]) GetClass() ArrayClassLike[V] {
	return Array[V]()
}

// Accessible[V]

func (v array_[V]) GetValue(index int) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v array_[V]) GetValues(
	first int,
	last int,
) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

// Sequential[V]

func (v array_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v array_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v array_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

// Updatable[V]

func (v array_[V]) SetValue(
	index int,
	value V,
) {
	// TBA - Implement the method.
}

func (v array_[V]) SetValues(
	index int,
	values Sequential[V],
) {
	// TBA - Implement the method.
}

// Public

func (v array_[V]) SortValues() {
	// TBA - Implement the method.
}

func (v array_[V]) SortValuesWithRanker(ranker RankingFunction[V]) {
	// TBA - Implement the method.
}

// Private

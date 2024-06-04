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

package generic

import (
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var setClass = map[string]any{}
var setMutex syn.Mutex

// Function

func Set[V any]() SetClassLike[V] {
	// Generate the name of the bound class type.
	var result_ SetClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	setMutex.Lock()
	var value = setClass[name]
	switch actual := value.(type) {
	case *setClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &setClass_[V]{
			// Initialize class constants.
		}
		setClass[name] = result_
	}
	setMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type setClass_[V any] struct {
	// Define class constants.
	defaultRanker_ RankingFunction[V]
}

// Constants

func (c *setClass_[V]) DefaultRanker() RankingFunction[V] {
	return c.defaultRanker_
}

// Constructors

func (c *setClass_[V]) Make() SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *setClass_[V]) MakeFromArray(values []V) SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *setClass_[V]) MakeFromSequence(values Sequential[V]) SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *setClass_[V]) MakeWithRanker(ranker RankingFunction[V]) SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
		ranker_: ranker,
	}
}

// Functions

func (c *setClass_[V]) And(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Or(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Sans(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Xor(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type set_[V any] struct {
	// Define instance attributes.
	class_ SetClassLike[V]
	ranker_ RankingFunction[V]
}

// Attributes

func (v *set_[V]) GetClass() SetClassLike[V] {
	return v.class_
}

func (v *set_[V]) SetRanker(ranker RankingFunction[V]) {
	v.ranker_ = ranker
}

// Sequential[V]

func (v *set_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *set_[V]) AddValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveAll() {
	// TBA - Implement the method.
}

// Private

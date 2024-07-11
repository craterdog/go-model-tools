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
	mod "github.com/craterdog/go-collection-framework/v4"
	syn "sync"
)

// CLASS ACCESS

// Reference

var stackClass = map[string]any{}
var stackMutex syn.Mutex

// Function

func Stack[V any]() StackClassLike[V] {
	// Generate the name of the bound class type.
	var class *stackClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for existing bound class type.
	stackMutex.Lock()
	var value = stackClass[name]
	switch actual := value.(type) {
	case *stackClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &stackClass_[V]{
			// Initialize class constants.
		}
		stackClass[name] = class
	}
	stackMutex.Unlock()

	// Return a reference to the bound class type.
	return class
}

// CLASS METHODS

// Target

type stackClass_[V any] struct {
	// Define class constants.
	notation_ NotationLike
	defaultCapacity_ uint
}

// Constructors

func (c *stackClass_[V]) Make() StackLike[V] {
	// Validate the arguments.
	switch {
	default:
		return &stack_[V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *stackClass_[V]) MakeWithCapacity(capacity uint) StackLike[V] {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(capacity):
		panic("The capacity attribute is required for each Stack.")
	default:
		return &stack_[V]{
			// Initialize instance attributes.
			class_: c,
			capacity_: capacity,
		}
	}
}

func (c *stackClass_[V]) MakeFromArray(values []V) StackLike[V] {
	// Validate the arguments.
	switch {
	default:
		return &stack_[V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

func (c *stackClass_[V]) MakeFromSequence(values Sequential[V]) StackLike[V] {
	// Validate the arguments.
	switch {
	default:
		return &stack_[V]{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// Constants

func (c *stackClass_[V]) Notation() NotationLike {
	return c.notation_
}

func (c *stackClass_[V]) DefaultCapacity() uint {
	return c.defaultCapacity_
}

// INSTANCE METHODS

// Target

type stack_[V any] struct {
	// Define instance attributes.
	class_ StackClassLike[V]
	capacity_ uint
}

// Attributes

func (v *stack_[V]) GetClass() StackClassLike[V] {
	return v.class_
}

func (v *stack_[V]) GetCapacity() uint {
	return v.capacity_
}

// Limited[V]

func (v *stack_[V]) AddValue(value V) {
	// TBA - Implement the method.
}

func (v *stack_[V]) RemoveAll() {
	// TBA - Implement the method.
}

// Sequential[V]

func (v *stack_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *stack_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *stack_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *stack_[V]) GetIterator() age.IteratorLike[V] {
	var result_ age.IteratorLike[V]
	// TBA - Implement the method.
	return result_
}

// Public

func (v *stack_[V]) RemoveTop() V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

// Private

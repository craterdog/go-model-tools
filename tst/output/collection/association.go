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
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var associationClass = map[string]any{}
var associationMutex syn.Mutex

// Function

func Association[K comparable, V any]() AssociationClassLike[K, V] {
	// Generate the name of the bound class type.
	var result_ AssociationClassLike[K, V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	associationMutex.Lock()
	var value = associationClass[name]
	switch actual := value.(type) {
	case *associationClass_[K, V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &associationClass_[K, V]{
			// Initialize class constants.
		}
		associationClass[name] = result_
	}
	associationMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type associationClass_[K comparable, V any] struct {
	// Define class constants.
	// This class has no private constants.

}

// Constructors

func (c *associationClass_[K, V]) MakeWithAttributes(
	key K,
	value V,
) AssociationLike[K, V] {
	return &association_[K, V]{
		// Initialize instance attributes.
		class_: c,
		key_: key,
		value_: value,
	}
}

// INSTANCE METHODS

// Target

type association_[K comparable, V any] struct {
	// Define instance attributes.
	class_ AssociationClassLike[K, V]
	key_ K
	value_ V
}

// Attributes

func (v *association_[K, V]) GetClass() AssociationClassLike[K, V] {
	return v.class_
}

func (v *association_[K, V]) GetKey() K {
	return v.key_
}

func (v *association_[K, V]) GetValue() V {
	return v.value_
}

func (v *association_[K, V]) SetValue(value V) {
	v.value_ = value
}

// Private

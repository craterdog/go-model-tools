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

package ast

import ()

// CLASS ACCESS

// Reference

var predicateClass = &predicateClass_{
	// Initialize class constants.
}

// Function

func Predicate() PredicateClassLike {
	return predicateClass
}

// CLASS METHODS

// Target

type predicateClass_ struct {
	// Define class constants.
}

// Constructors

func (c *predicateClass_) MakeWithAtom(atom AtomLike) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		atom_: atom,
	}
}

func (c *predicateClass_) MakeWithElement(element ElementLike) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		element_: element,
	}
}

func (c *predicateClass_) MakeWithFilter(filter FilterLike) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		filter_: filter,
	}
}

func (c *predicateClass_) MakeWithPrecedence(precedence PrecedenceLike) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		precedence_: precedence,
	}
}

// INSTANCE METHODS

// Target

type predicate_ struct {
	// Define instance attributes.
	class_ PredicateClassLike
	atom_ AtomLike
	element_ ElementLike
	filter_ FilterLike
	precedence_ PrecedenceLike
}

// Attributes

func (v *predicate_) GetClass() PredicateClassLike {
	return v.class_
}

func (v *predicate_) GetAtom() AtomLike {
	return v.atom_
}

func (v *predicate_) GetElement() ElementLike {
	return v.element_
}

func (v *predicate_) GetFilter() FilterLike {
	return v.filter_
}

func (v *predicate_) GetPrecedence() PrecedenceLike {
	return v.precedence_
}

// Private

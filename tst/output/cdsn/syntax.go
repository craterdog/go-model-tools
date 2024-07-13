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

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var syntaxClass = &syntaxClass_{
	// Initialize class constants.
}

// Function

func Syntax() SyntaxClassLike {
	return syntaxClass
}

// CLASS METHODS

// Target

type syntaxClass_ struct {
	// Define class constants.
}

// Constructors

func (c *syntaxClass_) Make(
	headers abs.Sequential[HeaderLike],
	rules abs.Sequential[RuleLike],
	lexigrams abs.Sequential[LexigramLike],
) SyntaxLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(headers):
		panic("The headers attribute is required for each Syntax.")
	case col.IsUndefined(rules):
		panic("The rules attribute is required for each Syntax.")
	case col.IsUndefined(lexigrams):
		panic("The lexigrams attribute is required for each Syntax.")
	default:
		return &syntax_{
			// Initialize instance attributes.
			class_: c,
			headers_: headers,
			rules_: rules,
			lexigrams_: lexigrams,
		}
	}
}

// INSTANCE METHODS

// Target

type syntax_ struct {
	// Define instance attributes.
	class_ SyntaxClassLike
	headers_ abs.Sequential[HeaderLike]
	rules_ abs.Sequential[RuleLike]
	lexigrams_ abs.Sequential[LexigramLike]
}

// Attributes

func (v *syntax_) GetClass() SyntaxClassLike {
	return v.class_
}

func (v *syntax_) GetHeaders() abs.Sequential[HeaderLike] {
	return v.headers_
}

func (v *syntax_) GetRules() abs.Sequential[RuleLike] {
	return v.rules_
}

func (v *syntax_) GetLexigrams() abs.Sequential[LexigramLike] {
	return v.lexigrams_
}

// Private

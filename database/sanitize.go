// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"html"
	"net/url"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// sanitize is a helper function to verify the provided input
// field does not contain HTML content. If the input field
// does contain HTML, then the function will sanitize and
// potentially remove the HTML if deemed malicious.
func sanitize(field string) string {
	// create new HTML input microcosm-cc/bluemonday policy
	//
	// NOTE: consider using bluemonday.UGCPolicy() if the
	// strict policy has too much user impact
	p := bluemonday.StrictPolicy()

	// create a URL query unescaped string from the field
	queryUnescaped, err := url.QueryUnescape(field)
	if err != nil {
		// overwrite URL query unescaped string with field
		queryUnescaped = field
	}

	// create an HTML escaped string from the field
	htmlEscaped := html.EscapeString(queryUnescaped)

	// create a microcosm-cc/bluemonday escaped string from the field
	bluemondayEscaped := p.Sanitize(queryUnescaped)

	// check if the field contains html
	if !strings.EqualFold(htmlEscaped, bluemondayEscaped) {
		// create new HTML input microcosm-cc/bluemonday policy
		return bluemondayEscaped
	}

	// return the unmodified field
	return field
}

// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_Service_Nullify(t *testing.T) {
	// setup types
	var s *Service

	want := &Service{
		ID:           sql.NullInt64{Int64: 0, Valid: false},
		BuildID:      sql.NullInt64{Int64: 0, Valid: false},
		RepoID:       sql.NullInt64{Int64: 0, Valid: false},
		Number:       sql.NullInt32{Int32: 0, Valid: false},
		Name:         sql.NullString{String: "", Valid: false},
		Image:        sql.NullString{String: "", Valid: false},
		Status:       sql.NullString{String: "", Valid: false},
		Error:        sql.NullString{String: "", Valid: false},
		ExitCode:     sql.NullInt32{Int32: 0, Valid: false},
		Created:      sql.NullInt64{Int64: 0, Valid: false},
		Started:      sql.NullInt64{Int64: 0, Valid: false},
		Finished:     sql.NullInt64{Int64: 0, Valid: false},
		Host:         sql.NullString{String: "", Valid: false},
		Runtime:      sql.NullString{String: "", Valid: false},
		Distribution: sql.NullString{String: "", Valid: false},
	}

	// setup tests
	tests := []struct {
		service *Service
		want    *Service
	}{
		{
			service: testService(),
			want:    testService(),
		},
		{
			service: s,
			want:    nil,
		},
		{
			service: new(Service),
			want:    want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.service.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Service_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Service)

	want.SetID(1)
	want.SetBuildID(1)
	want.SetRepoID(1)
	want.SetNumber(1)
	want.SetName("postgres")
	want.SetImage("postgres:12-alpine")
	want.SetStatus("running")
	want.SetError("")
	want.SetExitCode(0)
	want.SetCreated(1563474076)
	want.SetStarted(1563474078)
	want.SetFinished(1563474079)
	want.SetHost("example.company.com")
	want.SetRuntime("docker")
	want.SetDistribution("linux")

	// run test
	got := testService().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Service_Validate(t *testing.T) {
	tests := []struct {
		failure bool
		service *Service
	}{
		{
			failure: false,
			service: testService(),
		},
		{ // no build_id set for service
			failure: true,
			service: &Service{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
				Number: sql.NullInt32{Int32: 1, Valid: true},
				Name:   sql.NullString{String: "postgres", Valid: true},
				Image:  sql.NullString{String: "postgres:12-alpine", Valid: true},
			},
		},
		{ // no repo_id set for service
			failure: true,
			service: &Service{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Name:    sql.NullString{String: "postgres", Valid: true},
				Image:   sql.NullString{String: "postgres:12-alpine", Valid: true},
			},
		},
		{ // no number set for service
			failure: true,
			service: &Service{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Name:    sql.NullString{String: "postgres", Valid: true},
				Image:   sql.NullString{String: "postgres:12-alpine", Valid: true},
			},
		},
		{ // no name set for service
			failure: true,
			service: &Service{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Image:   sql.NullString{String: "postgres:12-alpine", Valid: true},
			},
		},
		{ // no image set for service
			failure: true,
			service: &Service{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Name:    sql.NullString{String: "postgres", Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.service.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestDatabase_ServiceFromLibrary(t *testing.T) {
	// setup types
	s := new(library.Service)

	s.SetID(1)
	s.SetBuildID(1)
	s.SetRepoID(1)
	s.SetNumber(1)
	s.SetName("postgres")
	s.SetImage("postgres:12-alpine")
	s.SetStatus("running")
	s.SetError("")
	s.SetExitCode(0)
	s.SetCreated(1563474076)
	s.SetStarted(1563474078)
	s.SetFinished(1563474079)
	s.SetHost("example.company.com")
	s.SetRuntime("docker")
	s.SetDistribution("linux")

	want := testService()

	// run test
	got := ServiceFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ServiceFromLibrary is %v, want %v", got, want)
	}
}

// testService is a test helper function to create a Service
// type with all fields set to a fake value.
func testService() *Service {
	return &Service{
		ID:           sql.NullInt64{Int64: 1, Valid: true},
		BuildID:      sql.NullInt64{Int64: 1, Valid: true},
		RepoID:       sql.NullInt64{Int64: 1, Valid: true},
		Number:       sql.NullInt32{Int32: 1, Valid: true},
		Name:         sql.NullString{String: "postgres", Valid: true},
		Image:        sql.NullString{String: "postgres:12-alpine", Valid: true},
		Status:       sql.NullString{String: "running", Valid: true},
		Error:        sql.NullString{String: "", Valid: false},
		ExitCode:     sql.NullInt32{Int32: 0, Valid: false},
		Created:      sql.NullInt64{Int64: 1563474076, Valid: true},
		Started:      sql.NullInt64{Int64: 1563474078, Valid: true},
		Finished:     sql.NullInt64{Int64: 1563474079, Valid: true},
		Host:         sql.NullString{String: "example.company.com", Valid: true},
		Runtime:      sql.NullString{String: "docker", Valid: true},
		Distribution: sql.NullString{String: "linux", Valid: true},
	}
}

package plan

import "testing"

import . "fig/model"
import . "fig/repos"

func TestNoDependencies(t *testing.T) {
	repo := NewMemoryRepository()
	WritePackage(repo, NewPackageBuilder("foo", "1.2.3").Config("default").End().Build())
	planner := NewPlanner(repo)
	configs, err := planner.Plan(NewDescriptor("foo","1.2.3","default"))
	if err != nil {
		t.Fatal(err)
	}
	if len(configs) != 1 {
		t.Errorf("Expected 1 config, got: %d", len(configs))
	}
}

func TestSimpleDependency(t *testing.T) {
	repo := NewMemoryRepository()
	foo := NewPackageBuilder("foo", "1.2.3").
		Config("default").Include("bar", "4.5.6", "default").End().
		Build()
	bar := NewPackageBuilder("bar", "4.5.6").
		Config("default").End().
		Build()
	WritePackage(repo, foo)
	WritePackage(repo, bar)
	planner := NewPlanner(repo)
	configs, err := planner.Plan(NewDescriptor("foo","1.2.3","default"))
	if err != nil {
		t.Fatal(err)
	}
	if len(configs) != 2 {
		t.Errorf("Expected 2 configs, got: %d", len(configs))
	}
}

func TestTransitiveDependency(t *testing.T) {
	repo := NewMemoryRepository()
	foo := NewPackageBuilder("foo", "1.2.3").
		Config("default").Include("bar", "4.5.6", "default").End().
		Build()
	bar := NewPackageBuilder("bar", "4.5.6").
		Config("default").Include("baz", "7.8.9", "default").End().
		Build()
	baz := NewPackageBuilder("baz", "7.8.9").
		Config("default").End().
		Build()
	WritePackage(repo, foo)
	WritePackage(repo, bar)
	WritePackage(repo, baz)
	planner := NewPlanner(repo)
	configs, err := planner.Plan(NewDescriptor("foo","1.2.3","default"))
	if err != nil {
		t.Fatal(err)
	}
	if len(configs) != 3 {
		t.Errorf("Expected 3 configs, got: %d", len(configs))
	}
}

package app_env_test

import (
	. "github.com/marugoshi/gobm/shared/app_env"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	ResetEnv()
	os.Setenv(EnvKey, "")
	if GetName() != "development" {
		t.Fatalf("Invalid Value: `loadEnv()` failed.")
	}

	if IsDevelopment() != true {
		t.Fatalf("Invalid Value `IsDevelopment()`.")
	}

	if IsStaging() == true {
		t.Fatalf("Invalid Value `IsStaging()`.")
	}

	if IsProduction() == true {
		t.Fatalf("Invalid Value `IsProduction()`.")
	}
}

func TestGetName(t *testing.T) {
	ResetEnv()
	os.Setenv(EnvKey, "development")
	if GetName() != "development" {
		t.Fatalf("GetName() development.")
	}

	ResetEnv()
	os.Setenv(EnvKey, "devElopmeNt")
	if GetName() != "development" {
		t.Fatalf("GetName() development.")
	}

	ResetEnv()
	os.Setenv(EnvKey, "staging")
	if GetName() != "staging" {
		t.Fatalf("GetName() staging.")
	}

	ResetEnv()
	os.Setenv(EnvKey, "StaginG")
	if GetName() != "staging" {
		t.Fatalf("GetName() staging.")
	}

	ResetEnv()
	os.Setenv(EnvKey, "production")
	if GetName() != "production" {
		t.Fatalf("GetName() production.")
	}

	ResetEnv()
	os.Setenv(EnvKey, "proDuctIoN")
	if GetName() != "production" {
		t.Fatalf("GetName() production.")
	}

	ResetEnv()
	os.Setenv(EnvKey, "")
	if GetName() != "development" {
		t.Fatalf("GetName() default development.")
	}
}

func TestIsDevelopment(t *testing.T) {
	ResetEnv()
	os.Setenv(EnvKey, "development")

	if GetName() != "development" {
		t.Fatalf("Invalid Value `Getname()`.")
	}

	if IsDevelopment() != true {
		t.Fatalf("Invalid Value `IsDevelopment()`.")
	}

	if IsStaging() == true {
		t.Fatalf("Invalid Value `IsStaging()`.")
	}

	if IsProduction() == true {
		t.Fatalf("Invalid Value `IsProduction()`.")
	}
}

func TestIsStaging(t *testing.T) {
	ResetEnv()
	os.Setenv(EnvKey, "staging")

	if GetName() != "staging" {
		t.Fatalf("Invalid Value `Getname()`.")
	}

	if IsDevelopment() == true {
		t.Fatalf("Invalid Value `IsDevelopment()`.")
	}

	if IsStaging() != true {
		t.Fatalf("Invalid Value `IsStaging()`.")
	}

	if IsProduction() == true {
		t.Fatalf("Invalid Value `IsProduction()`.")
	}
}

func TestIsProduction(t *testing.T) {
	ResetEnv()
	os.Setenv(EnvKey, "production")

	if GetName() != "production" {
		t.Fatalf("Invalid Value `Getname()`.")
	}

	if IsDevelopment() == true {
		t.Fatalf("Invalid Value `IsDevelopment()`.")
	}

	if IsStaging() == true {
		t.Fatalf("Invalid Value `IsStaging()`.")
	}

	if IsProduction() != true {
		t.Fatalf("Invalid Value `IsProduction()`.")
	}
}

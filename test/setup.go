package test

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

// TeardownFunc represents a function that will be used to destroy
// the test environment after tests are done
type TeardownFunc func() error

// ConfigurationFunc represents a function that will run any kind of
// setup to prepare the test environment. Also, it returns a teardown
// function to be called to destroy changes made after tests are done
type ConfigurationFunc func() (TeardownFunc, error)

// SkipSetup bool
var SkipSetup bool

// GetSetup receives a list of functions that are used to configure and teardown
// the test environment.
func GetSetup(configurations []ConfigurationFunc) func(*testing.M) {
	return func(m *testing.M) {
		var err error = nil
		var exitCode int

		teardownFuncs := []func() error{}

		// configure test environment
		for _, configure := range configurations {
			var teardown TeardownFunc

			if !SkipSetup {
				teardown, err = configure()
				if err != nil {
					fmt.Println(err)
					exitCode = 1
					teardownFuncs = append(teardownFuncs, teardown)
					break
				}
				teardownFuncs = append(teardownFuncs, teardown)
			}
		}

		// run tests
		if err == nil {
			exitCode = m.Run()
		}

		// teardown test environment
		for i := len(teardownFuncs) - 1; i >= 0; i-- {
			teardown := teardownFuncs[i]

			err = teardown()
			if err != nil {
				fmt.Println(err)
				exitCode = 1
			}
		}

		os.Exit(exitCode)
	}
}

// ConfigureTempDir creates a temporary directory within the module path
// to run tests, fork fixtures, etc
func ConfigureTempDir() (TeardownFunc, error) {
	testRelativePath := "./"
	testDir := "tmp"
	testPath := testRelativePath + testDir

	// start test isolated environment
	fmt.Println("creating temporary test environment")

	wdir, err := os.Getwd()
	if err != nil {
		return func() error {
			return nil
		}, fmt.Errorf("error getting current directory: %v", err)
	}

	if err := os.MkdirAll(testDir, os.ModePerm); err != nil {
		return func() error {
			return nil
		}, fmt.Errorf("error creating tmp dir")
	}

	err = os.Chdir(testPath)
	if err != nil {
		return func() error {
			return nil
		}, fmt.Errorf("error creating temporary directory for test files: %v", err)
	}

	teardown := func() error {
		fmt.Println("cleaning test environment")

		if err := os.Chdir(wdir); err != nil {
			return fmt.Errorf("failed to clean test environment")
		}

		err = os.RemoveAll(testPath)
		if err != nil {
			return fmt.Errorf("failed to remove test directory 'tmp': %v", err)
		}

		return nil
	}

	return teardown, nil
}

// ConfigureFlags parses command flags
func ConfigureFlags() (TeardownFunc, error) {
	flag.BoolVar(&SkipSetup, "skip-setup", false, "skip environment setup")
	flag.Parse()

	teardown := func() error { return nil }

	return teardown, nil
}

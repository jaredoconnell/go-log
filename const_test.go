package log_test

import (
	"testing"

	"go.arcalot.io/log/v2"
)

func TestLevelShouldPrint(t *testing.T) {
	testCases := map[string]struct {
		minimumLevel log.Level
		messageLevel log.Level
		shouldPrint  bool
	}{
		"debug-debug":     {minimumLevel: log.LevelDebug, messageLevel: log.LevelDebug, shouldPrint: true},
		"debug-info":      {minimumLevel: log.LevelDebug, messageLevel: log.LevelInfo, shouldPrint: true},
		"debug-warning":   {minimumLevel: log.LevelDebug, messageLevel: log.LevelWarning, shouldPrint: true},
		"debug-error":     {minimumLevel: log.LevelDebug, messageLevel: log.LevelError, shouldPrint: true},
		"info-debug":      {minimumLevel: log.LevelInfo, messageLevel: log.LevelDebug, shouldPrint: false},
		"info-info":       {minimumLevel: log.LevelInfo, messageLevel: log.LevelInfo, shouldPrint: true},
		"info-warning":    {minimumLevel: log.LevelInfo, messageLevel: log.LevelWarning, shouldPrint: true},
		"info-error":      {minimumLevel: log.LevelInfo, messageLevel: log.LevelError, shouldPrint: true},
		"warning-debug":   {minimumLevel: log.LevelWarning, messageLevel: log.LevelDebug, shouldPrint: false},
		"warning-info":    {minimumLevel: log.LevelWarning, messageLevel: log.LevelInfo, shouldPrint: false},
		"warning-warning": {minimumLevel: log.LevelWarning, messageLevel: log.LevelWarning, shouldPrint: true},
		"warning-error":   {minimumLevel: log.LevelWarning, messageLevel: log.LevelError, shouldPrint: true},
		"error-debug":     {minimumLevel: log.LevelError, messageLevel: log.LevelDebug, shouldPrint: false},
		"error-info":      {minimumLevel: log.LevelError, messageLevel: log.LevelInfo, shouldPrint: false},
		"error-warning":   {minimumLevel: log.LevelError, messageLevel: log.LevelWarning, shouldPrint: false},
		"error-error":     {minimumLevel: log.LevelError, messageLevel: log.LevelError, shouldPrint: true},
	}

	for name, tc := range testCases {
		testCase := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if shouldPrint := testCase.minimumLevel.ShouldPrint(testCase.messageLevel); shouldPrint != testCase.shouldPrint {
				t.Fatalf("Incorrect ShouldPrint value: %v", shouldPrint)
			}
		})
	}
}

func TestLevelValidate(t *testing.T) {
	testCases := map[string]struct {
		minimumLevel log.Level
		valid        bool
	}{
		"debug": {
			minimumLevel: log.LevelDebug,
			valid:        true,
		},
		"info": {
			minimumLevel: log.LevelInfo,
			valid:        true,
		},
		"warning": {
			minimumLevel: log.LevelWarning,
			valid:        true,
		},
		"error": {
			minimumLevel: log.LevelError,
			valid:        true,
		},
		"test": {
			minimumLevel: log.Level("test"),
			valid:        false,
		},
		"empty": {
			minimumLevel: log.Level(""),
			valid:        false,
		},
	}
	for name, tc := range testCases {
		testCase := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if err := testCase.minimumLevel.Validate(); (err == nil) != testCase.valid {
				t.Fatalf("Incorrect valid value: %v", err)
			}
		})
	}
}

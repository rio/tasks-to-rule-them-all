package server

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCharacterCounter(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "alphabet",
			input:    "abcdefghijklmnopqrstuvwxyz",
			expected: 26,
		},
		{
			name:     "see no evil, hear no evil, speak no evil, do no evil",
			input:    "🙈🙉🙊🐵",
			expected: 16,
		},
	}

	for _, tt := range testCases {
		result := countCharacters(tt.input)

		if result != tt.expected {
			t.Errorf("character length didn't match for case %q: %d != %d", tt.name, result, tt.expected)
		}
	}
}

func TestEchoHandler(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelError)

	request := httptest.NewRequest(
		http.MethodPost,
		"/echo",
		strings.NewReader(`{"message": test}`),
	)

	recorder := httptest.NewRecorder()

	Echo(recorder, request)

	result := recorder.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %d, expected %d", result.StatusCode, http.StatusOK)
	}

	contentType := result.Header.Get("Content-Type")
	expectedContentType := "application/json"
	if contentType != expectedContentType {
		t.Errorf("Unexpected Content-Type header: got %q, expected %q ", contentType, expectedContentType)
	}
}

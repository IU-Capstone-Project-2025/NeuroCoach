package services

import (
	"testing"
)

func TestOpenRouterClient_GetCurrentModel(t *testing.T) {
	client := NewOpenRouterClient("test-key")

	model := client.getCurrentModel()
	if model != availableModels[0] {
		t.Errorf("Expected first model %s, got %s", availableModels[0], model)
	}
}

func TestOpenRouterClient_SwitchToNextModel(t *testing.T) {
	client := NewOpenRouterClient("test-key")

	initialModel := client.getCurrentModel()
	client.switchToNextModel()
	nextModel := client.getCurrentModel()

	if initialModel == nextModel {
		t.Error("Model should have changed after switching")
	}

	if nextModel != availableModels[1] {
		t.Errorf("Expected second model %s, got %s", availableModels[1], nextModel)
	}
}

func TestOpenRouterClient_IsModelError(t *testing.T) {
	client := NewOpenRouterClient("test-key")

	testCases := []struct {
		error    string
		expected bool
	}{
		{"model unavailable", true},
		{"provider error", true},
		{"503 service unavailable", true},
		{"overloaded", true},
		{"network timeout", false},
		{"invalid request", false},
	}

	for _, tc := range testCases {
		result := client.isModelError(&mockError{tc.error})
		if result != tc.expected {
			t.Errorf("For error '%s', expected %v, got %v", tc.error, tc.expected, result)
		}
	}
}

func TestAvailableModels_Count(t *testing.T) {
	if len(availableModels) < 5 {
		t.Errorf("Expected at least 5 models, got %d", len(availableModels))
	}

	// Check that DeepSeek is first
	if availableModels[0] != "deepseek/deepseek-chat-v3-0324:free" {
		t.Errorf("Expected DeepSeek as first model, got %s", availableModels[0])
	}

	// Check that Cypher Alpha is second
	if availableModels[1] != "openrouter/cypher-alpha:free" {
		t.Errorf("Expected Cypher Alpha as second model, got %s", availableModels[1])
	}
}

type mockError struct {
	message string
}

func (e *mockError) Error() string {
	return e.message
}

package testcontainers_wiremock

import (
	"bytes"
	"context"
	"path/filepath"
	"strings"
	"testing"
)

func TestWireMock(t *testing.T) {
	// Create Container
	ctx := context.Background()
	container, err := RunContainer(ctx,
		WithMappingFile("hello", filepath.Join("testdata", "hello-world.json")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	statusCode, out, err := SendHttpGet(container, "/hello", nil)
	if err != nil {
		t.Fatal(err, "Failed to get a response")
	}
	if statusCode != 200 {
		t.Fatalf("expected HTTP-200 but got %d", statusCode)
	}
	if string(out) != "Hello, world!" {
		t.Fatalf("expected 'Hello, world!' but got %v", string(out))
	}
}

func TestWireMockWithResource(t *testing.T) {
	// Create Container
	ctx := context.Background()
	container, err := RunContainer(ctx,
		WithMappingFile("hello", filepath.Join("testdata", "hello-world-resource.json")),
		WithFile("hello-world-resource-response.xml", filepath.Join("testdata", "hello-world-resource-response.xml")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	statusCode, out, err := SendHttpGet(container, "/hello-from-file", nil)
	if err != nil {
		t.Fatal(err, "Failed to get a response")
	}
	if statusCode != 200 {
		t.Fatalf("expected HTTP-200 but got %d", statusCode)
	}
	if !strings.Contains(out, "Hello, world!") {
		t.Fatalf("expected 'Hello, world!' but got %v", string(out))
	}
}

func TestSendHttpGetWorksWithQueryParamsPassedInArgument(t *testing.T) {
	// Create Container
	ctx := context.Background()
	container, err := RunContainer(ctx,
		WithMappingFile("get", filepath.Join("testdata", "url-with-query-params.json")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	statusCode, out, err := SendHttpGet(container, "/get", map[string]string{"query": "test"})
	if err != nil {
		t.Fatal(err, "Failed to get a response")
	}
	if statusCode != 200 {
		t.Fatalf("expected HTTP-200 but got %d", statusCode)
	}
	if out != "" {
		t.Fatalf("expected 'Hello, world!' but got %v", string(out))
	}
}

func TestSendHttpGetWorksWithQueryParamsProvidedInURL(t *testing.T) {
	// Create Container
	ctx := context.Background()
	container, err := RunContainer(ctx,
		WithMappingFile("get", filepath.Join("testdata", "url-with-query-params.json")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	statusCode, out, err := SendHttpGet(container, "/get?query=test", nil)
	if err != nil {
		t.Fatal(err, "Failed to get a response")
	}
	if statusCode != 200 {
		t.Fatalf("expected HTTP-200 but got %d", statusCode)
	}
	if out != "" {
		t.Fatalf("expected 'Hello, world!' but got %v", string(out))
	}
}

func TestSendHttpDelete(t *testing.T) {
	// Create Container
	ctx := context.Background()
	container, err := RunContainer(ctx,
		WithMappingFile("delete", filepath.Join("testdata", "204-no-content.json")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	statusCode, out, err := SendHttpDelete(container, "/delete")
	if err != nil {
		t.Fatal(err, "Failed to get a response")
	}
	if statusCode != 204 {
		t.Fatalf("expected HTTP-200 but got %d", statusCode)
	}
	if out != "" {
		t.Fatalf("expected 'Hello, world!' but got %v", string(out))
	}
}

func TestSendHttpPatch(t *testing.T) {
	// Create Container
	ctx := context.Background()
	container, err := RunContainer(ctx,
		WithMappingFile("patch", filepath.Join("testdata", "200-patch.json")),
		WithFile("sample-model.json", filepath.Join("testdata", "sample-model.json")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	var jsonBody = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	statusCode, out, err := SendHttpPatch(container, "/patch", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err, "Failed to get a response")
	}
	if statusCode != 200 {
		t.Fatalf("expected HTTP-200 but got %d", statusCode)
	}
	if !strings.Contains(out, "sampleField1") {
		t.Fatalf("expected 'Hello, world!' but got %v", string(out))
	}
}

func TestSendHttpPut(t *testing.T) {
	// Create Container
	ctx := context.Background()
	container, err := RunContainer(ctx,
		WithMappingFile("put", filepath.Join("testdata", "200-put.json")),
		WithFile("sample-model.json", filepath.Join("testdata", "sample-model.json")),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	var jsonBody = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	statusCode, out, err := SendHttpPut(container, "/put", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err, "Failed to get a response")
	}
	if statusCode != 200 {
		t.Fatalf("expected HTTP-200 but got %d", statusCode)
	}
	if !strings.Contains(out, "sampleField1") {
		t.Fatalf("expected 'Hello, world!' but got %v", string(out))
	}
}

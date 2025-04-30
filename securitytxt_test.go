package securitytxt_test

import (
	"github.com/dreadl0ck/securitytxt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {

	// configure via env
	assert.NoError(t, os.Setenv("EXPIRES", "2025-03-26T11:00:00.000Z"))
	assert.NoError(t, os.Setenv("COMMENT", "this is a comment"))
	assert.NoError(t, os.Setenv("CONTACT", "mailto:security@org.com"))
	assert.NoError(t, os.Setenv("ACKNOWLEDGMENT", "https://example.com/halloffame"))
	assert.NoError(t, os.Setenv("CANONICAL", "https://example.com/canonical"))
	assert.NoError(t, os.Setenv("ENCRYPTION", "https://example.com/pgpkey.txt"))
	assert.NoError(t, os.Setenv("HIRING", "https://example.com/hiring"))
	assert.NoError(t, os.Setenv("PREFERRED_LANGUAGES", "en, de"))
	assert.NoError(t, os.Setenv("POLICY", "https://example.com/policy"))
	assert.NoError(t, os.Setenv("CSAF", "https://example.com/csaf"))

	// Create a request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request
	securitytxt.Handler().ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expected := `# this is a comment
Contact: mailto:security@org.com
Expires: 2025-03-26T11:00:00.000Z
Encryption: https://example.com/pgpkey.txt
Acknowledgments: https://example.com/halloffame
Preferred-Languages: en, de
Canonical: https://example.com/canonical
Policy: https://example.com/policy
Hiring: https://example.com/hiring
CSAF: https://example.com/csaf`

	body, _ := io.ReadAll(rr.Body)
	//if strings.TrimSpace(string(body)) != expected {
	//	t.Errorf("Expected body %s but got %s", expected, body)
	//}
	assert.Equal(t, expected, string(body))

	// Check content type
	if contentType := rr.Header().Get("Content-Type"); contentType != "text/plain" {
		t.Errorf("Expected content type application/json but got %s", contentType)
	}
}

package web

import (
	"testing"
)

func TestSanitizeHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "allows br tag",
			input:    "Hello<br>World",
			expected: "Hello<br>World",
		},
		{
			name:     "allows br self-closing tag",
			input:    "Hello<br/>World",
			expected: "Hello<br/>World",
		},
		{
			name:     "allows br tag with space",
			input:    "Hello<br />World",
			expected: "Hello<br />World",
		},
		{
			name:     "allows strong tag",
			input:    "This is <strong>important</strong>",
			expected: "This is <strong>important</strong>",
		},
		{
			name:     "allows em tag",
			input:    "This is <em>emphasized</em>",
			expected: "This is <em>emphasized</em>",
		},
		{
			name:     "blocks script tag",
			input:    "<script>alert('XSS')</script>",
			expected: "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;",
		},
		{
			name:     "blocks onclick attribute",
			input:    "<div onclick='alert(1)'>Click me</div>",
			expected: "&lt;div onclick=&#39;alert(1)&#39;&gt;Click me&lt;/div&gt;",
		},
		{
			name:     "blocks img tag",
			input:    "<img src=x onerror='alert(1)'>",
			expected: "&lt;img src=x onerror=&#39;alert(1)&#39;&gt;",
		},
		{
			name:     "allows multiple br tags",
			input:    "Line 1<br>Line 2<br>Line 3",
			expected: "Line 1<br>Line 2<br>Line 3",
		},
		{
			name:     "mixed allowed and blocked tags",
			input:    "Hello<br><script>evil</script><strong>World</strong>",
			expected: "Hello<br>&lt;script&gt;evil&lt;/script&gt;<strong>World</strong>",
		},
		{
			name:     "preserves special characters",
			input:    "5 < 10 & 10 > 5",
			expected: "5 &lt; 10 &amp; 10 &gt; 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizeHTML(tt.input)
			if result != tt.expected {
				t.Errorf("sanitizeHTML(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
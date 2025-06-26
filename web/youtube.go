package web

import (
	"fmt"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// YouTubeEmbed represents a YouTube embed node in the AST
type YouTubeEmbed struct {
	ast.BaseBlock
	URL string
}

// Dump implements ast.Node.Dump
func (n *YouTubeEmbed) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// Kind implements ast.Node.Kind
func (n *YouTubeEmbed) Kind() ast.NodeKind {
	return YouTubeEmbedKind
}

// YouTubeEmbedKind is the node kind for YouTube embeds
var YouTubeEmbedKind = ast.NewNodeKind("YouTubeEmbed")

// YouTubeParser parses YouTube embed tags as block elements
type YouTubeParser struct{}

// Trigger returns trigger characters for the parser
func (s *YouTubeParser) Trigger() []byte {
	return []byte{'<'}
}

// Open checks if the current line contains a YouTube embed
func (s *YouTubeParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, segment := reader.PeekLine()
	
	// Look for <youtube url="..." /> pattern on its own line (with optional whitespace)
	re := regexp.MustCompile(`^\s*<youtube\s+url="([^"]+)"\s*/>\s*$`)
	match := re.FindSubmatch(line)
	
	if match == nil {
		return nil, parser.NoChildren
	}
	
	// Extract URL
	url := string(match[1])
	
	// Advance the reader past this line
	reader.Advance(segment.Len())
	
	// Create YouTube embed node
	node := &YouTubeEmbed{
		URL: url,
	}
	
	return node, parser.NoChildren
}

// Continue is not used for this parser
func (s *YouTubeParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

// Close finalizes the node
func (s *YouTubeParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	// Nothing to do
}

// CanInterruptParagraph returns true if this parser can interrupt paragraphs
func (s *YouTubeParser) CanInterruptParagraph() bool {
	return true
}

// CanAcceptIndentedLine returns false
func (s *YouTubeParser) CanAcceptIndentedLine() bool {
	return false
}

// YouTubeRenderer renders YouTube embed nodes to HTML
type YouTubeRenderer struct{}

// RegisterFuncs registers rendering functions
func (r *YouTubeRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(YouTubeEmbedKind, r.renderYouTubeEmbed)
}

// renderYouTubeEmbed renders a YouTube embed to HTML
func (r *YouTubeRenderer) renderYouTubeEmbed(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	
	embed, ok := node.(*YouTubeEmbed)
	if !ok {
		return ast.WalkContinue, nil
	}
	
	// Extract video ID from YouTube URL
	videoID := extractYouTubeID(embed.URL)
	if videoID == "" {
		// If we can't extract ID, just render the original URL as a link
		w.WriteString(fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, embed.URL, embed.URL))
		return ast.WalkContinue, nil
	}
	
	// Render responsive YouTube embed
	embedHTML := fmt.Sprintf(`<div class="youtube-embed">
  <iframe 
    width="560" 
    height="315" 
    src="https://www.youtube.com/embed/%s" 
    title="YouTube video player" 
    frameborder="0" 
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
    allowfullscreen>
  </iframe>
</div>`, videoID)
	
	w.WriteString(embedHTML)
	return ast.WalkContinue, nil
}

// extractYouTubeID extracts the video ID from various YouTube URL formats
func extractYouTubeID(url string) string {
	// Handle various YouTube URL formats
	patterns := []string{
		`(?:youtube\.com/watch\?v=|youtu\.be/|youtube\.com/embed/)([a-zA-Z0-9_-]{11})`,
		`youtube\.com/watch\?.*v=([a-zA-Z0-9_-]{11})`,
	}
	
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(url)
		if len(match) > 1 {
			return match[1]
		}
	}
	
	return ""
}

// YouTubeExtension is the Goldmark extension for YouTube embeds
type YouTubeExtension struct{}

// Extend extends the Goldmark parser with YouTube embed support
func (e *YouTubeExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(&YouTubeParser{}, 200),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&YouTubeRenderer{}, 200),
		),
	)
}

// NewYouTubeExtension creates a new YouTube extension
func NewYouTubeExtension() goldmark.Extender {
	return &YouTubeExtension{}
}
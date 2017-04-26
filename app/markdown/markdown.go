package markdown

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"

	"aahframework.org/aah.v0-unstable"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
)

var (
	mdCache = make(map[string][]byte)

	markdownHTMLFlags = 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	markdownExtensions = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_AUTO_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS

	markdownOptions = blackfriday.Options{Extensions: markdownExtensions}
)

// ContentBasePath method returns the Markdown files base path.
func ContentBasePath() string {
	return filepath.Join(aah.AppBaseDir(), "content")
}

// FilePath method returns markdown file path from given path.
// it bacially remove any extension and adds ".md"
func FilePath(reqPath, prefix string) string {
	reqPath = strings.ToLower(strings.TrimPrefix(reqPath, "/"))
	reqPath = ess.StripExt(reqPath) + ".md"
	return filepath.Clean(filepath.Join(prefix, reqPath))
}

// ReadAll method reads the markdown file and returns the bytes.
func ReadAll(reqPath string) []byte {
	bytes, err := ioutil.ReadFile(reqPath)
	if err != nil {
		log.Error(err)
		return []byte("")
	}
	return bytes
}

// Parse method parsed the markdown content into html using blackfriday library
// and returns the byte slice.
func Parse(input []byte) []byte {
	htmlRender := blackfriday.HtmlRenderer(markdownHTMLFlags, "", "")
	return blackfriday.MarkdownOptions(input, htmlRender, markdownOptions)
}

// Get method returns the parsed markdown content for given URL path.
func Get(mdPath string) []byte {
	cache := aah.AppConfig().BoolDefault("markdown.cache", false)
	if cache {
		if c, found := mdCache[mdPath]; found {
			return c
		}
	}

	mf := ReadAll(mdPath)
	content := Parse(mf)

	if cache {
		// put it in the cache
		mdCache[mdPath] = content
	}

	return content
}

// ClearCache method clears the Markdown cache.
func ClearCache() {
	mdCache = make(map[string][]byte)
}

// ClearCacheByFile method clears cache by file.
func ClearCacheByFile(name string) {
	key := ""
	for k := range mdCache {
		if strings.Contains(k, name) {
			key = k
			break
		}
	}

	if !ess.IsStrEmpty(key) {
		delete(mdCache, key)
	}
}

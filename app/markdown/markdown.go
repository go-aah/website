package markdown

import (
	"bufio"
	"os"
	"strings"

	"github.com/russross/blackfriday"

	"aahframework.org/aah.v0"
	"aahframework.org/log.v0"

	"github.com/go-aah/website/app/models"
)

var (
	articleCache = make(map[string]*models.Article)
	// mdCache      = make(map[string][]byte)

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

	isCacheEnabled bool
)

// Parse method parsed the markdown content into html using blackfriday library
// and create Article object.
func Parse(lines []string) *models.Article {
	pos := 0
	for idx, l := range lines {
		if strings.TrimSpace(l) == "---" {
			pos = idx + 1
			break
		}
	}

	article := &models.Article{}

	for _, v := range lines[:pos] {
		if v == "---" {
			break
		}
		idx := strings.IndexByte(v, ':')
		if idx == -1 {
			continue
		}
		switch v[:idx] {
		case "Title":
			article.Title = strings.TrimSpace(v[idx+1:])
		case "Desc":
			article.Desc = strings.TrimSpace(v[idx+1:])
		case "Keywords":
			article.Keywords = strings.TrimSpace(v[idx+1:])
		}
	}

	content := strings.Join(lines[pos:], "\n")
	htmlRender := blackfriday.HtmlRenderer(markdownHTMLFlags, "", "")
	article.Content = string(blackfriday.MarkdownOptions([]byte(content), htmlRender, markdownOptions))

	return article
}

// Get method returns the parsed markdown content for given URL path.
func Get(mdPath string) (*models.Article, bool) {
	if isCacheEnabled {
		if article, found := articleCache[mdPath]; found {
			return article, true
		}
	}

	f, err := os.Open(mdPath)
	if err != nil {
		return nil, false
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	article := Parse(lines)
	article.File = mdPath

	if article.IsContent() && isCacheEnabled {
		articleCache[mdPath] = article
	}

	return article, article.IsContent()
}

// ClearCache method clears the Markdown cache.
func ClearCache() {
	log.Info("Clearing cache")
	articleCache = make(map[string]*models.Article)
}

// ClearCacheByFile method clears cache by file.
// func ClearCacheByFile(name string) {
// 	key := ""
// 	for k := range mdCache {
// 		if strings.Contains(k, name) {
// 			key = k
// 			break
// 		}
// 	}
//
// 	if !ess.IsStrEmpty(key) {
// 		delete(mdCache, key)
// 	}
// }

func init() {
	aah.OnStart(func(e *aah.Event) {
		isCacheEnabled = aah.AppConfig().BoolDefault("markdown.cache", false)
	})
}

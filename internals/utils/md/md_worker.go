package md

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// MdWorker Интерфейс для работы с Markdown текстом
type MdWorker interface {
	GetMainPageDocApi()
}

var log = logrus.New()

// GetMainPageDocApi Возвращает отрендереную из md в html страницу файл
func (worker MdWorkerImpl) GetMainPageDocApi() string {
	return convertMarkdownToHtml("doc/main-page/doc.md", "Документация главной страницы")
}

// Чтение Markdown файла
func readMdFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Errorln("Error opening file ", path, ". Error: ", err)
	}

	defer file.Close()

	md, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorln("Error read file. Error: ", err)
	}

	return md
}

// Конвертирование Markdown файла в HTML
func convertMarkdownToHtml(path string, title string) string {
	md := readMdFile(path)
	css := getCssMarkdownFile()
	htmlTemplate := getHtmlTemplate()

	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{
		Flags: htmlFlags,
		Title: title,
	}
	renderer := html.NewRenderer(opts)

	h := markdown.ToHTML(md, nil, renderer)

	return fmt.Sprintf(htmlTemplate, css, h)
}

func getCssMarkdownFile() string {
	css, err := ioutil.ReadFile("assets/github-markdown.css")
	if err != nil {
		log.Errorln("Error read file. Error: ", err)
	}

	return string(css)
}

func getHtmlTemplate() string {
	return `
		<html>
			<head>
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<style>
					body {
						margin: 0;
					}
					
					@media (max-width: 767px) {
						.markdown-body {
							padding: 15px;
						}
					}

					.markdown-body {
						padding: 200px;
					}

					%s
				</style>
			</head>
			<body>
				<article class="markdown-body">
					%s
				</article>
			</body>
		</html>
		`
}

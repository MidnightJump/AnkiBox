package utils2

import (
	"github.com/russross/blackfriday/v2"
	"strings"
)

// 解析Markdown，提取所有三级标题及其对应内容
func extractH3HeadingsWithContent(mdContent []byte) map[string]string {
	result := make(map[string]string)
	var currentTitle string
	var contentBuilder strings.Builder

	// 解析Markdown为AST
	rootNode := blackfriday.New(blackfriday.WithExtensions(blackfriday.CommonExtensions)).Parse(mdContent)

	// 遍历AST找到所有的三级标题及其内容
	var walker func(*blackfriday.Node)
	walker = func(node *blackfriday.Node) {
		if node.Type == blackfriday.Heading && node.Level == 3 { // 发现新的三级标题
			if currentTitle != "" {
				// 存储上一个标题的内容
				result[currentTitle] = contentBuilder.String()
				contentBuilder.Reset()
			}
			// 设置新的标题
			currentTitle = string(node.FirstChild.Literal)
		} else if currentTitle != "" && node.Type != blackfriday.Paragraph { // 收集文本内容
			// 遍历 Paragraph 的子节点获取文本
			for textNode := node.FirstChild; textNode != nil; textNode = textNode.Next {
				if textNode.Type == blackfriday.Text { // 只获取 Text 类型内容
					contentBuilder.WriteString(string(textNode.Literal))
				}
			}
			//contentBuilder.WriteString("\n") // 段落换行
		}

		// 遍历子节点
		for child := node.FirstChild; child != nil; child = child.Next {
			walker(child)
		}
	}

	walker(rootNode)

	// 存储最后一个标题的内容
	if currentTitle != "" {
		result[currentTitle] = contentBuilder.String()
	}

	return result
}

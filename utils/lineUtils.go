package utils2

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strings"
)

// 用来存储解析后的标题和内容
type Section struct {
	Title   string
	Content string
}

// 解析文本，找到所有以 ### 开头的文本，提取标题和内容
func ParseMarkdown(filename string) ([]Section, error) {
	file, err := os.Open(filename) // 打开文件
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("err:%e", err)
		}
	}(file)

	var sections []Section
	var currentSection *Section

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "### ") { // 找到以 ### 开头的行
			// 如果有正在处理的 Section，先保存
			if currentSection != nil {
				sections = append(sections, *currentSection)
			}

			// 新的 Section 标题
			currentSection = &Section{
				Title:   line[4:], // 去掉 "### " 的部分
				Content: "",
			}
		} else if currentSection != nil {
			// 将行作为内容添加到当前 Section
			currentSection.Content += line + "\n"
		}
	}

	// 最后一段内容
	if currentSection != nil {
		sections = append(sections, *currentSection)
	}

	// 检查是否发生了扫描错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}

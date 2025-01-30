package utils2

import (
	"errors"
	"fmt"
	"github.com/atselvan/ankiconnect"
	"regexp"
	"sort"
	"strconv"
)

type AnkiContent struct {
	Id      int
	Title   string
	Content string
}

func RegexFind(str string) (int, string, error) {
	// 定义正则表达式
	// 9、什么是Redis持久化？Redis的持久化机制是什么？各自的优缺点？
	re := regexp.MustCompile(`(\d+)、(.+)`)

	// 进行匹配
	match := re.FindStringSubmatch(str)

	// 检查匹配结果
	if len(match) > 2 {
		number := match[1]  // 第1个捕获组（数字）
		content := match[2] // 第2个捕获组（文字）
		if num, err := strconv.Atoi(number); err == nil {
			return num, content, nil
		}
	}
	return 0, "", errors.New("未匹配到正确的格式")
}
func GetAnkiContentFromSection(sections []Section) []AnkiContent {
	var res []AnkiContent
	//for k, v := range extractH3HeadingsWithContent(mdContent) {
	//	num, title, err := RegexFind(k)
	//	if err != nil {
	//		println(err)
	//	}
	//	res = append(res, AnkiContent{
	//		Id:      num,
	//		Title:   title,
	//		Content: v,
	//	})
	//}
	for _, v := range sections {
		num, title, err := RegexFind(v.Title)
		if err != nil {
			println(err)
		}
		res = append(res, AnkiContent{
			Id:      num,
			Title:   title,
			Content: v.Content,
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Id < res[j].Id
	})
	return res
}
func AddNote(deckName string, contents []AnkiContent) error {
	client := ankiconnect.NewClient()
	decks, restErr := client.Decks.GetAll()
	if restErr != nil {
		return fmt.Errorf(restErr.Error)
	}
	hasDeck := false
	for _, v := range *decks {
		if v == deckName {
			hasDeck = true
		}
	}
	if !hasDeck {
		return fmt.Errorf("不存在该Deck：%s，请手动创建", deckName)
	}
	count := 0
	for _, v := range contents {
		if v.Content == "" || v.Content == "\n" || v.Content == "\r" || v.Title == "" {
			fmt.Printf("数据格式不规范,已忽略：%+v", v)
			continue
		}
		note := ankiconnect.Note{
			DeckName:  deckName,
			ModelName: "basicMarkdown",
			Fields: ankiconnect.Fields{
				"Front": v.Title,
				"Back":  v.Content,
			},
		}
		restErr := client.Notes.Add(note)
		if restErr != nil {
			return fmt.Errorf(restErr.Error)
		}
		count++
	}
	fmt.Printf("%v条笔记已经更新到%s中", count, deckName)
	return nil
}

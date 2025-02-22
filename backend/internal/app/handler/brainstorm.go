package handler

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type BrainstormRequest struct {
	Topic string `json:"topic" validate:"required"`
}

type BrainstormResponse struct {
	Ideas []string `json:"ideas"`
}

func BrainstormHandler(c echo.Context) error {
	var req BrainstormRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	if err := c.Validate(req); err != nil {
		return echo.ErrBadRequest
	}

	accessToken, err := c.Cookie("AccessToken")
	if err != nil {
		return echo.ErrForbidden
	}

	if accessToken.Value != "dummy_token" {
		return echo.ErrForbidden
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful assistant that generates ideas for a brainstorming ideas."},
			{Role: "user", Content: "Give me 5 brainstorming ideas for " + req.Topic},
		},
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var ideas []string
	if len(resp.Choices) > 0 {
		ideas = strings.Split(resp.Choices[0].Message.Content, "\n")
	}

	// result := "1. 教育的なイベント: 学校や地域社会で虫歯予防についてのセミナーやワークショップを開催し、虫歯への理解と認識を深める。\n\n2. ソーシャルメディアを使った啓発: TikTokやInstagramなどのプラットフォームで、虫歯予防に関する情報や日常生活のヒントを共有する。\n\n3. 絵本やキッズブックの作成: 子供たちに虫歯とその予防方法を教える楽しい絵本やキッズブックを作成します。\n\n4. 虫歯予防キャンペーン: デンタルケア製品の無料サンプルを提供することで、一般的な歯科衛生習慣の普及につなげる。\n\n5. 歯医者とのパートナーシップ: 地元の歯科医師と連携して、一般的な歯科衛生診療や教育の提供を通じて虫歯予防を広める。"

	// 空のアイデアを削除してクリーンアップ
	var cleanedIdeas []string
	for _, idea := range ideas {
		if idea != "" {
			cleanedIdeas = append(cleanedIdeas, strings.TrimSpace(idea))
		}
	}

	return c.JSON(http.StatusOK, BrainstormResponse{Ideas: cleanedIdeas})
}

package EinoCompile

import (
	"context"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

var systemPrompt = `
# Role: 健身管理及预约系统的智能助手

## Core Competencies
- 跟据用户的提问，回答相应的内容
- 可能从从知识库中查询相关内容
- 可能需要从各个tool中获取相关内容

## Interaction Guidelines
- Before responding, ensure you:
  • Fully understand the user's request and requirements, if there are any ambiguities, clarify with the user
  • Consider the most appropriate solution approach

- When providing assistance:
  • Be clear and concise
  • Include practical examples when relevant
  • Suggest improvements or next steps if applicable

- If a request exceeds your capabilities:
  • Clearly communicate your limitations, suggest alternative approaches if possible

- If the question is compound or complex, you need to think step by step, avoiding giving low-quality answers directly.

## Context Information
- Current Date: {date}
- Related Documents: |-
==== doc start ====
  {documents}
==== doc end ====
`

type ChatTemplateConfig struct {
	FormatType schema.FormatType
	Templates  []schema.MessagesTemplate
}

// newChatTemplate component initialization function of node 'ChatTemplate' in graph 'My'
func newChatTemplate(ctx context.Context) (ctp prompt.ChatTemplate, err error) {
	// TODO Modify component configuration here.
	config := &ChatTemplateConfig{
		FormatType: schema.FString,
		Templates: []schema.MessagesTemplate{
			schema.SystemMessage(systemPrompt),
			schema.MessagesPlaceholder("history", true),
			schema.UserMessage("{content}"),
		},
	}
	ctp = prompt.FromMessages(config.FormatType, config.Templates...)
	return ctp, nil
}

package agent

import (
	"context"
	"errors"
	"fmt"
	"gin-blog/internal/EinoCompile"
	"io"

	"github.com/cloudwego/eino/schema"
)

func RunAgent(ctx context.Context, id string, msg string) (*schema.StreamReader[*schema.Message], error) {

	runner, err := EinoCompile.BuildMy(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to build agent graph: %w", err)
	}

	conversation := memory.GetConversation(id, true)

	userMessage := &EinoCompile.UserMessage{
		ID:      id,
		Query:   msg,
		History: conversation.GetMessages(),
	}

	sr, err := runner.Stream(ctx, userMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to stream: %w", err)
	}

	srs := sr.Copy(2)

	go func() {
		// for save to memory
		fullMsgs := make([]*schema.Message, 0)

		defer func() {
			// close stream if you used it
			srs[1].Close()

			// add user input to history
			conversation.Append(schema.UserMessage(msg))

			fullMsg, err := schema.ConcatMessages(fullMsgs)
			if err != nil {
				fmt.Println("error concatenating messages: ", err.Error())
			}
			// add agent response to history
			conversation.Append(fullMsg)
		}()

	outer:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done", ctx.Err())
				return
			default:
				chunk, err := srs[1].Recv()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break outer
					}
				}

				fullMsgs = append(fullMsgs, chunk)
			}
		}
	}()

	return srs[0], nil
}

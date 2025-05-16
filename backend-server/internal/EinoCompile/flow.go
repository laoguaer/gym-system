package EinoCompile

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
)

// newLambda1 component initialization function of node 'ReactAgent' in graph 'My'
func newLambda1(ctx context.Context) (lba *compose.Lambda, err error) {
	// TODO Modify component configuration here.
	config := &react.AgentConfig{
		MaxStep:            25,
		ToolReturnDirectly: map[string]struct{}{}}
	chatModelIns11, err := newChatModel(ctx)
	if err != nil {
		return nil, err
	}
	config.Model = chatModelIns11
	// toolIns21, err := newTool(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	courseTool, err := newCourseTool()
	if err != nil {
		return nil, err
	}
	coachTool, err := newCoachTool()
	if err != nil {
		return nil, err
	}
	reservationTool, err := newReservationTool()
	if err != nil {
		return nil, err
	}
	bookingTool, err := newBookingTool()
	if err != nil {
		return nil, err
	}

	config.ToolsConfig.Tools = []tool.BaseTool{courseTool, coachTool, reservationTool, bookingTool}
	ins, err := react.NewAgent(ctx, config)
	if err != nil {
		return nil, err
	}
	lba, err = compose.AnyLambda(ins.Generate, ins.Stream, nil, nil)
	if err != nil {
		return nil, err
	}
	return lba, nil
}

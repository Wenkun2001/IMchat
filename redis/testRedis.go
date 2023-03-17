package mq

import (
	"IMchat/utils"
	"context"
	"fmt"
	"testing"
	"time"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

// TestPublish 测试发布消息到redis
func TestPublish(t *testing.T) {
	msg := "current time: " + time.Now().Format("15:04:05")
	err := utils.Publish(ctx, utils.PublishKey, msg)
	if err != nil {

		fmt.Println(err)
	}
}

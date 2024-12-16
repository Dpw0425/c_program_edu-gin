package generator

import (
	"c_program_edu-gin/pkg/logger"
	"github.com/bwmarrin/snowflake"
)

func IDGenerator() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Panic(err)
	}

	return node.Generate().Int64()
}

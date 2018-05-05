package pipelines

import (
	"github.com/infinitbyte/framework/core/pipeline"
	"github.com/infinitbyte/framework/core/util"
	"time"
)

type LoggingJoint struct {
}

func (joint LoggingJoint) Name() string {
	return "logging"
}

func (joint LoggingJoint) Process(c *pipeline.Context) error {

	sql := c.MustGetString(sqlKey)

	util.FileAppendNewLine("log/executed_sql.log", "")
	util.FileAppendNewLine("log/executed_sql.log", time.Now().String())
	util.FileAppendNewLine("log/executed_sql.log", sql)
	util.FileAppendNewLine("log/executed_sql.log", "")

	return nil
}

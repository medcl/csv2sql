/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pipelines

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/cihub/seelog"
	"github.com/infinitbyte/framework/core/pipeline"
	"strings"
)

var rowFormat pipeline.ParaKey = "row_format"
var sheetName pipeline.ParaKey = "sheet_name"
var columnName pipeline.ParaKey = "column_name"
var dataFromIndex pipeline.ParaKey = "data_start_from_index"

type ReadCsvJoint struct {
}

func (joint ReadCsvJoint) Name() string {
	return "read_csv"
}

func (joint ReadCsvJoint) Process(c *pipeline.Context) error {

	log.Debug(c.Data)
	templates := c.MustGetStringArray(rowFormat)
	log.Debug("row templates: ", templates)

	xlsx, err := excelize.OpenFile("../test/Book1.xlsx")
	if err != nil {
		log.Error(err)
		return err
	}
	sheetMap := xlsx.GetSheetMap()
	log.Debug("sheets: ", sheetMap)

	colNames := c.MustGetStringArray(columnName)
	dataOffset := c.MustGetInt(dataFromIndex)

	rows := xlsx.GetRows(c.MustGetString(sheetName))
	for offset, row := range rows {
		if offset < dataOffset {
			log.Debugf("%v < data offset: %v,　ignore", offset, dataOffset)
			continue
		}

		colMap := map[string]string{}

		for k, colCell := range row {
			colName := colNames[k]
			colMap[colName] = colCell
			log.Debug("row:",offset, ": ", colName, "-", k, "-", colCell)
		}

		for _, x := range templates {
			line:=x
			log.Debug("template:",line)
			for k, v := range colMap {
				log.Debug(fmt.Sprintf("<{%v: }>",k),",", formatString(v))
				line = strings.Replace(line, fmt.Sprintf("<{%v: }>", k), formatString(v), -1)
			}
			log.Debug(line)
		}
	}

	return nil
}

func formatString(str string) string {
	str = strings.Replace(str, "\"", "", -1)
	str=fmt.Sprintf("'%s'",str)
	return str
}

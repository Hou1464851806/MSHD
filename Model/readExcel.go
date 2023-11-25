package Model

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ReadFile(filePath string) (datas [][]Location) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	sheets := []string{"region_code", "region_code(2)", "region_code(3)",
		"region_code(4)", "region_code(5)", "region_code(6)",
		"region_code(7)", "region_code(8)", "region_code(9)"}
	for _, sheet := range sheets {
		var data []Location
		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, row := range rows {
			tuples := Location{}
			tuples.Number = row[0]
			tuples.Province = row[1]
			tuples.City = row[2]
			tuples.County = row[3]
			tuples.Town = row[4]
			tuples.Village = row[5]
			//log.Println("tuples", tuples)
			data = append(data, tuples)
		}
		datas = append(datas, data)
		log.Println("Sheet:", sheet+"complete")
	}
	return datas
}

func InsertDataToLocationdb() {
	for _, data := range ReadFile("region_code.xlsx") {
		//log.Println(data)
		db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
		if err != nil {
			log.Println(err)
		}
		results := db.CreateInBatches(&data, 999)
		log.Println(results.Error)
		log.Println(results.RowsAffected)
	}
}

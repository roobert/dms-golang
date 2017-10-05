package main

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//func upsertClientData(c Client) error {
//	fmt.Println(c.Name, c.TimeStamp, c.Bitmap)
//
//	// fake bitmap for now
//	bitmap := make([]bool, 99)
//
//	// read bitmap from database
//	fmt.Println(bitmap)
//
//	// if bitmap exists update it
//	// UPDATE BITMAP
//
//	// else create new bitmap and insert
//	// FIXME: convert bitmap from bool slice to int
// date := c.TimeStamp.Format("2006-01-02")
//	stmt, _ := db.Prepare("INSERT INTO data (site, date, bitmap) VALUES (?, ?, ?)")
//	_, err := stmt.Exec(c.Name, c.TimeStamp, bitmap)
//
//	return err
//}

func fetchBitmap(c Client) ([]bool, error) {
	var bitmap []bool

	err := checkForMultipleRecords(c)
	checkErr(err)

	if rowExists(c) {
		bitmap = selectBitmapFromDB(c)
	} else {
		bitmap = newBitmap()
	}

	// return bitmap
	return bitmap, err
}

// update client in place?
func selectBitmapFromDB(c Client) []bool {
	//query := fmt.Sprintf("SELECT * FROM data WHERE 'name' == '%s' and 'date' == '%s')", c.Name, c.Date)
	stmt, _ := db.Prepare("SELECT * FROM data WHERE 'name' == ? AND 'date' == ?")
	rows, err := stmt.Query(c.Name, c.TimeStamp)
	checkErr(err)

	var name string
	var timeStamp time.Time
	var bitmap []bool

	for rows.Next() {
		err = rows.Scan(&name, &timeStamp, &bitmap)
		checkErr(err)
	}

	return bitmap
}

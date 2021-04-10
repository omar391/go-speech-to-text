package repository

import (
	"stt-service/models"
	"stt-service/utils"
)

var (
	parent_table string = "stt_data"
	fts_table    string = parent_table + "_fts"
)

// Migrate the schema on startup
func init() {
	db := utils.OpenSQLiteDB()
	db.AutoMigrate(models.STTData{})

	//Create a full text search virtual table for text audio data

	if !db.Migrator().HasTable(fts_table) {

		create_fts := `CREATE VIRTUAL TABLE ` + fts_table + ` USING FTS4(
			text,
			original_audio_file_name,
			user_id UNINDEXED, 
			original_audio_file_path UNINDEXED, 
			content='` + parent_table + `', 
			--content_id='id' 
		);
		
		CREATE TRIGGER ` + parent_table + `_insert AFTER INSERT ON ` + parent_table + `
			BEGIN
				INSERT INTO ` + fts_table + ` (rowid, text, original_audio_file_name)
				VALUES (new.id, new.text, new.original_audio_file_name);
			END;

		CREATE TRIGGER ` + parent_table + `_delete AFTER DELETE ON ` + parent_table + `
			BEGIN
				INSERT INTO ` + fts_table + ` (` + fts_table + `, rowid, text, original_audio_file_name)
				VALUES ('delete', old.id, old.text, old.original_audio_file_name);
			END;

		CREATE TRIGGER ` + parent_table + `_update AFTER UPDATE ON ` + parent_table + `
			BEGIN
				INSERT INTO ` + fts_table + ` (` + fts_table + `, rowid, text, original_audio_file_name)
				VALUES ('delete', old.id, old.text, old.original_audio_file_name);
				INSERT INTO ` + fts_table + ` (rowid, text, original_audio_file_name)
				VALUES (new.id, new.text, new.original_audio_file_name);
			END;`

		db.Exec(create_fts)
	}
}

// Create a new AUDIO file row
func CreateNewAudioField(audio_text models.STTData) {
	db := utils.OpenSQLiteDB()
	db.Create(audio_text)
}

//get count of audio data rows
func GetTotalRowCount() int64 {
	db := utils.OpenSQLiteDB()
	var count int64
	db.Model(&models.STTData{}).Count(&count)

	return count
}

//get all audio data by limit, offset
func GetAudioData(page_no int, limit int) []map[string]interface{} {
	db := utils.OpenSQLiteDB()

	// Get all records
	var results []map[string]interface{}
	db.Model(models.STTData{}).Offset(page_no * limit).Limit(limit).Find(&results)

	return results
}

//get all audio data by limit, offset, and text
func FilterAudioData(page_no int, limit int, text_to_find string) []map[string]interface{} {
	db := utils.OpenSQLiteDB()

	// Get all filtered records
	var results []map[string]interface{}
	offset := page_no * limit
	db.Model(models.STTData{}).Raw("SELECT * FROM ? WHERE ? MATCH '?' OFFSET ? LIMIT ?", fts_table, fts_table, text_to_find, offset, limit).Find(&results)

	return results
}

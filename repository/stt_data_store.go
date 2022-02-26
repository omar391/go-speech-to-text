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
	db.AutoMigrate(&models.STTData{})

	// Create a full text search virtual table for text audio data

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
func CreateNewAudioField(audio_text *models.STTData) {
	db := utils.OpenSQLiteDB()
	db.Create(audio_text)
}

//get count of audio data rows
func GetTotalDataRowCount(user_id uint) uint {
	db := utils.OpenSQLiteDB()
	var count int64
	db.Model(&models.STTData{}).Where("user_id = ?", user_id).Count(&count)

	return uint(count)
}

// get all audio data by limit, offset
func GetAllAudioData(offset int, limit int, user_id uint) []map[string]interface{} {
	db := utils.OpenSQLiteDB()

	// Get all records
	var results []map[string]interface{}
	//db.Model(models.STTData{}).Select("original_audio_file_name", "original_audio_file_path", "text").Offset(offset).Limit(limit).Where("user_id = ?", user_id).Find(&results)
	db.Model(models.STTData{}).Select("original_audio_file_name", "original_audio_file_path", "text").Where("user_id = ?", user_id).Find(&results)

	return results
}

// get all audio data by limit, offset, and text
func FilterAudioData(offset int, limit int, text_to_find string, user_id uint) []map[string]interface{} {
	db := utils.OpenSQLiteDB()

	// Get all filtered records
	var results []map[string]interface{}
	//db.Model(models.STTData{}).Raw("SELECT original_audio_file_name, original_audio_file_path, text FROM stt_data_fts WHERE user_id = ? AND stt_data_fts MATCH ? LIMIT 100 OFFSET 0", user_id, text_to_find, limit, offset).Find(&results)
	db.Model(models.STTData{}).Raw("SELECT original_audio_file_name, original_audio_file_path, text FROM stt_data_fts WHERE user_id = ? AND stt_data_fts MATCH ?", user_id, text_to_find).Find(&results)

	return results
}

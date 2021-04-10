package models

type STTData struct {
	ID                    uint
	UserId                uint   `gorm:"type:INTEGER NOT NULL;"`
	OriginalAudioFileName string `gorm:"type:TEXT NOT NULL; check:original_audio_file_name <> ''"`
	OriginalAudioFilePath string `gorm:"type:TEXT NOT NULL; check:original_audio_file_path <> ''"`
	Text                  string `gorm:"type:TEXT NULL; "`
	User                  User   `gorm:"foreignKey:user_id; constraint:OnDelete:CASCADE;"`
}

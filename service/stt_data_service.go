package service

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"stt-service/conf"
	"stt-service/models"
	"stt-service/repository"
	"time"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/floostack/transcoder/ffmpeg"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func getOffset(page_no int, limit int, user_id uint) int {
	// Get total rows for current session user
	count := repository.GetTotalDataRowCount(user_id)
	offset := (page_no - 1) * limit
	if uint(offset) > count {
		return -1
	}
	return offset
}

// Get all data by pagination
func GetAudioDataByPage(page_no int, limit int, user_id uint) []map[string]interface{} {
	offset := getOffset(page_no, limit, user_id)
	if offset == -1 {
		return nil
	}

	return repository.GetAllAudioData(offset, limit, user_id)
}

// Filter data with query string and pagination
func FilterAudioData(page_no int, limit int, text_to_find string, user_id uint) []map[string]interface{} {
	offset := getOffset(page_no, limit, user_id)
	if offset == -1 {
		return nil
	}

	return repository.FilterAudioData(offset, limit, text_to_find, user_id)
}

// Return audio transcript
func SaveTextFromAudio(orig_file_name string, orig_file_full_path string, user_id uint, is_keep_file bool) (string, error) {
	new_file_path, err := convertAudio(orig_file_name, orig_file_full_path)
	var text string
	if err != nil {
		text = "File conversion failed!"

	} else {
		// time for new file to be unlocked from previous conversion process
		time.Sleep(1 * time.Millisecond)

		// Process audio transcript in the cloud
		text, err = processAudioInGCP(new_file_path)
		if err != nil {
			text = "File cloud transcription task failed!"

		} else {
			text = strings.TrimRight(text, "\n")
		}
	}

	go func() {
		//remove the file if not asked to be stored
		if !is_keep_file {
			os.Remove(new_file_path)
			new_file_path = "File is deleted!"
		}

		if err == nil {
			//save audio transcript
			repository.CreateNewAudioField(&models.STTData{
				UserId:                user_id,
				OriginalAudioFileName: orig_file_name,
				OriginalAudioFilePath: new_file_path,
				Text:                  text,
			})
		}

		//delete original raw audio file
		os.Remove(orig_file_full_path)
	}()

	return text, err
}

// Process local audio file to Google Cloud Speech API
func processAudioInGCP(filename string) (string, error) {
	//set env variables
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", conf.Config.GOOGLE_APPLICATION_CREDENTIALS)

	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Send the contents of the audio file with the encoding and
	// and sample rate information to be transcripted.
	req := &speechpb.LongRunningRecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_FLAC,
			SampleRateHertz: 24000,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	}

	op, err := client.LongRunningRecognize(ctx, req)
	if err != nil {
		return "", err
	}
	resp, err := op.Wait(ctx)
	if err != nil {
		return "", err
	}

	// Print the results.
	var buffer bytes.Buffer
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			buffer.WriteString(alt.Transcript + "\n")
		}
	}
	return buffer.String(), nil
}

// Convert media to FLAC audio format of max 60 sec.
func convertAudio(original_name string, file_to_convert string) (string, error) {
	//raw command: ffmpeg -i ./data/voice.mp4 -vn -y -t 5 -ar 24000 -ac 1 -compression_level 12 ./data/out.flac
	d_true := true
	duration := "60"
	audio_rate := 24000
	audio_channels := 1
	compression_level := 12
	codec := "flac"

	opts := ffmpeg.Options{
		SkipVideo:        &d_true,
		Overwrite:        &d_true,
		Duration:         &duration,
		AudioRate:        &audio_rate,
		AudioChannels:    &audio_channels,
		CompressionLevel: &compression_level,
		AudioCodec:       &codec,
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   "ffmpeg",
		FfprobeBinPath:  "ffprobe",
		ProgressEnabled: false,
	}

	f, fr := ioutil.TempFile(conf.Config.DATA_DIR, original_name+"_*."+codec)
	if fr != nil {
		return "", fr
	}

	output_file := f.Name()

	_, err := ffmpeg.
		New(ffmpegConf).
		Input(file_to_convert).
		Output(output_file).
		WithOptions(opts).
		Start(opts)

	if err != nil {
		return "", err
	} else {
		f.Sync()
		return output_file, nil
	}
}

package ocrworker

import (
	"encoding/json"
	"testing"

	"io/ioutil"

	"github.com/couchbaselabs/go.assert"
	"github.com/couchbaselabs/logg"
)

func TestTesseractEngineWithRequest(t *testing.T) {

	engine := TesseractEngine{}
	bytes, err := ioutil.ReadFile("docs/testimage.png")
	assert.True(t, err == nil)

	cFlags := make(map[string]interface{})
	cFlags["tessedit_char_whitelist"] = "0123456789"

	ocrRequest := OcrRequest{
		ImgBytes:   bytes,
		EngineType: ENGINE_TESSERACT,
		EngineArgs: cFlags,
	}

	assert.True(t, err == nil)
	result, err := engine.ProcessRequest(ocrRequest)
	assert.True(t, err == nil)
	logg.LogTo("TEST", "result: %v", result)

}

func TestTesseractEngineWithJson(t *testing.T) {

	testJsons := []string{}
	testJsons = append(testJsons, `{"engine":"tesseract_exec"}`)
	testJsons = append(testJsons, `{"engine":"tesseract_exec", "engine_args":{}}`)
	testJsons = append(testJsons, `{"engine":"tesseract_exec", "engine_args":null}`)
	testJsons = append(testJsons, `{"engine":"tesseract_exec", "engine_args":{"config_vars":{"tessedit_char_whitelist":"0123456789"}, "psm":"0"}}`)

	for _, testJson := range testJsons {
		logg.LogTo("TEST", "testJson: %v", testJson)
		ocrRequest := OcrRequest{}
		err := json.Unmarshal([]byte(testJson), &ocrRequest)
		assert.True(t, err == nil)
		bytes, err := ioutil.ReadFile("docs/testimage.png")
		assert.True(t, err == nil)
		ocrRequest.ImgBytes = bytes
		engine := NewOcrEngine(ocrRequest.EngineType)
		result, err := engine.ProcessRequest(ocrRequest)
		logg.LogTo("TEST", "err: %v", err)
		assert.True(t, err == nil)
		logg.LogTo("TEST", "result: %v", result)

	}

}

func TestNewTesseractEngineArgs(t *testing.T) {
	testJson := `{"engine":"tesseract_exec", "engine_args":{"config_vars":{"tessedit_char_whitelist":"0123456789"}, "psm":"0"}}`
	ocrRequest := OcrRequest{}
	err := json.Unmarshal([]byte(testJson), &ocrRequest)
	assert.True(t, err == nil)
	engineArgs, err := NewTesseractEngineArgs(ocrRequest)
	assert.True(t, err == nil)
	assert.Equals(t, len(engineArgs.configVars), 1)
	assert.Equals(t, engineArgs.configVars["tessedit_char_whitelist"], "0123456789")
	assert.Equals(t, engineArgs.pageSegMode, "0")

}

func TestTesseractEngineWithFile(t *testing.T) {

	engine := TesseractEngine{}
	engineArgs := TesseractEngineArgs{}
	result, err := engine.processImageFile("docs/testimage.png", engineArgs)
	assert.True(t, err == nil)
	logg.LogTo("TEST", "result: %v", result)

}

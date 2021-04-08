package srv

import (
	"os"
	"stt-service/mock"
	"testing"
)

//TestMain setup mock data
func TestMain(m *testing.M) {
	mock.SetupAppConfig()

	code := m.Run()

	os.Exit(code)
}

//TestSearchAPI test the http response of online search api
func TestSearchAPI(t *testing.T) {

}

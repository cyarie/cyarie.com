package backend
import (
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/cyarie/cyarie.com/backend/router"
	"github.com/cyarie/cyarie.com/backend/settings"
	"fmt"
)

// Let's write two simple setup and teardown functions. Setup() will return an *httptest.Server setup using the
// router from the backend API. Teardown will close it after being deferred.
func Setup() *httptest.Server {
	context := &settings.AppContext{}
	router := router.ApiRouter(context)

	server := httptest.NewUnstartedServer(router)

	return server
}

func Teardown(server *httptest.Server) {
	server.Close()
}

func TestIndexHandler(t *testing.T) {
	server := Setup()
	server.Start()
	defer Teardown(server)

	req, err := http.NewRequest("GET", fmt.Sprint(server.URL), nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected a 200, received: %d", res.StatusCode)
	}
}

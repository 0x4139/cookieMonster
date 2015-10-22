package cookieMonster
import (
	"testing"
	"log"
	"net/http"
)

type DummyStorage struct {
}

func (storage DummyStorage) Get(id string) []byte {
	log.Printf("Fetching cookies for %s\n", id)
	return nil
}
func (storage DummyStorage) Set(id string, cookie []byte) {
	log.Printf("Setting cookies for id=%s data=%s\n", id, string(cookie))

}


func Test_cookieMonster(t *testing.T) {
	jar := CookieMonster{
		id:"dumb",
		storage:DummyStorage{},
	}
	client := http.Client{Jar: jar}
	_, err := client.Get("http://dubbelboer.com/302cookie.php")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("DONE!")
}

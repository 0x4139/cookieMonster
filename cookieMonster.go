package cookieMonster
import (
	"net/http"
	"net/url"
	"encoding/json"
)
// represents a basic interface to get the serialized string based on an interface so it can be used
// with different storage types
type CookieMonsterStorage interface {
	Get(id string) []byte
	Set(id string, cookie []byte)
}
// our cookie jar
// the identifier represents the way we uniquely represent the cookie in the db/file for example
type CookieMonster struct {
	Storage CookieMonsterStorage
	ID string
}

func (jar CookieMonster)  SetCookies(u *url.URL, cookies []*http.Cookie) {
	serializedCookies, err := json.Marshal(cookies)
	if err != nil {
		panic(err)
	}
	jar.Storage.Set(jar.ID,serializedCookies)
}

func (jar CookieMonster) Cookies(u *url.URL) []*http.Cookie {
	var cookies []*http.Cookie

	serializedCookies:=jar.Storage.Get(jar.ID)
	if serializedCookies==nil || len(serializedCookies)==0{
		return []*http.Cookie{}
	}
	err:=json.Unmarshal(serializedCookies,&cookies)
	if err!=nil{
		panic(err)
	}
	return cookies
}

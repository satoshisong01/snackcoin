package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/sks8982/snackcoin/blockchain"
)

const (
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) { //writer로 사용자에게 대답
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data) //Execute에 writer , data 필요
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil) //Get이면 아무 데이터 없이 add template을 씀
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.Blockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home) //route 생성  rw writer 유저에게 보내고싶은 데이터 작성
	http.HandleFunc("/add", add)
	fmt.Printf("http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil)) //go는 콘솔에러가 나오지않아 별도로 추가 에러처리
}

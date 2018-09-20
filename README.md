# chi-clova
This package is an extension for developing LINE's Clova Extension while using go-chi.
If you want to build the REST API on AWS Lambda in a single binary, using this package makes implementation easier.

## Additional information
When using go-chi, it is convenient to use the following extensions together.
- [chi-extension](https://github.com/WhiteRaven777/chi-extension)

# How to use.
go get -u github.com/WhiteRaven777/chi-clova

## code sample
``` go
import (
	mw "github.com/WhiteRaven777/chi-clova/middleware"
	"github.com/WhiteRaven777/chi-clova/clova"
	"github.com/go-chi/chi"
)

r := chi.NewRouter()

r.Use(mw.New("com.example.my_extension").Middleware)
r.Post("/callback", func(w http.ResponseWriter, r *http.Request) {
	message, err := clova.Clova(r).ParseRequest(r)
	if err != nil {
		log.Printf("invalid request")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var response *cek.ResponseMessage
	switch request := message.Request.(type) {
	case *cek.IntentRequest:
		switch request.Intent.Name {
		case "Clova.GuideIntent":
			response = cek.NewResponseBuilder().
				OutputSpeech(
					cek.NewOutputSpeechBuilder().
						AddSpeechText("話しかけてください", cek.SpeechInfoLangJA).
						Build()).
				Build()
		}
	case *cek.LaunchRequest:
		response = cek.NewResponseBuilder().
			OutputSpeech(
				cek.NewOutputSpeechBuilder().
					AddSpeechText("起動しました", cek.SpeechInfoLangJA).
					Build()).
			Build()
	}
	if response != nil {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
})
```
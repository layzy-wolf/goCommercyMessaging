package transport

import (
	"ApiGateway/config"
	"ApiGateway/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

var verify *Verify

func Handler(cfg config.Cfg) *mux.Router {
	r := mux.NewRouter()

	// Инициализация обработчиков
	authService := NewAuthHandler(cfg)
	verify = NewVerify(cfg)

	accountService := MakeAccountHandler(cfg)
	chatService := service.NewChatService(cfg)
	chatHandler := NewChatHandler(cfg)

	gChatHandler := NewGroupHandler(cfg)

	// Определение маршрутов для маршрутизатора
	r.Methods(http.MethodPost, http.MethodOptions).Path("/register").Handler(accountService.Register)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/login").Handler(authService.Login)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/verify").Handler(authService.Verify)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/remove").Handler(accountService.Remove)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/list").Handler(accountService.List)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/search").Handler(accountService.Search)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/add/contact").Handler(accountService.AddContact)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/remove/contact").Handler(accountService.RemoveContact)

	r.Methods(http.MethodPost, http.MethodOptions).Path("/show").Handler(chatHandler.Show)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/chats").Handler(chatHandler.Chats)

	r.Methods(http.MethodPost, http.MethodOptions).Path("/group/register").Handler(gChatHandler.Register)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/group/add").Handler(gChatHandler.AddToGroup)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/group/code").Handler(gChatHandler.GetCode)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/group/remove").Handler(gChatHandler.RemoveFromGroup)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/group/get").Handler(gChatHandler.GetMembers)

	// Создание ws подключения
	r.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		_, user, err := verify.Verify(r)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		r.Header.Set("user", user)
		chatService.ForwardMessage(w, r)
	})

	// Определение Middleware для разрешения CORS
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	return r
}

package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"html/template"
)

func main() {

	// Iniciando estrutura de validação
	StartValidator()

	// Iniciando serviço web
	StartServer()
}

func StartServer() {

	r := mux.NewRouter()

	// ------------------------- Entidade Players -------------------------------

	r.HandleFunc("/players/register", RegisterPlayer).Methods("POST")

	r.HandleFunc("/players/search/all", SearchAllPlayers).Methods("GET")

	r.HandleFunc("/players/search", SearchPlayer).Methods("GET")

	r.HandleFunc("/players/edit", EditPlayer).Methods("PUT")

	r.HandleFunc("/players/delete", DeletePlayer).Methods("DELETE")

	r.HandleFunc("/players/search/all/deleted", SearchAllDeletedPlayers).Methods("GET")

	// ------------------------ Entidade Matches ----------------------------------

	r.HandleFunc("/matches/register", RegisterMatch).Methods("POST")

	r.HandleFunc("/matches/search/all", SearchMatch).Methods("GET")

	r.HandleFunc("/matches/search/all", SearchAllMatches).Methods("GET")

	r.HandleFunc("/matches/delete", DeleteMatch).Methods("DELETE")

	r.HandleFunc("/matches/search/all/deleted", SearchAllDeletedMatches).Methods("GET")

	// ----------------------- Entidade Championships ----------------------------

	r.HandleFunc("/championships/register", RegisterChampionship).Methods("POST")

	r.HandleFunc("/championships/search/all", SearchAllChampionships).Methods("GET")

	r.HandleFunc("/championships/search", SearchChampionship).Methods("GET")

	r.HandleFunc("/championships/edit", EditChampionship).Methods("PUT")

	r.HandleFunc("/championships/delete", DeleteChampionship).Methods("DELETE")

	// ----------------------- Entidade Team ----------------------------

	r.HandleFunc("/team/register", RegisterTeam).Methods("POST")

	r.HandleFunc("/team/search", SearchTeam).Methods("GET")

	r.HandleFunc("/team/search/all", SearchAllTeams).Methods("GET")

	// ----------------------- Front End ----------------------------

	r.PathPrefix("/tmpl").Handler(http.StripPrefix("/tmpl", http.FileServer(http.Dir("./tmpl"))))

	r.HandleFunc("/index", IndexHandler)

	r.HandleFunc("/dashboard", DashboradHandler)

	r.HandleFunc("/player", PlayerHandler)

	r.HandleFunc("/register_player", RegisterPlayerHandler)

	r.HandleFunc("/team", TeamHandler)

	r.HandleFunc("/champ", ChampHandler)

	r.HandleFunc("/manager_champ", ManagerChampHandler)

	r.HandleFunc("/match", MatchHandler)

	r.HandleFunc("/ranking", RankingHandler)

	r.HandleFunc("/ranking_team", RankingTeamHandler)

	r.HandleFunc("/agenda", AgendaHandler)

	http.ListenAndServe(SERVER_HOST, r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
	w.WriteHeader(http.StatusOK)

	email := r.FormValue("email")
	password := r.FormValue("pass")

	p := Player{Email: email, Password: password}
	tmpl.Execute(w, p) // merge.
}

func DashboradHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/dashboard.html"))
	w.WriteHeader(http.StatusOK)

	data := map[string]string{}

	tmpl.Execute(w, data) // merge.
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/player.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func RegisterPlayerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/register_player.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/team.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func ChampHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/champ.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func ManagerChampHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/manager_champ.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/match.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func RankingHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/ranking.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func RankingTeamHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/ranking_team.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func AgendaHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("tmpl/agenda.html"))
	data := map[string]string{}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

package forum

type Discussion struct {
	id    int
	sujet string
	mess  []Message
}

func requestSayAllDiscussion(w http.ResponseWriter, r *http.Request) {
	for disc := range discussions {
		e, err := json.MarshalIndent(discussions[disc].sujet, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(e))
	}
}

func requestSayDiscussion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if (indice - 1) < len(discussions) {
		fmt.Println(discussions[indice-1].mess)
		e, err := json.MarshalIndent(discussions[indice-1].sujet, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(e))
		return
		/*for m := range discussions[indice-1].mess {
			fmt.Fprintln(w, string(discussions[indice-1].mess[m].id))
		}*/
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Cette discussion n'existe pas"))
}

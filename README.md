# Bookings and Reservations

- Built in Go version 1.15
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)

soda migration sql ceva pompos
unit , integrity testing with a test database repository
used channels to send emails(posibil using goroutines and channels) -- cred ca ceea ce am facut noi e un local mail server

problemele:
cand fac o rezervare pe make-reservation si e ceva gresit nu se salveaza in session startdate si enddate ca la firstname sau email etc
tot aici,cand ma duc pe summary nu am startdate si enddate din session
si pe pagina search availability cand nu am nici o camera disponibila sa imi apara eroarea sus(deja e facuta) DAR sa imi ramana datele pe pagina cand se da refresh
cand sunt pe make-reservation si fac o reservare gresita se pastreaza valorile din formular dar cele din arrival date departure date extrase din sesiune se pierd
o problema, pot insera rezervari la infinit cand durata e de o singura zi, adica 23-23 februarie si nu are sens,trebuie sa pun minim o noapte

in bookroom ma gandesc ca userul poate modifica url acela cu query si sa pune orice data vrea el. cred ca ar mai trb ceva securizare sau cum zice chatu
// Check room availability again (security check)
available, err := m.DB.SearchAvailabilityByDatesByRoomID(startDate, endDate, roomID)
if err != nil {
helpers.ServerError(w, err)
return
}
if !available {
http.Error(w, "Room is not available for the selected dates", http.StatusConflict)
return
}
de gandit

sa repar testele cel putin de la handlers sa verifice toate testcase-urile deja scrise

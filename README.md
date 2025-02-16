# Bookings and Reservations

- Built in Go version 1.15
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)

termen limita de facut pana pe 20 inclusiv pana sunt multumit

password si alte datele stocate in Server Side Session
soda migration sql ceva pompos
unit , integrity testing with a test database repository
used channels to send emails(posibil using goroutines and channels) -- cred ca ceea ce am facut noi e un local mail server

problemele:
DONE cand fac o rezervare pe make-reservation si e ceva gresit nu se salveaza in session startdate si enddate ca la firstname sau email etc
DONE tot aici,cand ma duc pe summary nu am startdate si enddate din session
DONE si pe pagina search availability cand nu am nici o camera disponibila sa imi apara eroarea sus(deja e facuta) DAR sa imi ramana datele pe pagina cand se da refresh
DONE sa repar testele cel putin de la handlers sa verifice toate testcase-urile deja scrise
DONE o problema, pot insera rezervari la infinit cand durata e de o singura zi, adica 23-23 februarie si nu are sens,trebuie sa pun minim o noapte

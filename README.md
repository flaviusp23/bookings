# Bookings and Reservations

- Built in Go version 1.15
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)

soda migration sql ceva pompos

problemele:
cand fac o rezervare pe make-reservation si e ceva gresit nu se salveaza in session startdate si enddate ca la firstname sau email etc
tot aici,cand ma duc pe summary nu am startdate si enddate din session
si pe pagina search availability cand nu am nici o camera disponibila sa imi apara eroarea sus(deja e facuta) DAR sa imi ramana datele pe pagina cand se da refresh

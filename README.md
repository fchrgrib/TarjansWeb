## Cara Menjalankan Web
### Dengan Docker
- pastikan anda memiliki docker desktop pada device anda
- buka docker desktop
- pull link github ini
- masuk pada folder utama dengan menggunakan terminal
- lakukan command `docker compose build`
- setelah itu lakukan command `docker compose up -d`
- frontend akan berjalan pada `http://localhost:3000` dan backend pada `http://localhost:8080`

### Tanpa Docker
- masuk pada folder backend terlebih dahulu
- pastikan anda sudah menginstall bahasa pemrograman go yang terbaru
- buka folder dengan terminal
- lakukan command `go mod tidy`
- setelah itu jalankan program dengan commad `go run main.go`
- setelah backend berjalan kita beralih pada folder frontend
- pastikan anda sudah menginstall node js versi terbaru
- pada folder frontend bukalah dengan terminal
- lalu masukkan command `npm install` setelah itu `npm start`
- frontend akan berjalan pada `http://localhost:3000` dan backend pada `http://localhost:8080`

## Cara Menggunakan Web
1. pertama-tama pilihlah terlebih dahulu ingin menampilkan `SCC` atau `Bridge` pada toggle button (**WAJIB**)
2. buatlah graph dengan format file txt. untuk cara membuat graph txt anda bisa melihat file tc1.txt pada folder test
3. setelah graph sudah dibuat tekan tombol `choose file` untuk memilih graph yang sudah dibuat.
4. lalu tekan tombol `POST` untuk mendapatkan hasilnya
5. jika anda ingin mengganti type tampilan dari `SCC` ke `Bridge` atau sebaliknya. Anda DIWAJIBKAN untuk menekan tombol `CLEAR GRAPH` terlebih dahulu  (**WAJIB**)

## Tentang Tarjans
- Kompleksitas dari tarjans adalah O(V+E)
- Tarjans mengecek apakah low[v] > disc[u] atau tidak jika iya maka (u, v) adalah bridge
- Back edge ditentukan jika semisal ada edge (A,B) jika A adalah ancestor dari B make ia adalah Back edge

## Framework
- Frontend : ReactJs
- Backend :Gin

## Library
- `react-graoh-vis` : untuk menampilkan graph pada front end
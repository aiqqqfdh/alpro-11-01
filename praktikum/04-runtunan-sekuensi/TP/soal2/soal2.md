Pertanyaan:
1. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi?
2. Apa output yang dihasilkan oleh program?
3. Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan:
◦ Kondisi 1: Apakah kondisi x > 5 benar? Jika ya, apa yang terjadi selanjutnya?
◦ Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Bagaimana hal ini
memengaruhi nilai result?
◦ Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Apa yang
terjadi?
 Kondisi 4: Bagaimana kondisi !(x < 15 && y < 10) dievaluasi? Apa
dampaknya pada result?

Jawab :
1.25
2.menampilkan variabel "result"
3.
◦ benar, menjalankan kondisi 
 if x > 5 {
	if y < 10 {
		result = x + y
	} else {
		result = x - y
	}
 }

◦ benar, 
 if z > 10 && x == 10 {
 result += z
 } else {
 result = z - x
 }
 karena dia menambah nilai result yg telah ada dengan nilai di variabel z

◦ kondisi "x == 10" benar, menjalankan
 result += 5

◦ melihat apakah nilai "x < 15 dan nilai y < 10" jika nilai x dan y sama-sama benar, maka akan menjalankan kondisi  result += 10 kalau tidak akan menjalankan result -= 10. jadi kalau kondisi benar akan ditambah 10 jika salah akan dikurang 10


package main

import "fmt"

const MaksData = 100

// 1. Struct paling dalam
type Data struct {
	RefleksiDiri string
	Skor         int
}

// 2. Struct isi tiap user
type TipeC struct {
	SelfAssessment     [10]Data
	IDPengguna         int64
	SkorSelfAssessment int
	Date               string
	n                  int
}

// 3. Struct terluar
type TipeB struct {
	C [MaksData]TipeC
	n int
}

var B TipeB

func main() {
	var pilihan, sub int
	for {
		fmt.Println("\n=======================================")
		fmt.Println("   SISTEM MANAJEMEN KESEHATAN MENTAL   ")
		fmt.Println("=======================================")
		fmt.Println("1. Input Data Assessment")
		fmt.Println("2. Cetak Laporan Data")
		fmt.Println("3. Kelola Data (Edit / Hapus)")
		fmt.Println("4. Cari Data (Sequential / Binary)")
		fmt.Println("5. Urutkan berdasarkan Skor (Desc)")
		fmt.Println("6. Urutkan berdasarkan Tanggal (Asc)")
		fmt.Println("7. Exit Program")
		fmt.Println("=======================================")
		fmt.Print("Pilih menu (1-7): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData(&B)
		case 2:
			cetakData(B)
		case 3:
			fmt.Print("-> Aksi: 1.Edit | 2.Hapus ? "); fmt.Scan(&sub)
			if sub == 1 { editData(&B) } else if sub == 2 { hapusData(&B) } else { fmt.Println("Pilihan salah!") }
		case 4:
			fmt.Print("-> Metode: 1.Sequential | 2.Binary ? "); fmt.Scan(&sub)
			if sub == 1 { cariSeq(B) } else if sub == 2 { cariBin(B) } else { fmt.Println("Pilihan salah!") }
		case 5:
			sortSkor(&B)
		case 6:
			sortTanggal(&B)
		case 7:
			fmt.Println("Program selesai. Jangan lupa istirahat!")
			return
		default:
			fmt.Println("Pilihan menu tidak valid!")
		}
	}
}

// --- HELPER MENCARI INDEKS ---
func cariIndeks(B TipeB, id int64) int {
	for i := 0; i < B.n; i++ {
		if B.C[i].IDPengguna == id { return i }
	}
	return -1
}

// --- HELPER KUESIONER OTOMATIS ---
func hitungSkorKuesioner() int {
	var totalSkor, jawaban int
	pertanyaan := []string{
		"1. Cemas mikirin tugas/Tubes?",
		"2. Pola tidur mengalami penurunan kualitas?",
		"3. Merasa hampa atau kosong setelah deadline selesai?",
		"4. Pusing mikirin error kodingan?",
		"5. Malas berinteraksi (terjadi penurunan kepekaan sosial)?",
	}

	for _, p := range pertanyaan {
		fmt.Println(p)
		for {
			fmt.Print("Jawaban: ")
			fmt.Scan(&jawaban)
			if jawaban >= 1 && jawaban <= 5 {
				totalSkor += jawaban
				break
			}
			fmt.Println("-> Input salah! Masukkan angka 1 sampai 5.")
		}
	}
	// Rumus Konversi ke 0-100
	return ((totalSkor - 5) * 100) / 20
}

// --- FUNGSI CRUD ---
func tambahData(B *TipeB) {
	if B.n >= MaksData { fmt.Println("Kapasitas Penuh!"); return }
	idx := B.n
	
	// Kuesioner wajib ditaruh di AWAL sebelum input ID
	fmt.Println("\n[WAJIB] Silakan isi kuesioner (Skala 1-5) terlebih dahulu:")
	skorOtomatis := hitungSkorKuesioner()
	
	B.C[idx].SelfAssessment[0].Skor = skorOtomatis
	B.C[idx].SkorSelfAssessment = skorOtomatis

	fmt.Printf("\n-> Kuesioner Selesai! Skor kamu: %d\n", skorOtomatis)
	fmt.Print("-> Masukkan ID(12Digit) & Tgl(YYYYMMDD) (Spasi): ")
	fmt.Scan(&B.C[idx].IDPengguna, &B.C[idx].Date)
	
	fmt.Print("-> Kesimpulan/Refleksi Utama (Tanpa spasi): ")
	fmt.Scan(&B.C[idx].SelfAssessment[0].RefleksiDiri)

	B.C[idx].n = 1
	B.n++
	
	fmt.Println("-> Data berhasil tersimpan!")
}

func cetakData(B TipeB) {
	if B.n == 0 { fmt.Println("Data kosong."); return }
	fmt.Printf("| %-14s | %-8s | %-16s | %-4s | %-26s |\n", "ID Pengguna", "Tanggal", "Refleksi", "Skor", "Tingkat Stres")
	for i := 0; i < B.n; i++ {
		status := "Stres berat! Cari bantuan"
		if B.C[i].SkorSelfAssessment <= 40 {
			status = "Aman, stres rendah"
		} else if B.C[i].SkorSelfAssessment <= 70 {
			status = "Stres sedang, butuh rehat"
		}
		fmt.Printf("| %-14d | %-8s | %-16s | %-4d | %-26s |\n",
			B.C[i].IDPengguna, B.C[i].Date, B.C[i].SelfAssessment[0].RefleksiDiri, B.C[i].SkorSelfAssessment, status)
	}
}

func editData(B *TipeB) {
	var id int64
	fmt.Print("Masukkan ID diedit: "); fmt.Scan(&id)
	if idx := cariIndeks(*B, id); idx != -1 {
		
		// Kuesioner wajib ditaruh di AWAL saat edit
		fmt.Println("\n[WAJIB] Silakan isi ulang kuesioner untuk update skor:")
		skorBaru := hitungSkorKuesioner() 
		B.C[idx].SelfAssessment[0].Skor = skorBaru
		B.C[idx].SkorSelfAssessment = skorBaru
		
		fmt.Printf("\n-> Skor Terupdate: %d\n", skorBaru)
		fmt.Print("-> Tgl & Refleksi Baru (Spasi): ")
		fmt.Scan(&B.C[idx].Date, &B.C[idx].SelfAssessment[0].RefleksiDiri)
		
		fmt.Println("-> Update sukses!")
	} else { fmt.Println("Data tidak ada!") }
}

func hapusData(B *TipeB) {
	var id int64
	fmt.Print("Masukkan ID dihapus: "); fmt.Scan(&id)
	if idx := cariIndeks(*B, id); idx != -1 {
		for i := idx; i < B.n-1; i++ { B.C[i] = B.C[i+1] }
		B.n--
		fmt.Println("Hapus sukses!")
	} else { fmt.Println("Data tidak ada!") }
}

// --- FUNGSI SEARCHING ---
func cariSeq(B TipeB) {
	var id int64
	fmt.Print("Cari ID: "); fmt.Scan(&id)
	if idx := cariIndeks(B, id); idx != -1 {
		fmt.Printf("Ketemu! Tgl: %s | Refleksi: %s | Skor: %d\n", B.C[idx].Date, B.C[idx].SelfAssessment[0].RefleksiDiri, B.C[idx].SkorSelfAssessment)
	} else { fmt.Println("Data tidak ada.") }
}

func cariBin(B TipeB) {
	var id int64
	fmt.Print("Cari ID (Pastikan urut ID dulu): "); fmt.Scan(&id)
	kiri, kanan, posisi := 0, B.n-1, -1
	for kiri <= kanan && posisi == -1 {
		tengah := (kiri + kanan) / 2
		if id < B.C[tengah].IDPengguna { kanan = tengah - 1 } else if id > B.C[tengah].IDPengguna { kiri = tengah + 1 } else { posisi = tengah }
	}
	if posisi != -1 {
		fmt.Printf("Ketemu! Tgl: %s | Refleksi: %s | Skor: %d\n", B.C[posisi].Date, B.C[posisi].SelfAssessment[0].RefleksiDiri, B.C[posisi].SkorSelfAssessment)
	} else { fmt.Println("Data tidak ada.") }
}

// --- FUNGSI SORTING ---
func sortSkor(B *TipeB) {
	for i := 0; i < B.n-1; i++ {
		m := i
		for j := i + 1; j < B.n; j++ {
			if B.C[j].SkorSelfAssessment > B.C[m].SkorSelfAssessment { m = j }
		}
		B.C[i], B.C[m] = B.C[m], B.C[i]
	}
	fmt.Println("Selesai diurutkan dari skor terbesar.")
}

func sortTanggal(B *TipeB) {
	for i := 1; i < B.n; i++ {
		key, j := B.C[i], i-1
		for j >= 0 && key.Date < B.C[j].Date {
			B.C[j+1] = B.C[j]
			j--
		}
		B.C[j+1] = key
	}
	fmt.Println("Selesai diurutkan dari tanggal terlama.")
}
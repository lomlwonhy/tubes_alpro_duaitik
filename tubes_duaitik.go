package main

import "fmt"


type Assessment struct {
	ID        int
	UserID    string
	Tanggal   string 
	SkorTotal int
	Kategori  string
}

const MAX_DATA = 100
var daftarAsesmen [MAX_DATA]Assessment
var jumlahData int = 0
var nextID int = 1

func tambahAsesmen(userID string, tanggal string, q1, q2, q3, q4, q5 int) {
	if jumlahData >= MAX_DATA {
		fmt.Println("[Error] Kapasitas penyimpanan riwayat sudah penuh!")
		return
	}
	
	skor := q1 + q2 + q3 + q4 + q5
	kat := "Sehat"
	if skor > 18 {
		kat = "Stres Berat"
	} else if skor > 10 {
		kat = "Stres Ringan"
	}

	fmt.Printf("\n➜ HASIL ASESMEN: Skor Total = %d | Kategori = %s\n", skor, kat)

	daftarAsesmen[jumlahData] = Assessment{
		ID:        nextID,
		UserID:    userID,
		Tanggal:   tanggal, 
		SkorTotal: skor,
		Kategori:  kat,
	}
	nextID++
	jumlahData++
	fmt.Println("Data assessment berhasil ditambahkan ke dalam sistem!")
}

func tampilkanSemua() {
	if jumlahData == 0 {
		fmt.Println("Belum ada riwayat data asesmen.")
		return
	}
	fmt.Println("\n=================== DAFTAR RIWAYAT ASESMEN ===================")
	for i := 0; i < jumlahData; i++ {
		a := daftarAsesmen[i]
		fmt.Printf("ID: %-3d | User: %-6s | Tgl: %-10s | Skor Total: %-2d | Status: %s\n", 
			a.ID, a.UserID, a.Tanggal, a.SkorTotal, a.Kategori)
	}
}

func ubahAsesmen(id int, newUserID string, newTanggal string, q1, q2, q3, q4, q5 int) {
	idx := -1
	for i := 0; i < jumlahData; i++ {
		if daftarAsesmen[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Data ID tidak ditemukan! Gagal mengubah data.")
		return
	}

	skor := q1 + q2 + q3 + q4 + q5
	kat := "Sehat"
	if skor > 18 {
		kat = "Stres Berat"
	} else if skor > 10 {
		kat = "Stres Ringan"
	}

	fmt.Printf("\n➜ HASIL UPDATE ASESMEN: Skor Total Baru = %d | Kategori = %s\n", skor, kat)

	daftarAsesmen[idx].UserID = newUserID
	daftarAsesmen[idx].Tanggal = newTanggal
	daftarAsesmen[idx].SkorTotal = skor
	daftarAsesmen[idx].Kategori = kat

	fmt.Println("Data assessment berhasil diperbarui!")
}

func hapusAsesmen(id int) {
	idx := -1
	for i := 0; i < jumlahData; i++ {
		if daftarAsesmen[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Data ID tidak ditemukan!")
		return
	}
	
	for i := idx; i < jumlahData-1; i++ {
		daftarAsesmen[i] = daftarAsesmen[i+1]
	}
	jumlahData--
	fmt.Println("Data assessment berhasil dihapus dari sistem!")
}

func cariBerdasarkanUser(userID string) {
	ditemukan := false
	fmt.Printf("\n--- Hasil Pencarian Sequential Search untuk User: %s ---\n", userID)
	for i := 0; i < jumlahData; i++ {
		if daftarAsesmen[i].UserID == userID {
			a := daftarAsesmen[i]
			fmt.Printf("[Ketemu] ID: %d | Skor Total: %d | Status Kesehatan: %s\n", a.ID, a.SkorTotal, a.Kategori)
			ditemukan = true
		}
	}
	if !ditemukan { 
		fmt.Println("Data dengan User ID tersebut tidak ditemukan.") 
	}
}

func binarySearchUser(targetUserID string) {
	if jumlahData == 0 {
		fmt.Println("Belum ada data untuk dicari.")
		return
	}

	for i := 1; i < jumlahData; i++ {
		key := daftarAsesmen[i]
		j := i - 1
		for j >= 0 && daftarAsesmen[j].UserID > key.UserID {
			daftarAsesmen[j+1] = daftarAsesmen[j]
			j--
		}
		daftarAsesmen[j+1] = key
	}

	low, high := 0, jumlahData-1
	idxKetemu := -1
	
	for low <= high {
		mid := (low + high) / 2
		if daftarAsesmen[mid].UserID == targetUserID {
			idxKetemu = mid
			break 
		} else if daftarAsesmen[mid].UserID < targetUserID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idxKetemu != -1 {
		a := daftarAsesmen[idxKetemu]
		fmt.Printf("\n[Ketemu via Binary Search] ID: %d | User ID: %s | Tgl: %s | Skor: %d | Status: %s\n", 
			a.ID, a.UserID, a.Tanggal, a.SkorTotal, a.Kategori)
	} else {
		fmt.Println("Data dengan User ID tersebut tidak dapat ditemukan.")
	}
}

func selectionSortByTanggal() {
	for i := 0; i < jumlahData-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if daftarAsesmen[j].Tanggal < daftarAsesmen[minIdx].Tanggal {
				minIdx = j
			}
		}
		daftarAsesmen[i], daftarAsesmen[minIdx] = daftarAsesmen[minIdx], daftarAsesmen[i]
	}
	fmt.Println("Riwayat berhasil diurutkan berdasarkan tanggal secara Ascending (Selection Sort).")
}

func insertionSortBySkor() {
	for i := 1; i < jumlahData; i++ {
		key := daftarAsesmen[i]
		j := i - 1
		for j >= 0 && daftarAsesmen[j].SkorTotal < key.SkorTotal {
			daftarAsesmen[j+1] = daftarAsesmen[j]
			j--
		}
		daftarAsesmen[j+1] = key
	}
	fmt.Println("Riwayat berhasil diurutkan berdasarkan skor tertinggi ke terendah (Insertion Sort).")
}

func tampilkanStatistik() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data riwayat untuk dianalisis.")
		return
	}
	fmt.Println("\n================ LAPORAN STATISTIK KESEHATAN MENTAL ================")
	
	fmt.Println("a. 5 Hasil Terakhir Self-Assessment Pengguna:")
	start := 0
	if jumlahData > 5 { 
		start = jumlahData - 5 
	}
	for i := jumlahData - 1; i >= start; i-- {
		fmt.Printf("   - ID: %d | User: %s | Skor: %d (%s)\n", 
			daftarAsesmen[i].ID, daftarAsesmen[i].UserID, daftarAsesmen[i].SkorTotal, daftarAsesmen[i].Kategori)
	}
	
	totalSkor := 0
	for i := 0; i < jumlahData; i++ { 
		totalSkor += daftarAsesmen[i].SkorTotal 
	}
	rataRata := float64(totalSkor) / float64(jumlahData)
	fmt.Printf("b. Rata-rata skor self-assessment pengguna: %.2f\n", rataRata)
}

func main() {
	var pilihan, targetInt int
	var targetStr, targetTanggal string 

	pertanyaan := [5]string{
		"1. Seberapa sering Anda merasa gugup atau cemas tanpa alasan yang jelas?",
		"2. Seberapa sering Anda merasa begitu gelisah sehingga tidak bisa duduk tenang?",
		"3. Seberapa sering Anda merasa putus asa atau tidak memiliki harapan?",
		"4. Seberapa sering Anda merasa bahwa segala sesuatu membutuhkan usaha yang sangat berat?",
		"5. Seberapa sering Anda merasa sangat sedih sehingga tidak ada yang bisa menghibur Anda?",
	}

	for {
		fmt.Println("\n=============================================")
		fmt.Println("          APLIKASI GO-MENTAL (CLI)           ")
		fmt.Println("=============================================")
		fmt.Println("1. Tampilkan Semua Riwayat Asesmen")
		fmt.Println("2. Tambah Data Asesmen Baru")
		fmt.Println("3. Ubah Data Asesmen Berdasarkan ID")
		fmt.Println("4. Hapus Data Asesmen Berdasarkan ID")
		fmt.Println("5. Cari User ID (Sequential Search)")
		fmt.Println("6. Cari User ID (Binary Search)")
		fmt.Println("7. Urutkan Berdasarkan Tanggal (Selection Sort)")
		fmt.Println("8. Urutkan Berdasarkan Skor Total (Insertion Sort)")
		fmt.Println("9. Tampilkan Laporan Statistik & Rata-rata")
		fmt.Println("10. Keluar Aplikasi")
		fmt.Print("Pilih Opsi Menu (1-10): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanSemua()
		case 2:
			fmt.Print("Masukkan User ID Baru: ")
			fmt.Scan(&targetStr)
			fmt.Print("Masukkan Tanggal (contoh: 2026-05-27): ")
			fmt.Scan(&targetTanggal)
			
			var q [5]int
			fmt.Println("\n--- JAWAB PERTANYAAN (Skala 1: Tidak Pernah, s/d 5: Selalu) ---")
			for i := 0; i < 5; i++ {
				for {
					fmt.Println(pertanyaan[i])
					fmt.Print("Jawaban Anda (1-5): ")
					fmt.Scan(&q[i])
					if q[i] >= 1 && q[i] <= 5 { break }
					fmt.Println("Input tidak valid! Harap masukkan angka antara 1 sampai 5.\n")
				}
				fmt.Println() 
			}
			tambahAsesmen(targetStr, targetTanggal, q[0], q[1], q[2], q[3], q[4])

		case 3:
			fmt.Print("Masukkan ID Data yang ingin Anda ubah: ")
			fmt.Scan(&targetInt)
			
			fmt.Print("Masukkan User ID Pengganti: ")
			fmt.Scan(&targetStr)
			fmt.Print("Masukkan Tanggal Baru (contoh: 2026-05-27): ")
			fmt.Scan(&targetTanggal)
			
			var q [5]int
			fmt.Println("\n--- ISI KEMBALI KUESIONER BARU (Skala 1-5) ---")
			for i := 0; i < 5; i++ {
				for {
					fmt.Println(pertanyaan[i])
					fmt.Print("Jawaban Baru (1-5): ")
					fmt.Scan(&q[i])
					if q[i] >= 1 && q[i] <= 5 { break }
					fmt.Println("Input tidak valid! Masukkan angka antara 1 sampai 5.\n")
				}
				fmt.Println()
			}
			ubahAsesmen(targetInt, targetStr, targetTanggal, q[0], q[1], q[2], q[3], q[4])

		case 4:
			fmt.Print("Masukkan ID Data yang akan dihapus: ")
			fmt.Scan(&targetInt)
			hapusAsesmen(targetInt)
		case 5:
			fmt.Print("Masukkan User ID yang ingin dicari: ")
			fmt.Scan(&targetStr)
			cariBerdasarkanUser(targetStr)
		case 6:
			fmt.Print("Masukkan User ID yang ingin dicari (Binary Search): ")
			fmt.Scan(&targetStr)
			binarySearchUser(targetStr)
		case 7:
			selectionSortByTanggal()
			tampilkanSemua()
		case 8:
			insertionSortBySkor()
			tampilkanSemua()
		case 9:
			tampilkanStatistik()
		case 10:
			fmt.Println("Terima kasih telah menggunakan aplikasi Go-Mental!")
			return
		default:
			fmt.Println("Pilihan menu tidak valid, silakan coba kembali.")
		}
	}
}

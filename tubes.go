package main

import "fmt"

type Milestone struct {
	ID, Deadline, TingkatStres int
	NamaTugas                  string
	StatusSelesai              bool
}

type KamusBot struct {
	KataKunci, Respon string
}

type ProfilUser struct {
	Nama, Fokus string
}

type CatatanMood struct {
	Mood       [30]int
	JumlahHari int
}

// I.S.: Layar terminal dalam keadaan siap.
// F.S.: Teks pilihan menu tercetak di layar.
func TampilkanMenuUtama() {
	fmt.Println("==================================================")
	fmt.Println("     AI CHATBOT MENTAL MILESTONE & WELLBEING")
	fmt.Println("==================================================")
	fmt.Println("1. Atur Profil Pengguna")
	fmt.Println("2. Mulai Sesi Chat (Konsultasi Mental)")
	fmt.Println("3. Manajemen Pencapaian Kerja (Milestone)")
	fmt.Println("4. Daily Mood Check-in (Catat Perasaan)")
	fmt.Println("5. Analisis & Ringkasan Kesejahteraan")
	fmt.Println("0. Keluar Aplikasi")
	fmt.Println("--------------------------------------------------")
}

// I.S.: Layar terminal dalam keadaan siap.
// F.S.: Teks pilihan menu tercetak di layar.
func TampilkanMenuManajemen() {
	fmt.Println("==================================================")
	fmt.Println("            MANAJEMEN PENCAPAIAN KERJA            ")
	fmt.Println("==================================================")
	fmt.Println("1. Tambah Pekerjaan Baru")
	fmt.Println("2. Lihat Daftar Semua Pekerjaan")
	fmt.Println("3. Update Detail Pekerjaan")
	fmt.Println("4. Hapus Pekerjaan")
	fmt.Println("5. Cari Pekerjaan")
	fmt.Println("6. Urutkan Pekerjaan")
	fmt.Println("9. Kembali ke Menu Utama")
	fmt.Println("--------------------------------------------------")
}

// I.S.: Variabel profilAktif belum memiliki data yang terpersonalisasi.
// F.S.: Variabel profilAktif terisi dengan Nama dan Fokus utama pengguna.
func AturProfil(profil *ProfilUser) {
	var pilihan int

	fmt.Println("==================================================")
	fmt.Println("               ATUR PROFIL PENGGUNA               ")
	fmt.Println("==================================================")
	fmt.Print("Masukkan Nama Panggilan (Tanpa Spasi): ")
	fmt.Scan(&profil.Nama)
	fmt.Println("Pilih Fokus Utama Anda saat ini:")
	fmt.Println("1. Akademik & Perkuliahan")
	fmt.Println("2. Organisasi & Kepanitiaan")
	fmt.Println("3. Persiapan Karier / Lomba")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		profil.Fokus = "Akademik & Perkuliahan"
	} else if pilihan == 2 {
		profil.Fokus = "Organisasi & Kepanitiaan"
	} else if pilihan == 3 {
		profil.Fokus = "Karier & Lomba"
	} else {
		profil.Fokus = "Umum"
	}
	fmt.Println("Mantap! Profil berhasil disimpan. Halo,", profil.Nama)
}

// I.S.: Array Mood mungkin kosong atau berisi data sebagian.
// F.S.: Data mood bertambah, jumlahHari bertambah, dan rata-rata tercetak.
func CatatMood(catatan *CatatanMood) {
	var inputMood, i, total int
	var rataRata float64

	if catatan.JumlahHari >= 30 {
		fmt.Println("Catatan bulan ini sudah penuh. Tetap semangat!")
		return
	}

	fmt.Println("==================================================")
	fmt.Println("               DAILY MOOD CHECK-IN                ")
	fmt.Println("==================================================")
	fmt.Println("Skala Mood: 1 (Sangat Buruk) - 5 (Sangat Baik)")
	fmt.Print("Bagaimana perasaanmu hari ini? (1-5): ")
	fmt.Scan(&inputMood)

	if inputMood >= 1 && inputMood <= 5 {
		catatan.Mood[catatan.JumlahHari] = inputMood
		catatan.JumlahHari = catatan.JumlahHari + 1
		fmt.Println("Mood hari ini berhasil dicatat.")

		total = 0
		for i = 0; i < catatan.JumlahHari; i++ {
			total = total + catatan.Mood[i]
		}
		rataRata = float64(total) / float64(catatan.JumlahHari)

		fmt.Print("Rata-rata mood kamu sejauh ini: ")
		fmt.Printf("%.2f\n", rataRata)

		if rataRata < 3.0 {
			fmt.Println("Perhatian: Rata-rata moodmu sedang rendah. Jangan lupa istirahat!")
		}
	} else {
		fmt.Println("Input tidak valid. Masukkan angka 1 sampai 5.")
	}
}

// I.S.: Array arrMilestone berisi sejumlah data (jumlahData).
// F.S.: Layar menampilkan persentase progres dan nama tugas dengan beban tertinggi.
func TampilkanAnalisis(arrMilestone [100]Milestone, jumlahData int) {
	var i, tugasSelesai, tugasBelum, maxStresIdx int
	var persentase float64

	fmt.Println("==================================================")
	fmt.Println("        ANALISIS & RINGKASAN KESEJAHTERAAN        ")
	fmt.Println("==================================================")

	if jumlahData == 0 {
		fmt.Println("Belum ada data pekerjaan untuk dianalisis.")
		return
	}

	tugasSelesai = 0
	tugasBelum = 0
	maxStresIdx = 0

	for i = 0; i < jumlahData; i++ {
		if arrMilestone[i].StatusSelesai {
			tugasSelesai = tugasSelesai + 1
		} else {
			tugasBelum = tugasBelum + 1
		}

		if arrMilestone[i].TingkatStres > arrMilestone[maxStresIdx].TingkatStres {
			maxStresIdx = i
		}
	}

	persentase = (float64(tugasSelesai) / float64(jumlahData)) * 100

	fmt.Println("Total Pekerjaan     :", jumlahData)
	fmt.Println("Pekerjaan Selesai   :", tugasSelesai)
	fmt.Println("Pekerjaan Tertunda  :", tugasBelum)
	fmt.Printf("Progres Penyelesaian: %.2f%%\n", persentase)
	fmt.Println("--------------------------------------------------")
	fmt.Println("Peringatan Beban Mental:")
	fmt.Println("Tugas yang paling membebani pikiran saat ini adalah '", arrMilestone[maxStresIdx].NamaTugas, "' (Skala: ", arrMilestone[maxStresIdx].TingkatStres, ")")
	fmt.Println("Saran: Coba cicil tugas ini lebih awal atau diskusi sama teman.")
	fmt.Println("==================================================")
}

// I.S.: Array memiliki kapasitas yang cukup (belum penuh).
// F.S.: Data baru tefunc TambahMilestrsimpan di indeks ke-jumlahData, dan jumlahData bertambah 1.
one(arrMilestone *[100]Milestone, jumlahData *int) {
	var idBaru, deadlineBaru, stresBaru int
	var namaBaru string

	if *jumlahData >= 100 {
		fmt.Println("Kapasitas penuh, tidak bisa nambah data lagi!")
		return
	}

	fmt.Print("Masukkan ID Pekerjaan (Angka): ")
	fmt.Scan(&idBaru)
	fmt.Print("Masukkan Nama Tugas Pekerjaan Sekarang (Tanpa Spasi): ")
	fmt.Scan(&namaBaru)
	fmt.Print("Masukkan Sisa Hari Deadline (Angka): ")
	fmt.Scan(&deadlineBaru)
	fmt.Print("Masukkan Tingkat Stres (1-10): ")
	fmt.Scan(&stresBaru)

	arrMilestone[*jumlahData].ID = idBaru
	arrMilestone[*jumlahData].NamaTugas = namaBaru
	arrMilestone[*jumlahData].Deadline = deadlineBaru
	arrMilestone[*jumlahData].TingkatStres = stresBaru
	arrMilestone[*jumlahData].StatusSelesai = false

	*jumlahData = *jumlahData + 1
	fmt.Println("Pekerjaan baru berhasil ditambahkan.")
}

// I.S.: Array berisi data pekerjaan.
// F.S.: Seluruh elemen array tercetak di terminal dengan format yang rapi.
func TampilkanSemuaMilestone(arrMilestone [100]Milestone, jumlahData int) {
	var i int

	if jumlahData == 0 {
		fmt.Println("Daftar pekerjaan kamu masih kosong.")
		return
	}

	fmt.Println("==================================================")
	fmt.Println("             DAFTAR PENCAPAIAN KERJA              ")
	fmt.Println("==================================================")

	for i = 0; i < jumlahData; i++ {
		fmt.Print("ID: ", arrMilestone[i].ID, " | Tugas: ", arrMilestone[i].NamaTugas)
		fmt.Print(" | Sisa Hari: ", arrMilestone[i].Deadline, " | Tingkat Stres: ", arrMilestone[i].TingkatStres)

		if arrMilestone[i].StatusSelesai {
			fmt.Println(" | Status: Selesai")
		} else {
			fmt.Println(" | Status: Belum")
		}
	}
	fmt.Println("--------------------------------------------------")
}

// I.S.: Terdapat data dengan ID target di dalam array.
// F.S.: Atribut data (Nama, Deadline, Stres, atau Status) pada ID tersebut berhasil diperbarui.
func UpdateMilestone(arrMilestone *[100]Milestone, jumlahData int) {
	var idTarget, i, pilihanUpdate int
	var ditemukan bool

	if jumlahData == 0 {
		fmt.Println("Data masih kosong, belum ada yang bisa di-update.")
		return
	}

	fmt.Print("Masukkan ID Pekerjaan yang mau diubah: ")
	fmt.Scan(&idTarget)

	ditemukan = false
	i = 0

	for i < jumlahData && !ditemukan {
		if arrMilestone[i].ID == idTarget {
			ditemukan = true
			fmt.Println("Data ditemukan! Pilih bagian yang mau diubah:")
			fmt.Println("1. Nama Pekerjaan")
			fmt.Println("2. Sisa Hari Deadline")
			fmt.Println("3. Tingkat Stres")
			fmt.Println("4. Status Penyelesaian")
			fmt.Print("Masukkan pilihan (1-4): ")
			fmt.Scan(&pilihanUpdate)

			if pilihanUpdate == 1 {
				fmt.Print("Masukkan Nama Baru (Tanpa Spasi): ")
				fmt.Scan(&arrMilestone[i].NamaTugas)
			} else if pilihanUpdate == 2 {
				fmt.Print("Masukkan Sisa Hari Baru (Angka): ")
				fmt.Scan(&arrMilestone[i].Deadline)
			} else if pilihanUpdate == 3 {
				fmt.Print("Masukkan Tingkat Stres Baru (1-10): ")
				fmt.Scan(&arrMilestone[i].TingkatStres)
			} else if pilihanUpdate == 4 {
				if arrMilestone[i].StatusSelesai {
					arrMilestone[i].StatusSelesai = false
					fmt.Println("Status diubah menjadi: Belum Selesai")
				} else {
					arrMilestone[i].StatusSelesai = true
					fmt.Println("Status diubah menjadi: Selesai")
				}
			} else {
				fmt.Println("Pilihan update tidak valid.")
			}

			if pilihanUpdate >= 1 && pilihanUpdate <= 3 {
				fmt.Println("Detail pekerjaan berhasil diperbarui!")
			}
		} else {
			i++
		}
	}

	if !ditemukan {
		fmt.Println("ID Pekerjaan tidak ditemukan.")
	}
}

// I.S.: Terdapat data target yang ingin dihapus.
// F.S.: Data terhapus, elemen array di bawahnya bergeser ke atas, dan jumlahData berkurang 1.
func HapusMilestone(arrMilestone *[100]Milestone, jumlahData *int) {
	var idTarget, i, j int
	var ditemukan bool

	if *jumlahData == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	fmt.Print("Masukkan ID Pekerjaan yang mau dihapus: ")
	fmt.Scan(&idTarget)

	ditemukan = false
	i = 0

	for i < *jumlahData && !ditemukan {
		if arrMilestone[i].ID == idTarget {
			ditemukan = true
		} else {
			i++
		}
	}

	if ditemukan {
		for j = i; j < *jumlahData-1; j++ {
			arrMilestone[j] = arrMilestone[j+1]
		}
		*jumlahData = *jumlahData - 1
		fmt.Println("Pekerjaan berhasil dihapus dari sistem!")
	} else {
		fmt.Println("ID Pekerjaan tidak ditemukan.")
	}
}

// I.S.: Elemen array tidak terurut berdasarkan ID/Deadline.
// F.S.: Elemen array tersusun rapi dari ID/Deadline terkecil hingga terbesar.
func UrutkanByID(arrMilestone *[100]Milestone, jumlahData int) {
	var i, j int
	var temp Milestone

	for i = 1; i < jumlahData; i++ {
		temp = arrMilestone[i]
		j = i - 1
		for j >= 0 && arrMilestone[j].ID > temp.ID {
			arrMilestone[j+1] = arrMilestone[j]
			j = j - 1
		}
		arrMilestone[j+1] = temp
	}
}

// I.S.: Target pencarian (Nama/ID) diinputkan oleh pengguna.
// F.S.: Jika ditemukan, data dicetak. Jika tidak, mencetak pesan error.
func MenuCariPekerjaan(arrMilestone *[100]Milestone, jumlahData int) {
	var pilihan, targetID, indeks, low, high, i, mid int
	var targetNama string
	var ditemukan bool

	if jumlahData == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Berdasarkan Nama Tugas")
	fmt.Println("2. Berdasarkan ID Tugas")
	fmt.Print("Masukkan pilihan (1/2): ")
	fmt.Scan(&pilihan)

	indeks = -1
	ditemukan = false

	if pilihan == 1 {
		fmt.Print("Masukkan Nama Pekerjaan: ")
		fmt.Scan(&targetNama)

		i = 0
		for i < jumlahData && !ditemukan {
			if arrMilestone[i].NamaTugas == targetNama {
				ditemukan = true
				indeks = i
			} else {
				i++
			}
		}
	} else if pilihan == 2 {
		fmt.Print("Masukkan ID Pekerjaan: ")
		fmt.Scan(&targetID)

		UrutkanByID(arrMilestone, jumlahData)

		low = 0
		high = jumlahData - 1
		for low <= high && !ditemukan {
			mid = (low + high) / 2
			if arrMilestone[mid].ID == targetID {
				ditemukan = true
				indeks = mid
			} else if arrMilestone[mid].ID < targetID {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	} else {
		fmt.Println("Pilihan salah.")
		return
	}

	if ditemukan {
		fmt.Println("Pekerjaan Ditemukan!")
		fmt.Print("ID: ", arrMilestone[indeks].ID, " | Tugas: ", arrMilestone[indeks].NamaTugas)
		fmt.Print(" | Sisa Hari: ", arrMilestone[indeks].Deadline, " | Tingkat Stres: ", arrMilestone[indeks].TingkatStres)

		if arrMilestone[indeks].StatusSelesai {
			fmt.Println(" | Status: Selesai")
		} else {
			fmt.Println(" | Status: Belum")
		}
	} else {
		fmt.Println("Pekerjaan tidak ditemukan. Silakan periksa kembali ketikan Anda.")
	}
}

// I.S.: Elemen array tidak terurut berdasarkan ID/Deadline.
// F.S.: Elemen array tersusun rapi dari ID/Deadline terkecil hingga terbesar.
func UrutkanDeadlineTerdekat(arrMilestone *[100]Milestone, jumlahData int) {
	var i, j int
	var temp Milestone

	for i = 1; i < jumlahData; i++ {
		temp = arrMilestone[i]
		j = i - 1

		for j >= 0 && arrMilestone[j].Deadline > temp.Deadline {
			arrMilestone[j+1] = arrMilestone[j]
			j = j - 1
		}
		arrMilestone[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Deadline terdekat!")
}

// I.S.: Elemen array tidak terurut berdasarkan Tingkat Stres.
// F.S.: Elemen array tersusun rapi dari Tingkat Stres paling tinggi di indeks teratas.
func UrutkanStresTertinggi(arrMilestone *[100]Milestone, jumlahData int) {
	var i, j, maxIdx int
	var temp Milestone

	for i = 0; i < jumlahData-1; i++ {
		maxIdx = i
		for j = i + 1; j < jumlahData; j++ {
			if arrMilestone[j].TingkatStres > arrMilestone[maxIdx].TingkatStres {
				maxIdx = j
			}
		}
		temp = arrMilestone[maxIdx]
		arrMilestone[maxIdx] = arrMilestone[i]
		arrMilestone[i] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Tingkat Stres tertinggi!")
}

// I.S.: Array arrMilestone sudah berisi data pekerjaan sebanyak jumlahData (susunan data bisa dalam keadaan acak/tidak terurut).
// F.S.: Elemen di dalam array arrMilestone telah berubah susunannya menjadi terurut sesuai pilihan user, atau mencetak pesan peringatan jika data kosong/input tidak valid.
func MenuSorting(arrMilestone *[100]Milestone, jumlahData int) {
	var pilihan int

	if jumlahData == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	fmt.Println("1. Urutkan Deadline Terdekat")
	fmt.Println("2. Urutkan Tingkat Stres Tertinggi")
	fmt.Print("Pilih metode urut (1/2): ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		UrutkanDeadlineTerdekat(arrMilestone, jumlahData)
	} else if pilihan == 2 {
		UrutkanStresTertinggi(arrMilestone, jumlahData)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// I.S.: Pengguna memilih menu 3 di halaman utama.
// F.S.: Loop berjalan untuk menerima opsi CRUD, dan berhenti jika opsi 'Kembali' dipilih.
func JalankanMenuManajemen(arrMilestone *[100]Milestone, jumlahData *int) {
	var pilihan int
	var menuBerjalan bool

	menuBerjalan = true

	for menuBerjalan {
		TampilkanMenuManajemen()
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahMilestone(arrMilestone, jumlahData)
		case 2:
			TampilkanSemuaMilestone(*arrMilestone, *jumlahData)
		case 3:
			UpdateMilestone(arrMilestone, *jumlahData)
		case 4:
			HapusMilestone(arrMilestone, jumlahData)
		case 5:
			MenuCariPekerjaan(arrMilestone, *jumlahData)
		case 6:
			MenuSorting(arrMilestone, *jumlahData)
		case 9:
			menuBerjalan = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// I.S.: Array daftarKamus masih kosong.
// F.S.: Array daftarKamus terisi penuh dengan pasangan kata kunci dan responsnya.
func InisialisasiKamus(arrKamus *[50]KamusBot, jumlahKamus *int) {
	arrKamus[0].KataKunci = "stres"
	arrKamus[0].Respon = "Tarik napas dalam-dalam. Istirahat sejenak sangat penting untuk mental Anda."

	arrKamus[1].KataKunci = "lelah"
	arrKamus[1].Respon = "Kerja keras Anda berharga, tapi tubuh Anda juga butuh tidur yang cukup."

	arrKamus[2].KataKunci = "pusing"
	arrKamus[2].Respon = "Coba jauhkan mata dari layar sejenak dan minum air putih."

	arrKamus[3].KataKunci = "banyak"
	arrKamus[3].Respon = "Fokus pada satu tugas pada satu waktu. Jangan lihat semuanya sekaligus."

	arrKamus[4].KataKunci = "selesai"
	arrKamus[4].Respon = "Kerja bagus! Rayakan pencapaian kecil ini untuk menjaga semangat."

	arrKamus[5].KataKunci = "malas"
	arrKamus[5].Respon = "Wajar kok. Anggap aja lagi pause sebentar. Coba cicil 5 menit aja, kalau masih berat, baru istirahat."

	arrKamus[6].KataKunci = "bingung"
	arrKamus[6].Respon = "Pecah masalah besarnya jadi bagian-bagian kecil. Selesaikan satu-satu biar otak gak *overload*."

	arrKamus[7].KataKunci = "begadang"
	arrKamus[7].Respon = "Jangan dipaksakan. Jam tidur yang berantakan bikin fokus hancur seharian. Mending tutup layar dan tidur sekarang."

	arrKamus[8].KataKunci = "rapat"
	arrKamus[8].Respon = "Terlalu banyak koordinasi memang menguras energi. Kasih batas waktu kapan kamu harus berhenti mikirin urusan organisasi."

	arrKamus[9].KataKunci = "error"
	arrKamus[9].Respon = "Tinggalkan keyboard sebentar. Cuci muka atau cari udara segar, biasanya solusi muncul pas otak lagi rileks."

	arrKamus[10].KataKunci = "burnout"
	arrKamus[10].Respon = "Ini tanda kamu butuh libur total dari rutinitas. Jangan buka laptop hari ini, lakukan hal yang kamu suka."

	arrKamus[11].KataKunci = "kelompok"
	arrKamus[11].Respon = "Dinamika kerja bareng orang lain memang menantang. Komunikasikan bebanmu, jangan dipikul sendiri semua."

	arrKamus[12].KataKunci = "takut"
	arrKamus[12].Respon = "Masa depan memang gak pasti. Fokus aja sama apa yang bisa kamu kontrol hari ini, selangkah demi selangkah."

	arrKamus[13].KataKunci = "bosan"
	arrKamus[13].Respon = "Cari suasana baru. Pindah tempat nugas, atau coba ganti metode kerjamu biar rutinitasnya gak monoton."

	arrKamus[14].KataKunci = "overthinking"
	arrKamus[14].Respon = "Pikiranmu sedang menjebakmu dengan skenario terburuk. Tuliskan apa yang kamu cemaskan biar kepalamu lebih lega."

	arrKamus[15].KataKunci = "marah"
	arrKamus[15].Respon = "Tarik napas. Jangan ambil keputusan atau ngetik pesan apa pun saat emosi sedang tinggi. Tenangkan diri dulu."

	arrKamus[16].KataKunci = "gagal"
	arrKamus[16].Respon = "Gak apa-apa, ini bagian dari proses belajar. Evaluasi pelan-pelan di mana letak kurangnya tanpa perlu menyalahkan diri sendiri."

	arrKamus[17].KataKunci = "sendiri"
	arrKamus[17].Respon = "Kamu gak sendirian. Hubungi teman atau orang terdekat, sekadar ngobrol ringan bisa bikin perasaan jauh lebih baik."

	arrKamus[18].KataKunci = "tugas"
	arrKamus[18].Respon = "Daftar kerjaan memang kelihatan menumpuk. Kerjakan dari yang *deadline*-nya paling dekat, abaikan yang lain dulu."

	arrKamus[19].KataKunci = "menyerah"
	arrKamus[19].Respon = "Istirahatlah, tapi jangan berhenti. Ingat lagi alasan awal kenapa kamu memulai semua ini."

	arrKamus[20].KataKunci = "degdegan"
	arrKamus[20].Respon = "Wajar merasa gugup sebelum menghadapi sesuatu yang penting. Tarik napas panjang. Fokus pada persiapanmu, sisanya biarkan mengalir."

	arrKamus[21].KataKunci = "gugup"
	arrKamus[21].Respon = "Rasa gugup itu bukti bahwa kamu peduli dengan hasilnya. Jangan biarkan panik menguasai, kamu sudah berlatih untuk ini."

	arrKamus[22].KataKunci = "waktu"
	arrKamus[22].Respon = "Merasa kehabisan waktu atau telat memulai itu hanya ilusi. Setiap orang punya garis waktunya sendiri. Mulai saja porsimu hari ini."

	arrKamus[23].KataKunci = "tertinggal"
	arrKamus[23].Respon = "Melihat pencapaian orang lain memang kadang bikin *insecure*. Matikan media sosial sebentar dan fokus pada progresmu sendiri."

	arrKamus[24].KataKunci = "sosial"
	arrKamus[24].Respon = "Baterai sosialmu sedang habis. Menarik diri sebentar ke kamar untuk main konsol atau sendiri itu hal yang sangat perlu untuk *recharge*."

	arrKamus[25].KataKunci = "teman"
	arrKamus[25].Respon = "Teman yang baik itu penting, tapi jangan sampai hubungan sosial malah jadi sumber stres. Pilih teman yang bisa mendukungmu."

	arrKamus[26].KataKunci = "musuhan"
	arrKamus[26].Respon = "Konflik dengan teman memang bikin pikiran berat. Kasih jeda waktu untuk mendinginkan kepala sebelum mencoba menyelesaikan masalahnya."

	arrKamus[27].KataKunci = "berantem"
	arrKamus[27].Respon = "Emosi yang meledak-ledak tidak akan menyelesaikan debat. Mundur selangkah, evaluasi apakah masalah ini layak menguras energimu."

	arrKamus[28].KataKunci = "kecewa"
	arrKamus[28].Respon = "Tidak semua hal berjalan sesuai ekspektasi. Terima rasa kecewa itu, lalu cari plan B. Selalu ada jalan keluar lain."

	arrKamus[29].KataKunci = "ngoding"
	arrKamus[29].Respon = "Logika yang ruwet tidak akan selesai kalau dipaksa. Tinggalkan layar, kadang solusi dari algoritma muncul saat otak sedang tidak memikirkannya."

	arrKamus[30].KataKunci = "organisasi"
	arrKamus[30].Respon = "Jangan sampai urusan kepanitiaan atau himpunan mengorbankan akademismu. Belajar bilang 'tidak' untuk tugas di luar kapasitasmu."

	arrKamus[31].KataKunci = "insecure"
	arrKamus[31].Respon = "Tidak perlu membandingkan fisik atau kemampuanmu dengan orang lain. Fokus saja merawat diri dan mengasah keahlian yang kamu miliki."

	arrKamus[32].KataKunci = "depresi"
	arrKamus[32].Respon = "Tolong jangan pendam ini sendirian. Bicaralah dengan orang terdekat atau profesional. Kehadiranmu sangat berharga."

	arrKamus[33].KataKunci = "menangis"
	arrKamus[33].Respon = "Menangis bukan tanda lemah, itu cara tubuh membuang stres. Keluarkan saja semuanya sampai dadamu terasa lebih lega."

	arrKamus[34].KataKunci = "ngantuk"
	arrKamus[34].Respon = "Tubuhmu menagih haknya untuk istirahat. Pekerjaan bisa menunggu besok, tapi kesehatanmu tidak. Tutup mata sekarang."

	arrKamus[35].KataKunci = "insomnia"
	arrKamus[35].Respon = "Jika kamu kesulitan tidur, coba matikan semua layar 1 jam sebelum tidur dan lakukan aktivitas yang menenangkan seperti membaca buku atau meditasi ringan."

	arrKamus[36].KataKunci = "cemas"
	arrKamus[36].Respon = "Cemas itu normal, tapi jangan biarkan itu menguasai. Fokus pada apa yang bisa kamu lakukan sekarang, bukan apa yang mungkin terjadi."

	arrKamus[37].KataKunci = "sibuk"
	arrKamus[37].Respon = "Kesibukan memang tidak bisa dihindari, tapi jangan sampai itu membuatmu lupa untuk merawat diri. Sisihkan waktu untuk istirahat dan hal yang kamu suka."

	arrKamus[38].KataKunci = "deadline"
	arrKamus[38].Respon = "Deadline memang bikin deg-degan, tapi jangan sampai itu membuatmu panik. Buat rencana kerja yang jelas dan mulai dari yang paling penting."

	arrKamus[39].KataKunci = "panik"
	arrKamus[39].Respon = "Panik hanya akan mengacaukan logikamu. Berhenti dulu sejenak, minum air, dan urai masalahnya pelan-pelan."

	arrKamus[40].KataKunci = "sidang"
	arrKamus[40].Respon = "Kamu sudah mempersiapkan kodingan dan konsep ini dengan baik. Kuasai materinya dan hadapi dengan percaya diri."

	arrKamus[41].KataKunci = "praktikum"
	arrKamus[41].Respon = "Tugas praktikum memang datang bertubi-tubi. Kerjakan bagian mendasarnya dulu sebelum memikirkan fitur yang rumit."

	arrKamus[42].KataKunci = "revisi"
	arrKamus[42].Respon = "Revisi adalah tanda program dan pemahamanmu berkembang menjadi lebih baik. Selesaikan satu per satu perbaikan tersebut."

	arrKamus[43].KataKunci = "ujian"
	arrKamus[43].Respon = "Jangan terlalu tegang menghadapi ujian. Baca soalnya dengan teliti dan kerjakan dari materi yang paling kamu kuasai."

	arrKamus[44].KataKunci = "nilai"
	arrKamus[44].Respon = "Nilai di atas kertas penting, tetapi proses pemahaman logikamu jauh lebih berharga untuk masa depan kariermu."

	arrKamus[45].KataKunci = "ipk"
	arrKamus[45].Respon = "IPK bukan satu-satunya penentu kesuksesan. Terus asah keahlian praktis dan portofoliomu di samping menjaga nilai akademik."

	arrKamus[46].KataKunci = "minder"
	arrKamus[46].Respon = "Setiap orang punya kecepatan belajar yang berbeda. Fokus saja melampaui dirimu yang kemarin, bukan orang lain."

	arrKamus[47].KataKunci = "ambis"
	arrKamus[47].Respon = "Menjadi ambisius itu bagus, tetapi jangan sampai ambisimu membakar habis kesehatan mental dan fisikmu sendiri."

	arrKamus[48].KataKunci = "istirahat"
	arrKamus[48].Respon = "Istirahat bukan berarti malas. Itu adalah investasi energi agar kamu bisa berpikir lebih tajam di sesi berikutnya."

	arrKamus[49].KataKunci = "semangat"
	arrKamus[49].Respon = "Senang mendengarnya! Pertahankan energi positif ini dan manfaatkan momentumnya untuk menyelesaikan targetmu."

	*jumlahKamus = 50
}

// I.S.: Input teks pengguna diterima.
// F.S.: Bot mencetak respons sesuai kata kunci, atau respons default jika tidak ada kecocokan.
func JalankanSesiChat(arrKamus [50]KamusBot, jumlahKamus int, profil ProfilUser) {
	var inputUser, respon string
	var chatBerjalan, ketemu bool
	var i int

	fmt.Println("==================================================")
	fmt.Println("              SESI KONSULTASI MENTAL              ")
	fmt.Println("==================================================")
	fmt.Println("Halo", profil.Nama, ", fokus kita sekarang adalah:", profil.Fokus)
	fmt.Println("Ketik 1 kata kunci perasaan Anda (misal: stres, lelah, pusing)")
	fmt.Println("Ketik 'keluar' untuk mengakhiri sesi chat.")

	chatBerjalan = true

	for chatBerjalan {
		fmt.Print("\nAnda: ")
		fmt.Scan(&inputUser)

		if inputUser == "keluar" {
			chatBerjalan = false
			fmt.Println("Bot: Terima kasih sudah bercerita. Jaga kesehatan mentalmu,", profil.Nama, "!")
		} else {
			ketemu = false
			respon = "Bot: Saya di sini untuk mendengar. Tetap semangat dan istirahatlah jika perlu."
			i = 0
a
			for i < jumlahKamus && !ketemu {
				if inputUser == arrKamus[i].KataKunci {
					respon = "Bot: " + arrKamus[i].Respon
					ketemu = true
				} else {
					i++
				}
			}
			fmt.Println(respon)
		}
	}
}

func main() {
	var daftarMilestone [100]Milestone
	var daftarKamus [50]KamusBot
	var profilAktif ProfilUser
	var catatanHarian CatatanMood
	var jumlahMilestone, jumlahKamus, pilihan int
	var aplikasiBerjalan bool

	jumlahMilestone = 0
	catatanHarian.JumlahHari = 0
	profilAktif.Nama = "User"
	profilAktif.Fokus = "Belum Diatur"

	InisialisasiKamus(&daftarKamus, &jumlahKamus)

	aplikasiBerjalan = true

	for aplikasiBerjalan {
		TampilkanMenuUtama()
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			AturProfil(&profilAktif)
		} else if pilihan == 2 {
			JalankanSesiChat(daftarKamus, jumlahKamus, profilAktif)
		} else if pilihan == 3 {
			JalankanMenuManajemen(&daftarMilestone, &jumlahMilestone)
		} else if pilihan == 4 {
			CatatMood(&catatanHarian)
		} else if pilihan == 5 {
			TampilkanAnalisis(daftarMilestone, jumlahMilestone)
		} else if pilihan == 0 {
			aplikasiBerjalan = false
			fmt.Println("Aplikasi ditutup. Semoga hari", profilAktif.Nama, "produktif dan tenang.")
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

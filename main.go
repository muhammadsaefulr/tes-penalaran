package main

import (
	"fmt"
	"strings"
)

const artikel = `
Dengan SISKO dapat selesai dalam waktu singkat. Untuk mempermudah bagian administrasi kurikulum sekolah, SISKO menyediakan fasilitas istimewa yang merupakan inti dari sistem kurikulum sekolah yaitu membantu dalam pembuatan penjadwalan mata pelajaran sekolah yang dapat diproses tidak lebih lama dari 10 menit. Administrator hanya akan memasukkan kondisi dari masing-masing guru yang akan mengajar baik itu dalam 1 minggu seorang guru dapat mengajar berapa jam, selain itu dapat juga melakukan pemesanan tempat dan penempatan hari libur masing-masing guru dalam 1 minggu masa mengajar. Setelah semua kondisi dimasukkan, sistem akan memproses semua data tersebut sehingga menghasilkan jadwal yang optimal dan dapat langsung dipakai karena sistem akan mendeteksi sehingga tidak akan ada jadwal yang bertumpukan satu dengan yang lainnya. Setelah permasalahan penjadwalan dapat ditangani dengan baik, hal yang tidak kalah pentingnya adalah memasukkan data siswa. Program SISKO telah menyediakan fasilitas untuk penanganan penilaian siswa yang secara langsung memasukkan nilai ke dalam raport dan siap dicetak. Untuk sistem penilaian siswa, yang dapat melakukan pengisian hanya Guru yang mengajar mata pelajaran. Sistem penilaian telah disesuaikan dengan KBK sehingga masing-masing guru dapat memasukkan deskripsi narasi dari mata pelajaran. Untuk menampilkan data penilaian dapat disesuaikan kembali dengan kebijaksanaan dari masing-masing lembaga pendidikan apakah ingin menampilkan data nilai akhir siswa maupun menampilkan data nilai siswa setiap kali mengadakan test ataupun tugas tertentu. Selain Modul untuk penjadwalan dan Modul Penilaian siswa, SISKO juga memberikan fasilitas untuk bagian administrasi keuangan sekolah dalam hal pembayaran SPP siswa. Bagian administrasi dapat langsung mengecek siapa siswa yang mempunyai tunggakan SPP dan untuk detail histori pembayaran SPP dari masing-masing siswa dapat dicetak seperti mencetak buku tabungan di bank sehingga mempermudah pekerjaan pihak administrasi keuangan. Administrasi keuangan dapat langsung melakukan pengaturan data pembayaran masing-masing siswa sesuai dengan kebutuhan dan dapat diubah sewaktu-waktu apabila ada kenaikan pembayaran SPP. Apabila siswa tersebut akan melakukan pembayaran, petugas dapat langsung memasukkan data. Hal sama juga dapat dilakukan untuk Data pembayaran Sumbangan Sukarela dan Tabungan Karyawisata.
`

func main() {
	fmt.Println("Artikel asli:")
	fmt.Println(artikel)

	var kataLama, kataBaru string
	fmt.Print("\nMasukkan kata yang ingin diganti: ")
	fmt.Scanln(&kataLama)
	fmt.Print("Masukkan kata pengganti: ")
	fmt.Scanln(&kataBaru)

	artikelBaru := strings.ReplaceAll(artikel, kataLama, kataBaru)

	fmt.Printf("\nArtikel setelah penggantian kata '%s' dengan '%s':\n\n%s\n", kataLama, kataBaru, artikelBaru)
}

# Day 2

Created: 2022-10-26 08:25
Tags: #journal 

Hari kedua pembelajaran Aqua-Developer Batch 2.2, kami diberikan pembekalan oleh Mas Gerin mengenai Git, yaitu sistem manajemen perubahan dokumen (*version control*) yang dapat memudahkan proses sinkronisasi dokumen utama dari berbagai versi dokumen. Pertemuan kali ini, mentoring lebih banyak diisi dengan sesi *hands-on* dimana kami diajak untuk lebih membiasakan diri untuk menggunakan Git melalui perintah-perintah yang umum digunakan dalam pengembangan suatu sistem. Selain itu, kami juga diperkenalkan dengan konsep *trunk based development*, yaitu strategi pengembangan dimana semua *developer* melakukan pengembangan langsung pada *main branch* yang juga disebut sebagai *trunk*.

## Version Control
Ketika banyak *developer* yang terlibat dalam pengembangan suatu sistem dimana perubahan *source code* dari sistem dapat terjadi dengan sangat cepat dan masif, konflik antara perubahan-perubahan yang telah dilakukan akan sering ditemukan. *Version control* memungkinkan kita untuk mengembalikan *source code* ke keadaan yang semula, melihat jejak perubahan yang terjadi dan membandingkannya, serta melihat siapa yang berkontribusi terhadap perubahan tersebut. *Version control* sendiri, menurut Atlassian, merupakan praktek untuk merekam dan mengatur perubahan pada *software code*. Dengan sistem ini, kita dapat mengatur perubahan yang terjadi terhadap *source code* di sepanjang waktu. Salah satu *version control system* yang biasa digunakan ialah Git, yang mana akan kita bahas lebih lanjut pada artikel ini.

## Git Command
Dalam bagian ini, aku akan menuliskan penjelasan singkat untuk perintah-perintah Git yang biasanya digunakan dalam pengembangan suatu sistem. 

**Git Help**
Perintah `git help` dapat digunakan untuk mencari laman manual terkait perintah Git yang akan dipakai. Perintah tersebut dapat digunakan sebagai berikut.

`git help -a` untuk melihat semua *git command* yang dapat digunakan.
`git help <command>` untuk membaca manual terkait *git command* tertentu.

**Git Init**
Perintah `git init` dapat digunakan untuk menginisialisasi repository git baru pada direktori lokal kita. Melalui perintah ini, Git membuat direktori tersembunyi *.git* yang mana menyimpan semua objek yang dibutuhkan Git untuk merekam jejak histori perubahan dokumen kita. Berikut perintah-perintah inisialisasi yang biasa digunakan.

`git init` untuk mengubah direktori yang sedang diakses menjadi sebuah *git directory*. 
`git init <directory_name>` untuk mengubah direktori yang dituliskan menjadi sebuah *git directory*.
`git init <options>` untuk menjalankan inisialisasi dengan fungsi tertentu, lebih lanjut dapat gunakan `git help init`.

**Git Status**
Perintah `git status` dapat digunakan untuk mengetahui perubahan yang terjadi terhadap dokumen-dokumen dalam *git repository*. Git akan membandingkan keadaan dokumen pada *index file* dengan keadaan dokumen saat ini dan menampilkan perbedaan antara keduanya. Perubahan yang terjadi dapat dikembalikan ke keadaan awal dengan melalui perintah `git restore`  atau ditambahkan ke proses *staging* melalui perintah `git add`. Selain itu, perintah ini dapat menunjukkan *branch* yang sedang kita tempati dan kesesuaian antara *local repository* dan *remote repository*.

**Git Restore**
Perintah `git restore` dapat digunakan untuk mengembalikan konten dokumen ke *restore point* yang tercatat pada *index file*. Sederhananya, perintah ini ialah untuk meng-*undo* perubahan yang dibuat. Untuk itu, perintah ini dapat dilakukan melalui beberapa contoh berikut.

`git restore <document_name>` untuk mengembalikan perubahan yang terjadi pada *document_name*.
`git restore --staged <document_name>` untuk mengembalikan perubahan yang terjadi pada dokumen yang telah ditambahkan ke tahapan *staging*.

**Git Add**
Perintah `git add` dapat digunakan untuk menambahkan perubahan yang terjadi pada direktori tempat kita bekerja (*working directory*) kedalam tahapan *staging*. Perintah ini akan memperbaharui *index* dengan perubahan konten yang terjadi pada *working directory* dan mempersiapkan dokumen tersebut untuk tahapan selanjutnya yaitu *commit*. Beberapa contoh perintah `git add`.

`git add .` untuk menambahkan perubahan pada seluruh dokumen dalam *repository* ke tahapan *staging*.
`git add <file_name>` untuk menambahkan dokumen tertentu ke tahapan *staging*.

**Git Commit**
Perintah `git commit` digunakan untuk merekam perubahan yang terjadi pada *repository* berdasarkan konten *index* terbaru. Perintah ini biasanya akan membuat titik baru yang menjadi awal mula dari suatu percabangan (*branch*). *Commit* umumnya dilakukan menggunakan perintah berikut.

`git commit -m "commit message"` untuk melakukan *commit* dengan pesan tertentu terkait perubahan yang telah dilakukan.

Pembahasan mengenai konvensi dalam memberikan *commit message* dapat dibaca lebih lanjut pada artikel yang tertera di referensi.

**Git Remote**

Perintah `git remote` dapat digunakan untuk mengatur koneksi ke *remote repository*. Dengan perintah ini, kita dapat melihat dan menghapus *remote repository* yang telah terhubung atau menambahkan *remote repository* baru. Beberapa perintah yang sering digunakan ialah sebagai berikut

`git remote -v` untuk melihat koneksi dengan *remote repository* yang telah dibuat.
`git remote add <remote_name> <remote_url>` untuk menambahkan *remote repository* baru.

**Git Push**
Perintah `git push` dapat digunakan untuk menyesuaikan perubahan pada *remote repository* dengan *local repository*. Melalui perintah ini, kita dapat mengirimkan perubahan yang telah di-*commit* ke *remote branch*, yang mana harus sesuai dengan *local branch* tempat kita bekerja. Sehingga, kita harus selalu sadar pada *branch* mana kita melakukan perubahan ketika akan melakukan *push*. *Push* ini biasanya dilakukan sebagai berikut,

`git push <remote_name> <branch_name>` melakukan *push* ke *branch* tertentu pada *remote repository*.

**Git Fetch**
Perintah `git fetch` dapat digunakan untuk mendapatkan informasi terkait perubahan yang terjadi pada *remote repository*. Hal ini dapat dilakukan dengan perintah berikut.

`git fetch --all` untuk melihat perubahan pada seluruh *remote repository* yang terhubung.
`git fetch <remote_name>` untuk melihat perubahan pada *remote repository* tertentu.

**Git Pull**
Berbeda dengan perintah *fetch*, `git pull` digunakan untuk melihat dan mengunduh perubahan konten dari *remote repository* sehingga *local repository* dapat diperbaharui dengan konten tersebut. Permintaan ini dapat dilakukan melalui perintah berikut.

`git pull` untuk melakukan *pull* terhadap *master branch* pada *remote repository* yang kemudian akan diintegrasikan dengan perubahan yang telah dibuat pada *local repository*.
`git pull <remote_name> <branch_name>` untuk melakukan *pull* terhadap *branch* tertentu pada *remote repository* yang telah ditentukan.

**Git Branch**
Perintah `git branch` dapat dilakukan untuk membentuk percabangan atau *branch* baru dari *master branch*. *Branch* sendiri merepresentasikan suatu pengembangan yang independen sehingga kita dapat memisahkan perubahan yang akan dilakukan dari pengembangan utama sebelumnya nantinya digabungkan kembali.

`git branch <branch_name>` untuk membentuk percabangan baru.

**Git Merge**
Perintah ini dapat digunakan untuk menggabungkan dua *branch* ke dalam satu kesatuan. Sebagai catatan, kita tidak dapat berada dalam *branch* yang akan digabung. Misalkan, suatu fitur baru dikembangkan pada *branch* dengan nama *newFeature* akan diaplikasikan pada *branch* utama dengan nama *main*. Untuk itu, kita dapat melakukan *merge* melalui beberapa tahapan berikut.

`git checkout main` untuk bekerja atau masuk ke *main branch*.
`git merge newFeature` untuk menggabungkan fitur baru dari *newFeature branch* ke *main branch*.

**Git Stash**
Perintah `git stash` dapat digunakan untuk menyimpan perubahan konten secara sementara seperti halnya melakukan *cut* dan *paste*. Perintah-perintah berikut menunjukkan cara untuk menggunakan *stash*.

`git stash` untuk menyimpan perubahan konten selayaknya *cut*.
`git stash list` untuk melihat daftar *stash* yang tersimpan.
`git stash show` untuk melihat perubahan yang terjadi pada suatu *stash entry*.
`git stash pop` untuk mengaplikasikan perubahan dalam *stash* kepada suatu dokumen, dan kemudian menghapus *stash entry* tersebut.
`git stash apply` memiliki fungsi yang sama seperti *pop* namun *stash entry* tidak dihapus.

## Trunk Based Development
Umumnya terdapat dua metode pengembangan yang dilakukan dengan menggunakan *version control* yaitu *feature based development* dan *trunk based development*. Dalam *feature based development*, *developer* akan membuat *branch*, biasanya dari *trunk*, yang khusus digunakan untuk mengembangkan fitur tertentu. Setelah selesai dikembangkan, fitur tersebut akan di-*merge* kembali ke *trunk*. Sedangkan dalam *trunk based development*, pengembangan dilakukan melalui tugas-tugas kecil yang dapat diselesaikan dalam waktu yang relatif cepat.

Pengembangan fitur dibagi ke dalam bagian-bagian kecil yang dapat diselesaikan dengan segera. Melalui metode ini, kita dapat dihindarkan dari proses-proses *merging* yang melelahkan. Mulai dari *code review* yang panjang hingga penyelesaian konflik yang sulit karenanya *trunk based development* dianggap sebagai strategi *version control* yang lebih baik dibandingkan *feature based development*. Sebagai contoh dari penerapan *trunk based development*, Anda dapat membaca artikel [Moving to Trunk-Based Development](https://medium.com/inside-bukalapak/moving-to-trunk-based-development-6efa394c681c) yang mana membahas mengenai proses migrasi ke *trunk based development* di Bukalapak.


---
## References
1. https://www.atlassian.com/git/tutorials/what-is-version-control
2. https://www.freecodecamp.org/news/writing-good-commit-messages-a-practical-guide/
3. https://medium.com/rupesh-tiwari/conventional-git-commit-messages-and-linting-76e1fbb9e14a
4. https://trunkbaseddevelopment.com/
5. https://cloud.google.com/architecture/devops/devops-tech-trunk-based-development
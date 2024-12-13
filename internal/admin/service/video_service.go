package service

import "go-project/internal/admin/repository"

// VideoService adalah struktur yang menyediakan logika bisnis terkait video
// Struktur ini berkomunikasi dengan lapisan repository untuk menangani operasi basis data terkait video.
type VideoService struct {
	Repo *repository.VideoRepository // Repository untuk berinteraksi dengan basis data
}

// Fungsi ini menerima parameter repository.VideoRepository dan mengembalikan instansi VideoService yang baru.
func NewVideoService(repo *repository.VideoRepository) *VideoService {
	return &VideoService{Repo: repo}
}

// Fungsi ini menerima objek video yang akan disimpan dan mengembalikan ID video yang baru dibuat serta error jika terjadi kesalahan.
func (s *VideoService) CreateVideo(video repository.Video) (int, error) {
	return s.Repo.Create(video)
}

// Fungsi ini mengembalikan daftar video yang ada di dalam basis data serta error jika terjadi kesalahan.
func (s *VideoService) GetAllVideos() ([]repository.Video, error) {
	return s.Repo.GetAll()
}

// Fungsi ini mengembalikan objek video yang sesuai dengan ID yang diberikan dan error jika terjadi kesalahan.
func (s *VideoService) GetVideoByID(id int) (*repository.Video, error) {
	return s.Repo.GetByID(id)
}

// Fungsi ini menerima objek video yang berisi data terbaru dan memperbarui entri video yang ada di basis data.
func (s *VideoService) UpdateVideo(video repository.Video) error {
	return s.Repo.Update(video)
}

// Fungsi ini menerima ID video yang ingin dihapus dan menghapusnya dari basis data.
func (s *VideoService) DeleteVideo(id int) error {
	return s.Repo.Delete(id)
}

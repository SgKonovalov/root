package handlers

const (
	ContentType       = "Content-type"
	ApplicationJson   = "application/json"
	MultipartFormData = "multipart/form-data"
	//Заголовок для файла фотографии
	PhotoHeader = "file"
	//Заголовок для обозначения является ли фотография главной
	IsMainPhoto = "Main"
	//URL домашней страницы
	HomeURL = "/"
	//URL для получения списка всех объявлений.
	GetAdvsListURL = "/getAdvsList"
	//URL для получения 1 объявления.
	GetOneAdvURL = "/getAdv/:id"
	//URL для добавления нового объявления.
	AddNewAdvURL = "/add"
	//URL для добавления новой фотографии.
	AddPhotoAtAdvURL = "/uploadPhoto/:id"
)

package repository

const (
	GetAdvsList = `SELECT adv.adv_name, adv.price, a_p.photo
	FROM advertisement as adv
	LEFT JOIN photos AS a_p on adv.adv_id = a_p.adv_id
	WHERE main_photo = true`
	LongGetOnePhoto = `SELECT adv.adv_name, adv.price, adv.description, a_p.photo
	FROM advertisement as adv
	LEFT JOIN photos AS a_p on adv.adv_id = a_p.adv_id
	WHERE adv.adv_id = $1`
	ShortGetOnePhoto = `SELECT adv_name, price FROM advertisement WHERE adv_id = $1`
	AddNewAdv        = `INSERT INTO advertisement (adv_name, description, price, pub_date)
	VALUES ($1, $2, $3, $4)
	RETURNING adv_id`
	PhotoAdd = `INSERT INTO photos (adv_id, photo, main_photo) VALUES ($1, $2, $3)`
)

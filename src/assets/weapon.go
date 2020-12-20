package assets

type WeaponType int

const (
	Pistol WeaponType = iota
	Automatic
	RPG
	Flamethrower
)

type Weapon struct {
	ammo           string
	bulletInterval int     // offset between two bullets in ms
	bulletRange    int     // in pixels
	bulletSpeed    float64 // pixels per ms
	description    string
	image          string
	imageRotated   string
	length         string
	name           string
	price          int
	sound          SoundEffect
	weaponType     WeaponType
	weight         string
}

func (w *Weapon) Ammo() string {
	return w.name
}

func (w *Weapon) BulletInterval() int {
	return w.bulletInterval
}

func (w *Weapon) BulletRange() int {
	return w.bulletRange
}

func (w *Weapon) BulletSpeed() float64 {
	return w.bulletSpeed
}

func (w *Weapon) Description() string {
	return w.description
}

func (w *Weapon) Image() string {
	return w.image
}

func (w *Weapon) ImageRotated() string {
	return w.imageRotated
}

func (w *Weapon) Length() string {
	return w.length
}

func (w *Weapon) Name() string {
	return w.name
}

func (w *Weapon) Price() int {
	return w.price
}

func (w *Weapon) Sound() SoundEffect {
	return w.sound
}

func (w *Weapon) WeaponType() WeaponType {
	return w.weaponType
}

func (w *Weapon) Weight() string {
	return w.weight
}

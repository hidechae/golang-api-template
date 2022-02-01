package identifier

type UserID uint64

func NewUserID(id uint64) UserID {
	return UserID(id)
}

func (id UserID) Int() uint64 {
	return uint64(id)
}

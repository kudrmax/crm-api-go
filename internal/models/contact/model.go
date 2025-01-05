package contact

import "time"

type Contact struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Telegram  string    `json:"telegram"`
	Birthday  string    `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContactUpdateData struct {
	Name     string
	Phone    string
	Telegram string
	Birthday string
}

//class MContact(Base):
//__tablename__ = "contact"
//
//id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
//name = Column(String, unique=True, nullable=False)
//phone = Column(String)
//telegram = Column(String)
//birthday = Column(Date)

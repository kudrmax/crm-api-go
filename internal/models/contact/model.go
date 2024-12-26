package contact

type Contact struct {
	Id       int
	Name     string
	Phone    string
	Telegram string
	Birthday string

	//CreatedAt time.Time
	//UpdatedAt time.Time
	//CreatedBy string
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

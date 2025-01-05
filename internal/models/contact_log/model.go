package contact_log

import "time"

type ContactLog struct {
	Id         int
	ContactId  int // `json:"contact_id"`
	Datetime   time.Time
	LogMessage string

	//CreatedAt  time.Time
	//UpdatedAt  time.Time
	//CreatedBy  string
}

func (l ContactLog) SetTimestampToNow() {
	l.Datetime = time.Now()
}

type ContactLogUpdateData struct {
	ContactId  int
	Datetime   time.Time
	LogMessage string
}

//class MLog(Base):
//__tablename__ = "contact_log"
//
//id = Column(Integer, primary_key=True)
//contact_id = Column(UUID(as_uuid=True), ForeignKey('contact.id', ondelete="CASCADE"), nullable=False)
//datetime = Column(DateTime(timezone=True), default=lambda: datetime.datetime.now(datetime.timezone.utc))
//log = Column(String, nullable=False, default="")
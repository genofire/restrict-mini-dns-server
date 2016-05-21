package model

import (
	"net"
	"time"
)

type DNSEntry struct {
	Id        uint64    `xorm:"autoincr pk not null unique"  json:"id"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	Owner_id  int64     `xorm:"not null 'login_id'"`
	Subdomain string    `xorm:"unique 'subdomain'" json:"subdomain"`
	IPv4      net.IP    `xorm:"'ipv4'" json:"ipv4"`
	IPv6      net.IP    `xorm:"'ipv6'" json:"ipv6"`
}

// TableName define table name for Xorm
func (e *DNSEntry) TableName() string {
	return "dnsentry"
}

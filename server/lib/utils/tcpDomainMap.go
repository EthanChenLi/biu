package utils

import "sync"

/**
 * tcp 域名映射表
 * 维护tcp域名和FID的映射关系，只做更新和新增不做删除
 */

var TcpDomainMap *TcpDomain


type TcpDomain struct {
	 TcpDomainMap map[string]uint16
	 lock sync.Mutex
}


//实例化tcp维护表
func InitTcpDomainMap() {
	TcpDomainMap =  &TcpDomain{
		TcpDomainMap: make(map[string]uint16),  //domain:fid
	}
}


//直接更新域名表
func (this *TcpDomain)WriteTcpDomainMap(domain string, fid uint16){
	this.lock.Lock()
	defer this.lock.Unlock()
	this.TcpDomainMap[domain] = fid
}
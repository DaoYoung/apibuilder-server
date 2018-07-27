// WebSocket连接管理
//
// liamylian
// 2018/03/09

package conn

import (
	"errors"
	"sync"
)

// 连接池
type Pool struct {
	Name       string
	peers     *sync.Map
}

// 新建
func NewPool(name string) *Pool {
	return &Pool{
		Name:       name,
		peers:     new(sync.Map),
	}
}

// 添加连接节点
func (p *Pool) Add(peer *Peer) error {
	_, exist := p.peers.LoadOrStore(peer.Uid, peer)
	if exist {
		return errors.New("already exist")
	}

	return nil
}

// 移除连接节点
func (p *Pool) Remove(peer *Peer) error {
	p.peers.Delete(peer.Uid)
	return nil
}

// 当前节点数量
func (p *Pool) GetPeerCount() int64 {
	var count int64 = 0
	p.Range(func(*Peer) bool {
		count ++
		return true
	})

	return count
}

// 查找节点
func (p *Pool) GetPeers(kv map[string]string, limit int) []*Peer {
	peers := make([]*Peer, 0)
	p.Range(func(peer *Peer) bool {
		for key, val := range kv {
			if !peer.Remote.Match(key, val) {
				return true
			}
		}
		peers = append(peers, peer)
		if limit == 0 || len(peers) < limit {
			return true
		} else {
			return false
		}
	})

	return peers
}

// 遍历节点
func (p *Pool) Range(f func(peer *Peer) bool) {
	p.peers.Range(func(k, v interface{}) bool {
		peer, ok := v.(*Peer)
		if !ok || peer.IsDestroyed() {
			p.Remove(peer)
			return true
		}

		return f(peer)
	})
}
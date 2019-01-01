package mlib

import (
	"errors"
)

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

/*歌曲实体*/
type MusicManager struct {
	musics []MusicEntry
}

/*歌曲库包装体*/
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

/*播放器中歌曲个数*/
func (m *MusicManager) Len() int {
	return len(m.musics)
}

/*根据索引找歌曲*/
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

/*根据歌名找歌曲*/
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

/*新增歌曲*/
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

/*删除歌曲*/
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]
	if index < len(m.musics)-1 { //中间的元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { //删除仅有的一个元素
		m.musics = make([]MusicEntry, 0)
	} else { //删除最后一个元素
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}

func (m *MusicManager) RemoveByName(Name string) *MusicEntry {
	if len(Name) == 0 {
		return nil
	}
	var removedMusic *MusicEntry
	for k, v := range m.musics {
		if v.Name == Name {
			removedMusic = &m.musics[k]
			m.musics = append(m.musics[:k], m.musics[k+1:]...)
		}
	}
	return removedMusic
}

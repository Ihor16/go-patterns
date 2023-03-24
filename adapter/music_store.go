package adapter

type MusicStore struct {
	adapters []MusicStoreAdapter
}

func NewMusicStore(adapters []MusicStoreAdapter) *MusicStore {
	return &MusicStore{adapters: adapters}
}

func (m *MusicStore) GetRecords() []Record {
	var combined []Record
	for _, adapter := range m.adapters {
		combined = append(combined, adapter.GetRecords()...)
	}
	return combined
}

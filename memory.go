package godb

func (db *Storage) Write(bucket string, key string, payload []byte) error {
	mtx := db.memTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	m, ok := db.memStorage[bucket]
	if ok {
		delete(m, key)
	}
	return db.diskWrite(bucket, key, payload)
}

func (db *Storage) Read(bucket string, key string) (bool, []byte, error) {
	mtx := db.memTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	m, ok := db.memStorage[bucket]
	if !ok {
		m = make(map[string][]byte)
	}
	cacheData, ok := m[key]
	if !ok {
		diskData, err := db.DiskRead(bucket, key)
		if err != nil {
			return false, nil, err
		}
		m[key] = diskData
		db.memStorage[bucket] = m
		return false, diskData, nil
	}
	return true, cacheData, nil
}

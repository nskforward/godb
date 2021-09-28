package godb

func (db *Storage) Write(bucket string, key string, payload []byte) error {
	mtx := db.memTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	m, ok := db.payloadStorage[bucket]
	if ok {
		delete(m, key)
	}
	if !db.diskKeyExists(bucket, key) {
		_, ok := db.keysStorage[bucket]
		if ok {
			delete(db.keysStorage, bucket)
		}
	}
	return db.diskWrite(bucket, key, payload)
}

func (db *Storage) Read(bucket string, key string) ([]byte, error) {
	mtx := db.memTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	m, ok := db.payloadStorage[bucket]
	if !ok {
		m = make(map[string][]byte)
	}
	cacheData, ok := m[key]
	if !ok {
		diskData, err := db.diskRead(bucket, key)
		if err != nil {
			return nil, err
		}
		m[key] = diskData
		db.payloadStorage[bucket] = m
		return diskData, nil
	}
	return cacheData, nil
}

func (db *Storage) Remove(bucket, key string) error {
	mtx := db.memTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	_, ok := db.keysStorage[bucket]
	if ok {
		delete(db.keysStorage, bucket)
	}
	m, ok := db.payloadStorage[bucket]
	if ok {
		delete(m, key)
	}
	return db.diskRemove(bucket, key)
}

func (db *Storage) RemoveAll(bucket string) error {
	mtx := db.memTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	_, ok := db.keysStorage[bucket]
	if ok {
		delete(db.keysStorage, bucket)
	}
	delete(db.payloadStorage, bucket)
	return db.diskRemoveAll(bucket)
}

func (db *Storage) Keys(bucket string) ([]string, error) {
	keys, ok := db.keysStorage[bucket]
	if ok {
		return keys, nil
	}
	keys, err := db.diskKeys(bucket)
	if err != nil {
		return nil, err
	}
	db.keysStorage[bucket] = keys
	return keys, nil
}

package main

// contains checks if a key is present in the keys slice
func contains(keys []uint32, key uint32) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

// remove removes a key from the keys slice
func remove(keys []uint32, key uint32) []uint32 {
	for i, k := range keys {
		if k == key {
			// Remove the key from the slice
			keys = append(keys[:i], keys[i+1:]...)
			break
		}
	}
	return keys
}
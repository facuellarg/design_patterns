# Singleton
singleton implementation in go, was added a mutex for parallelism support
```Go
	mutext.Lock()
	defer mutext.Unlock()
	if managerSingleton == nil {
		managerSingleton = &manager{}
	}
	return managerSingleton
```
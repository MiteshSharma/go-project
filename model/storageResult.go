package model

type StorageResult struct {
	Data interface{}
	Err  *AppError
}

func NewStorageResult(data interface{}, err *AppError) *StorageResult {
	storageResult := &StorageResult{
		Data: data,
		Err:  err,
	}
	return storageResult
}

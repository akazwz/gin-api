package request

import "mime/multipart"

type UploadFile struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required"`
}

type MultipartUpload struct {
	UploadId   string                `json:"upload_id" form:"upload_id" binding:"required"`
	Key        string                `json:"key" form:"key" binding:"required"`
	PartNumber int32                 `json:"part_number" form:"part_number" binding:"required"`
	File       *multipart.FileHeader `json:"file" form:"file" binding:"required"`
}

type UploadChunkFile struct {
	ChunkFile  *multipart.FileHeader `json:"chunk_file" form:"chunk_file" binding:"required"`
	ChunkIndex string                `json:"chunk_index" form:"chunk_index" binding:"required"`
	ChunkHash  string                `json:"chunk_hash" form:"chunk_hash" binding:"required"`
	ChunkSum   int                   `json:"chunk_sum" form:"chunk_sum" binding:"required"`
	FileHash   string                `json:"file_hash" form:"file_hash" binding:"required"`
}

type MergeChunkFile struct {
	FileHash string `json:"file_hash" form:"file_hash" binding:"required"`
	ChunkSum int    `json:"chunk_sum" form:"chunk_sum" binding:""`
}

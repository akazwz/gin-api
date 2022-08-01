package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/utils"
	"github.com/gin-gonic/gin"
)

// UploadChunk 上传分块文件
func UploadChunk(c *gin.Context) {
	// 获取 分块上传参数
	chunkFileUp := request.UploadChunkFile{}
	err := c.ShouldBind(&chunkFileUp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	// 参数信息
	chunkIndex := chunkFileUp.ChunkIndex
	chunkHash := chunkFileUp.ChunkHash
	chunkSum := chunkFileUp.ChunkSum
	fileHash := chunkFileUp.FileHash

	// 获取 file header
	fileHeader := chunkFileUp.ChunkFile

	// 文件信息
	contentType := fileHeader.Header.Get("Content-Type")
	filename := fileHeader.Filename
	size := fileHeader.Size

	// 获取 分块文件 hash
	sha256Hash, err := utils.HashFileByAlgo(fileHeader, "sha256")
	if err != nil {
		response.BadRequest(api.CodeCommonSuccess, nil, "获取文件hash失败", c)
		return
	}

	// hash 不同， 文件不完整
	if chunkHash != sha256Hash {
		response.BadRequest(api.CodeCommonFailed, nil, "文件不完整", c)
		return
	}

	//  分块文件 保存的 文件夹路径
	dir := fmt.Sprintf("public/file/%s", chunkFileUp.FileHash)
	// 单个 分块文件 完整的文件路径, 分块文件命名为 index-hash
	dst := fmt.Sprintf("public/file/%s/%s", fileHash, chunkIndex+"-"+chunkHash)

	// 判断文件夹是否存在,不存在创建文件夹
	pathExists := utils.PathExists(dir)
	if !pathExists {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			response.BadRequest(api.CodeCommonFailed, nil, "创建文件夹失败", c)
			return
		}
	}

	// 判断分块文件是否已经存在,已经存在直接返回成功
	exists := utils.PathExists(dst)
	if exists {
		response.Created(api.CodeCommonSuccess, nil, "已经存在", c)
		return
	}

	// 保存文件
	err = c.SaveUploadedFile(chunkFileUp.ChunkFile, dst)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "保存文件失败", c)
		return
	}

	// 分块上传成功返回
	response.Created(api.CodeCommonSuccess, gin.H{
		"content_type": contentType,
		"filename":     filename,
		"size":         size,
		"chunk_index":  chunkFileUp.ChunkIndex,
		"chunk_sum":    chunkSum,
		"hash_sha256":  sha256Hash,
	}, "success", c)
}

// MergeChunk 合并分块文件
func MergeChunk(c *gin.Context) {
	// 获取参数
	var mergeChunk request.MergeChunkFile
	err := c.ShouldBind(&mergeChunk)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "参数错误", c)
		return
	}

	// 参数信息
	fileHash := mergeChunk.FileHash
	chunkSum := mergeChunk.ChunkSum

	//  分块文件 保存的 文件夹路径
	dir := fmt.Sprintf("public/file/%s", fileHash)

	completeFile := fmt.Sprintf("%s/complete", dir)
	// 判断是否存在完整文件
	exists := utils.PathExists(completeFile)
	if exists {
		response.Created(api.CodeStatusSuccess, nil, "已经存在完整文件", c)
		return
	}

	// 读取文件夹下 所有的分块文件
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		response.BadRequest(api.CodeCommonFailed, nil, "读取文件夹失败", c)
		return
	}

	// 判断所有分块是否完整
	if chunkSum != len(files) {
		response.BadRequest(api.CodeCommonFailed, nil, "文件分块不完整", c)
		return
	}
	// 合并文件， 完整文件为 hash/complete
	timeSpend, err := utils.MergeChunkFile(dir)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "合并文件失败", c)
		return
	}
	response.Ok(api.CodeCommonSuccess, gin.H{
		"time_spend": timeSpend,
	}, "合并成功", c)
}

// ChunkState 查询分块状态
func ChunkState(c *gin.Context) {
	hash := c.Query("hash")
	// 文件夹路径
	dir := fmt.Sprintf("public/file/%s", hash)
	// 判断文件夹是否存在
	dirExists := utils.PathExists(dir)
	// 文件夹不存在
	if !dirExists {
		response.BadRequest(api.CodeStatusNoFile, nil, "没有此文件", c)
		return
	}

	// 完整文件路径
	completeFile := fmt.Sprintf("%s/complete", dir)
	// 判断是否存在完整文件
	exists := utils.PathExists(completeFile)

	// 存在完整文件， 上传成功
	if exists {
		response.Created(api.CodeStatusSuccess, nil, "文件上传成功", c)
		return
	}
	// 按照文件名index排序读取文件夹内的文件
	files, _ := ioutil.ReadDir(dir)
	sort.Slice(files, func(i, j int) bool {
		// 获取文件 index
		filename := files[i].Name()
		index := strings.Split(filename, "-")[0]

		indexInt, _ := strconv.Atoi(index)
		nextInt, _ := strconv.Atoi(strings.Split(files[j].Name(), "-")[0])
		return indexInt < nextInt
	})

	var indexes []int
	// 遍历分块文件
	for _, file := range files {
		filename := file.Name()
		index := strings.Split(filename, "-")[0]
		indexInt, _ := strconv.Atoi(index)
		indexes = append(indexes, indexInt)
	}

	if len(indexes) == 0 {
		response.Ok(api.CodeStatusEmpty, nil, "暂无分片文件", c)
		return
	}
	response.Ok(api.CodeStatusSuccess, gin.H{
		"indexed": indexes,
	}, "获取已上传分块index", c)
}

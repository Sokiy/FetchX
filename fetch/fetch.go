package fetch

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"
)
import "gougou/common"

var (
	config common.Config = common.GetConfig()
)

func StartFetch() {

	// 下载的目录路径
	dirPath := filepath.Join(common.GetCurrPath(), config.DownloadDirName, config.DownloadSubDirName)

	if common.IsDirExist(dirPath) {

		if config.StartPage == 0 || config.StartPage == 1 {
			config.StartPage = 1
			// 如果开始页码为 1 则清空之前的下载目录
			common.ClearDir(dirPath)
			common.CreateDir(dirPath)
		}
	} else {
		common.CreateDir(dirPath)
	}

	pageList := getPageList()
	multiFetch(pageList)
}

// 根据总数获取总共的 page 数并组成一个数组
func getPageList() []int {
	total := config.Total - (config.StartPage-1)*config.PerPage
	pageNum := total / config.PerPage
	if total%config.PerPage != 0 {
		pageNum += 1
	}
	pageList := make([]int, pageNum)
	for i := 0; i < pageNum; i++ {
		pageList[i] = config.StartPage + i
	}
	return pageList
}

// // multiFetch 多协程抓取数据
// // :param pageList
func multiFetch(pageList []int) {
	// wg 提供了等待一组 goroutine 结束的方法
	var wg sync.WaitGroup
	syncChan := make(chan int)

	for i := 0; i < config.ProcessNum; i++ {
		// 添加一个等待执行的 goroutine 数量
		wg.Add(1)
		go func() {
			// defer 表示延迟执行，表示在这个 goroutine 结束之前自动执行 wg.Done(), wg.Done() 表示执行完成了
			defer wg.Done()

			for {
				// 从 syncChan channel 中不停的取值, syncChan 中没有值时，会把 goroutine 挂起，不会占用 CPU 资源
				// ok true channel 开启中  false channel 关闭
				value, ok := <-syncChan
				if !ok {
					return // channel已关闭，表示数组为空，退出goroutine
				}
				url := common.GenerateFetchURL(value)
				log.Println("Fetch Url is: ", url)
				content := common.GetReqContent(url)
				if len(content) == 0 {
					// 重试抓取数据
					time.Sleep(10 * time.Second)
					// 把索引值重新放入 channel 中
					syncChan <- value
					fmt.Println("Something wrong, retry fetch data, page is:", value)
				}
				fileName := fmt.Sprintf(config.FilePrefix+"%d", value)
				common.SaveContentToFile(fileName, content)
			}
		}()
	}

	// 将数组中的值发送到channel
	for _, value := range pageList {
		syncChan <- value
	}

	close(syncChan) // 关闭channel，表示没有更多的值

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("fetch data success")
}

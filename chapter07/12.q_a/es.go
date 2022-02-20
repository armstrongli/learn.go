package main

// batch

var sharedCli = &esClientWithBuffer{}

type esClientWithBuffer struct {
	batchBuffer [][]interface{}
	messageCh   chan interface{}
	shortBuffer []interface{}
}

func (cli *esClientWithBuffer) pushBatch() {
	// todo 队列操作
}

func (cli *esClientWithBuffer) prepareBatch() {
	for msg := range cli.messageCh {
		if len(cli.shortBuffer) == batchSize {
			cli.batchBuffer = append(cli.batchBuffer, cli.shortBuffer)
			cli.shortBuffer = []interface{}{}
		}
		cli.shortBuffer = append(cli.shortBuffer, msg)
	}
}

func pushToElasticSearchService(data interface{}) {
	sharedCli.messageCh <- data
}

var batchSize int = 20

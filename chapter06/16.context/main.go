package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// withCancel()
	// withTimeout()
	// withValue()
	withDeadline()
}

func withDeadline() {
	now := time.Now()
	newtime := now.Add(1 * time.Second)
	ctx, _ := context.WithDeadline(context.TODO(), newtime)
	go tv(ctx)
	go mobile(ctx)
	go game(ctx)

	time.Sleep(2 * time.Second)
}

func tv(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电视")
			return
		default:
		}
		fmt.Println("看电视")
		time.Sleep(300 * time.Millisecond)
	}
}

func mobile(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关手机")
			return
		default:
		}
		fmt.Println("玩儿手机")
		time.Sleep(300 * time.Millisecond)
	}
}

func game(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关游戏机")
			return
		default:
		}
		fmt.Println("玩儿游戏机")
		time.Sleep(300 * time.Millisecond)
	}
}

func withValue() {
	ctx := context.WithValue(context.TODO(), "1", "钱包")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("withValue: 1", ctx.Value("1"))
		fmt.Println("withValue: 2", ctx.Value("2"))
		fmt.Println("withValue: 3", ctx.Value("3"))
		fmt.Println("withValue: 4", ctx.Value("4"))
	}(ctx)
	goToPapa(ctx)

	time.Sleep(2 * time.Second)
}

func goToPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, "2", "充电宝")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToPapa: 1", ctx.Value("1"))
		fmt.Println("goToPapa: 2", ctx.Value("2"))
		fmt.Println("goToPapa: 3", ctx.Value("3"))
		fmt.Println("goToPapa: 4", ctx.Value("4"))
	}(ctx)
	goToMama(ctx)
}

func goToMama(ctx context.Context) {
	ctx = context.WithValue(ctx, "3", "小夹克")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToMama: 1", ctx.Value("1"))
		fmt.Println("goToMama: 2", ctx.Value("2"))
		fmt.Println("goToMama: 3", ctx.Value("3"))
		fmt.Println("goToMama: 4", ctx.Value("4"))
	}(ctx)
	goToGrandma(ctx)
}

func goToGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "大苹果")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToGrandma: 1", ctx.Value("1"))
		fmt.Println("goToGrandma: 2", ctx.Value("2"))
		fmt.Println("goToGrandma: 3", ctx.Value("3"))
		fmt.Println("goToGrandma: 4", ctx.Value("4"))
	}(ctx)
	goToParty(ctx)
}

func goToParty(ctx context.Context) {
	fmt.Println("goToParty: 1", ctx.Value("1"))
	fmt.Println("goToParty: 2", ctx.Value("2"))
	fmt.Println("goToParty: 3", ctx.Value("3"))
	fmt.Println("goToParty: 4", ctx.Value("4"))
}

func withTimeout() {
	ctx, _ := context.WithTimeout(context.TODO(), 1*time.Second)
	fmt.Println("开始部署望远镜，发送信号")
	go distributeMainFrame(ctx)
	go distributeMainBody(ctx)
	go distributeCover(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("任务超时没有完成")
	}

	time.Sleep(20 * time.Second) // 等待20秒后收到任务取消的消息。
}

func distributeMainFrame(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainFrame")
		return
	default:
	}
	fmt.Println("部署:distributeMainFrame")
}

func distributeMainBody(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainBody")
		return
	default:
	}
	fmt.Println("部署:distributeMainBody")
}

func distributeCover(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeCover")
		return
	default:
	}
	fmt.Println("部署:distributeCover")
}

func withCancel() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，要买材料")
	go buyFlour(ctx)
	go buyOil(ctx)
	go buyEgg(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("没电了，取消购买所有食材")
	cancel() // 当 调用 cancel 后，所有由此上下文衍生出的context都会cancel

	time.Sleep(1 * time.Second)
}

func buyFlour(ctx context.Context) {
	fmt.Println("去买面")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买面了")
		return
	default:
	}
	fmt.Println("买面")
}

func buyOil(ctx context.Context) {
	fmt.Println("去买油")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买油了")
		return
	default:
	}
	fmt.Println("买油")
}

func buyEgg(ctx1 context.Context) {
	ctx, _ := context.WithCancel(ctx1)
	// defer cancel() // 无论是否调用衍生出来的ctx的cancel，Done返回的channel都会关闭。
	fmt.Println("去买蛋")
	// time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了")
		return
	default:
	}
	fmt.Println("买蛋")
	go buySEgg(ctx)
	go buyBEgg(ctx)
	time.Sleep(1 * time.Second)
}

func buySEgg(ctx context.Context) {
	fmt.Println("去买蛋:buySEgg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了:buySEgg")
		return
	default:
	}
	fmt.Println("买蛋")
}

func buyBEgg(ctx context.Context) {
	fmt.Println("buyBEgg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了:buyBEgg")
		return
	default:
	}
	fmt.Println("买蛋")
}

package main

func main() {

	bc := newBlockChain()

	cli := CLI{bc}

	cli.Run()

}

package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"os"
	"os/exec"
	"bufio"
)

// 定义 selpg_args 结构体，存储用户从命令行输入的各项参数信息
type selpg_args struct {
	start_page int
	end_page int
	input_file string
	destination string
	page_length int
	file_form string
}

var progname string

// show Usage
func Usage() {
	fmt.Printf("%s is a tool to select pages from your chosen file\n\n", progname)
	fmt.Printf("Usage:\n\n")
	fmt.Printf("selpg -s start_page -e end_page [-f (speciy how the page is sperated)| -l lines_per_page_default_72] [-d destination] [filename]\n\n")
	fmt.Printf("If no file specified, %s will read input from stdin, and use control-D to end.\n\n", progname)
}

// 配置 Pflag
func FlagInit(args *selpg_args) {
	flag.IntVarP(&args.start_page,"startPage","s",-1,"Start page number")
	flag.IntVarP(&args.end_page,"endPage","e",-1,"End page number")
	flag.IntVarP(&args.page_length,"pageLength","l",72,"Lines of a page")
	flag.StringVarP(&args.file_form,"type","f","l","File type")
	flag.StringVarP(&args.destination,"destination","d","","Set destination")
	flag.Usage = Usage
	flag.Parse()
}

// 检测 selpg_args 结构体实例中的每个部分
func ProcessArgs(args *selpg_args) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr,"%s: not enough arguments\n\n",progname)
		Usage()
		os.Exit(1)
	}
	if os.Args[1] != "-s" {
		fmt.Fprintf(os.Stderr,"%s: first arg should be -s\n\n",progname)
		Usage()
		os.Exit(1)
	}
	end_index := 2
	if len(os.Args[1]) == 2 {
		end_index = 3
	}
	if os.Args[end_index] != "-e" {
		fmt.Fprintf(os.Stderr,"%s: last arg should be -e\n\n",progname)
		Usage()
		os.Exit(1)
	}
	if args.start_page > args.end_page || args.start_page < 0 || args.end_page < 0 {
		fmt.Fprintln(os.Stderr,"Invalid args")
		Usage()
		os.Exit(1)
	}
}

// 根据用户输入的参数进行操作
func ProcessInput(args *selpg_args) {
	var stdin io.WriteCloser
	var err error
	var cmd *exec.Cmd

	if args.destination != "" {
		cmd = exec.Command("cat","-n")
		stdin,err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

	if flag.NArg() > 0 {
		args.input_file = flag.Arg(0)
		output,err := os.Open(args.input_file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader := bufio.NewReader(output)
		if args.file_form == "l" {
			count := 0
			for {
				line, _, err := reader.ReadLine()
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				if count/args.page_length >= args.start_page {
					if count/args.page_length > args.end_page {
						break
					} else {
						PrintOrWrite(args, string(line), stdin)
					}
				}
				count++
			}
		} else {
			for pageNum := 0; pageNum <= args.end_page; pageNum++ {
				line, err := reader.ReadString('\f')
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				PrintOrWrite(args, string(line), stdin)
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		count := 0
		target := ""
		for scanner.Scan() {
			line := scanner.Text()
			line += "\n"
			if count/args.page_length >= args.start_page {
				if count/args.page_length <= args.end_page {
					target += line
				}
			}
			count ++
		}
		PrintOrWrite(args, string(target), stdin)
	}

	if args.destination != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// 输出
func PrintOrWrite(args *selpg_args, line string, stdin io.WriteCloser) {
	if args.destination != "" {
		stdin.Write([]byte(line + "\n"))
	} else {
		fmt.Println(line)
	}
}

// main 函数
func main() {
	progname = os.Args[0]
	var args selpg_args
	FlagInit(&args)
	ProcessArgs(&args)
	ProcessInput(&args)
}
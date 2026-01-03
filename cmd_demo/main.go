package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

// calculate 执行数学运算
func calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		// 取模运算需要转换为整数
		return float64(int(a) % int(b)), nil
	default:
		return 0, fmt.Errorf("不支持的运算符: %s", operation)
	}
}

func main() {
	app := &cli.App{
		Name:  "calculator",
		Usage: "执行基本数学运算 (+, -, *, /, %)",
		Commands: []*cli.Command{
			{
				Name:      "add",
				Aliases:   []string{"+"},
				Usage:     "执行加法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("需要两个参数")
					}
					a, err := strconv.ParseFloat(c.Args().Get(0), 64)
					if err != nil {
						return fmt.Errorf("无效的第一个数字: %v", err)
					}
					b, err := strconv.ParseFloat(c.Args().Get(1), 64)
					if err != nil {
						return fmt.Errorf("无效的第二个数字: %v", err)
					}
					result, err := calculate(a, b, "+")
					if err != nil {
						return err
					}
					fmt.Printf("%.2f + %.2f = %.2f\n", a, b, result)
					return nil
				},
			},
			{
				Name:      "subtract",
				Aliases:   []string{"-"},
				Usage:     "执行减法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("需要两个参数")
					}
					a, err := strconv.ParseFloat(c.Args().Get(0), 64)
					if err != nil {
						return fmt.Errorf("无效的第一个数字: %v", err)
					}
					b, err := strconv.ParseFloat(c.Args().Get(1), 64)
					if err != nil {
						return fmt.Errorf("无效的第二个数字: %v", err)
					}
					result, err := calculate(a, b, "-")
					if err != nil {
						return err
					}
					fmt.Printf("%.2f - %.2f = %.2f\n", a, b, result)
					return nil
				},
			},
			{
				Name:      "multiply",
				Aliases:   []string{"*"},
				Usage:     "执行乘法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("需要两个参数")
					}
					a, err := strconv.ParseFloat(c.Args().Get(0), 64)
					if err != nil {
						return fmt.Errorf("无效的第一个数字: %v", err)
					}
					b, err := strconv.ParseFloat(c.Args().Get(1), 64)
					if err != nil {
						return fmt.Errorf("无效的第二个数字: %v", err)
					}
					result, err := calculate(a, b, "*")
					if err != nil {
						return err
					}
					fmt.Printf("%.2f * %.2f = %.2f\n", a, b, result)
					return nil
				},
			},
			{
				Name:      "divide",
				Aliases:   []string{"/"},
				Usage:     "执行除法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("需要两个参数")
					}
					a, err := strconv.ParseFloat(c.Args().Get(0), 64)
					if err != nil {
						return fmt.Errorf("无效的第一个数字: %v", err)
					}
					b, err := strconv.ParseFloat(c.Args().Get(1), 64)
					if err != nil {
						return fmt.Errorf("无效的第二个数字: %v", err)
					}
					result, err := calculate(a, b, "/")
					if err != nil {
						return err
					}
					fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
					return nil
				},
			},
			{
				Name:      "modulus",
				Aliases:   []string{"%"},
				Usage:     "执行取模运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("需要两个参数")
					}
					a, err := strconv.ParseFloat(c.Args().Get(0), 64)
					if err != nil {
						return fmt.Errorf("无效的第一个数字: %v", err)
					}
					b, err := strconv.ParseFloat(c.Args().Get(1), 64)
					if err != nil {
						return fmt.Errorf("无效的第二个数字: %v", err)
					}
					result, err := calculate(a, b, "%")
					if err != nil {
						return err
					}
					fmt.Printf("%.2f %% %.2f = %.2f\n", a, b, result)
					return nil
				},
			},
			{
				Name:      "calc",
				Usage:     "通用计算命令",
				ArgsUsage: "<num1> <operator> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() != 3 {
						return fmt.Errorf("需要三个参数: 数字1 运算符 数字2")
					}
					a, err := strconv.ParseFloat(c.Args().Get(0), 64)
					if err != nil {
						return fmt.Errorf("无效的第一个数字: %v", err)
					}
					operator := c.Args().Get(1)
					b, err := strconv.ParseFloat(c.Args().Get(2), 64)
					if err != nil {
						return fmt.Errorf("无效的第二个数字: %v", err)
					}
					result, err := calculate(a, b, operator)
					if err != nil {
						return err
					}
					fmt.Printf("%.2f %s %.2f = %.2f\n", a, operator, b, result)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

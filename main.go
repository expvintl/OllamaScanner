package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"ollamaScan/models"
	"ollamaScan/utils"

	"github.com/schollz/progressbar/v3"
)

func ScanIP(ip string, port int) models.OllamaInfo {
	c := &http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest("GET", "http://"+ip+":"+strconv.Itoa(port)+"/api/tags", nil)
	if err != nil {
		return models.OllamaInfo{}
	}
	r, err := c.Do(req)
	if err != nil {
		return models.OllamaInfo{}
	}
	defer r.Body.Close()
	d, _ := io.ReadAll(r.Body)
	info := models.OllamaInfo{
		Host: ip,
		Port: port,
	}
	err = json.Unmarshal(d, &info)
	if err != nil {
		return models.OllamaInfo{}
	}
	return info
}
func main() {
	ipList := flag.String("l", "./ips.txt", "IP List File (default=./ips.txt)")
	threads := flag.Int("t", 50, "Thread Num (default=50)")
	outPath := flag.String("o", "./out.txt", "Output Save File (default=./out.txt)")
	useJson := flag.Bool("json", false, "Output JSON Format (default=false)")
	d, err := utils.ReadFile(*ipList)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	list := strings.Split(d, "\n")
	pro := progressbar.Default(int64(len(list)), "Scanning...")
	pool := utils.PoolInfo{}
	pool.NewPool(*threads)
	results := make([]models.OllamaInfo, 2000)
	for i := 0; i < len(list); i++ {
		pool.AddTask(func() {
			info := ScanIP(list[i], 11434)
			if info.Port != 0 {
				results = append(results, info)
			}
			pro.Add(1)
		})
	}
	defer pool.Pool.Release()
	pool.TaskWaitGroup.Wait()
	if *useJson {
		buff := strings.Builder{}
		out, _ := json.Marshal(results)
		buff.Write(out)
		if err := utils.WriteFile(*outPath, buff.String()); err != nil {
			fmt.Println("Save Error:", err)
		}
	} else {
		buff := strings.Builder{}
		for _, i := range results {
			if i.Port == 0 && len(i.Models) == 0 { //skip invalid data
				continue
			}
			buff.WriteString("http://" + i.Host + ":" + strconv.Itoa(i.Port) + "/\n")
			for _, j := range i.Models {
				buff.WriteString("\t" + j.Name + " ModelSize:" + utils.FormatBytes(j.Size) + " " + " ParamSize:" + j.Details.ParameterSize + " QuantLevel:" + j.Details.QuantizationLevel + "\n")
			}
		}
		if err := utils.WriteFile(*outPath, buff.String()); err != nil {
			fmt.Println("Save Error:", err)
		}
	}
	fmt.Println("Done! Saved to", outPath)
}

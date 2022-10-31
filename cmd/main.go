package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"

	// "fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Pm2Data []struct {
	Pid    int    `json:"pid"`
	Name   string `json:"name"`
	Pm2Env struct {
		ExitCode    int    `json:"exit_code"`
		NodeVersion string `json:"node_version"`
		Versioning  struct {
			Type                 string      `json:"type"`
			URL                  string      `json:"url"`
			Revision             string      `json:"revision"`
			Comment              string      `json:"comment"`
			Unstaged             bool        `json:"unstaged"`
			Branch               string      `json:"branch"`
			Remotes              []string    `json:"remotes"`
			Remote               string      `json:"remote"`
			BranchExistsOnRemote bool        `json:"branch_exists_on_remote"`
			Ahead                bool        `json:"ahead"`
			NextRev              interface{} `json:"next_rev"`
			PrevRev              string      `json:"prev_rev"`
			UpdateTime           time.Time   `json:"update_time"`
			RepoPath             string      `json:"repo_path"`
		} `json:"versioning"`
		Version          string `json:"version"`
		UnstableRestarts int    `json:"unstable_restarts"`
		RestartTime      int    `json:"restart_time"`
		PmID             int    `json:"pm_id"`
		CreatedAt        int64  `json:"created_at"`
		AxmDynamic       struct {
		} `json:"axm_dynamic"`
		AxmOptions struct {
			Error                       bool `json:"error"`
			Heapdump                    bool `json:"heapdump"`
			FeatureProfilerHeapsnapshot bool `json:"feature.profiler.heapsnapshot"`
			FeatureProfilerHeapsampling bool `json:"feature.profiler.heapsampling"`
			FeatureProfilerCPUJs        bool `json:"feature.profiler.cpu_js"`
			Latency                     bool `json:"latency"`
			CatchExceptions             bool `json:"catchExceptions"`
			Profiling                   bool `json:"profiling"`
			Metrics                     struct {
				HTTP      bool `json:"http"`
				Runtime   bool `json:"runtime"`
				EventLoop bool `json:"eventLoop"`
				Network   bool `json:"network"`
				V8        bool `json:"v8"`
			} `json:"metrics"`
			Standalone bool `json:"standalone"`
			Tracing    struct {
				Outbound bool `json:"outbound"`
				Enabled  bool `json:"enabled"`
			} `json:"tracing"`
			ModuleConf struct {
			} `json:"module_conf"`
			Apm struct {
				Version string `json:"version"`
				Type    string `json:"type"`
			} `json:"apm"`
			ModuleName    string `json:"module_name"`
			ModuleVersion string `json:"module_version"`
		} `json:"axm_options"`
		AxmMonitor struct {
			UsedHeapSize struct {
				Value    string `json:"value"`
				Type     string `json:"type"`
				Unit     string `json:"unit"`
				Historic bool   `json:"historic"`
			} `json:"Used Heap Size"`
			HeapUsage struct {
				Value    float64 `json:"value"`
				Type     string  `json:"type"`
				Unit     string  `json:"unit"`
				Historic bool    `json:"historic"`
			} `json:"Heap Usage"`
			HeapSize struct {
				Value    string `json:"value"`
				Type     string `json:"type"`
				Unit     string `json:"unit"`
				Historic bool   `json:"historic"`
			} `json:"Heap Size"`
		} `json:"axm_monitor"`
		AxmActions []struct {
			ActionName string `json:"action_name"`
			ActionType string `json:"action_type"`
			Arity      int    `json:"arity"`
		} `json:"axm_actions"`
		PmUptime           int64         `json:"pm_uptime"`
		Status             string        `json:"status"`
		UniqueID           string        `json:"unique_id"`
		Pm2Home            string        `json:"PM2_HOME"`
		Pm2Usage           string        `json:"PM2_USAGE"`
		NAMING_FAILED      string        `json:"_"`
		CondaShlvl         string        `json:"CONDA_SHLVL"`
		CondaPythonExe     string        `json:"CONDA_PYTHON_EXE"`
		CeConda            string        `json:"_CE_CONDA"`
		CeM                string        `json:"_CE_M"`
		CondaExe           string        `json:"CONDA_EXE"`
		Oldpwd             string        `json:"OLDPWD"`
		Shlvl              string        `json:"SHLVL"`
		CfUserTextEncoding string        `json:"__CF_USER_TEXT_ENCODING"`
		ItermSessionID     string        `json:"ITERM_SESSION_ID"`
		Logname            string        `json:"LOGNAME"`
		XpcServiceName     string        `json:"XPC_SERVICE_NAME"`
		User               string        `json:"USER"`
		Tmpdir             string        `json:"TMPDIR"`
		Home               string        `json:"HOME"`
		Term               string        `json:"TERM"`
		CommandMode        string        `json:"COMMAND_MODE"`
		Colorterm          string        `json:"COLORTERM"`
		LcTerminal         string        `json:"LC_TERMINAL"`
		Path               string        `json:"PATH"`
		TermProgram        string        `json:"TERM_PROGRAM"`
		TermProgramVersion string        `json:"TERM_PROGRAM_VERSION"`
		CFBundleIdentifier string        `json:"__CFBundleIdentifier"`
		Shell              string        `json:"SHELL"`
		Pwd                string        `json:"PWD"`
		Lang               string        `json:"LANG"`
		XpcFlags           string        `json:"XPC_FLAGS"`
		ItermProfile       string        `json:"ITERM_PROFILE"`
		Colorfgbg          string        `json:"COLORFGBG"`
		LcTerminalVersion  string        `json:"LC_TERMINAL_VERSION"`
		SSHAuthSock        string        `json:"SSH_AUTH_SOCK"`
		TermSessionID      string        `json:"TERM_SESSION_ID"`
		NodeAppInstance    int           `json:"NODE_APP_INSTANCE"`
		VizionRunning      bool          `json:"vizion_running"`
		KmLink             bool          `json:"km_link"`
		PmPidPath          string        `json:"pm_pid_path"`
		PmErrLogPath       string        `json:"pm_err_log_path"`
		PmOutLogPath       string        `json:"pm_out_log_path"`
		Instances          int           `json:"instances"`
		ExecMode           string        `json:"exec_mode"`
		ExecInterpreter    string        `json:"exec_interpreter"`
		PmCwd              string        `json:"pm_cwd"`
		SourceMapSupport   bool          `json:"source_map_support"`
		PmExecPath         string        `json:"pm_exec_path"`
		NodeArgs           []interface{} `json:"node_args"`
		Name               string        `json:"name"`
		FilterEnv          []interface{} `json:"filter_env"`
		Namespace          string        `json:"namespace"`
		Env                struct {
			UniqueID string `json:"unique_id"`
			Socket   struct {
			} `json:"socket"`
			Pm2Home            string `json:"PM2_HOME"`
			Pm2Usage           string `json:"PM2_USAGE"`
			NAMING_FAILED      string `json:"_"`
			CondaShlvl         string `json:"CONDA_SHLVL"`
			CondaPythonExe     string `json:"CONDA_PYTHON_EXE"`
			CeConda            string `json:"_CE_CONDA"`
			CeM                string `json:"_CE_M"`
			CondaExe           string `json:"CONDA_EXE"`
			Oldpwd             string `json:"OLDPWD"`
			Shlvl              string `json:"SHLVL"`
			CfUserTextEncoding string `json:"__CF_USER_TEXT_ENCODING"`
			ItermSessionID     string `json:"ITERM_SESSION_ID"`
			Logname            string `json:"LOGNAME"`
			XpcServiceName     string `json:"XPC_SERVICE_NAME"`
			User               string `json:"USER"`
			Tmpdir             string `json:"TMPDIR"`
			Home               string `json:"HOME"`
			Term               string `json:"TERM"`
			CommandMode        string `json:"COMMAND_MODE"`
			Colorterm          string `json:"COLORTERM"`
			LcTerminal         string `json:"LC_TERMINAL"`
			Path               string `json:"PATH"`
			TermProgram        string `json:"TERM_PROGRAM"`
			TermProgramVersion string `json:"TERM_PROGRAM_VERSION"`
			CFBundleIdentifier string `json:"__CFBundleIdentifier"`
			Shell              string `json:"SHELL"`
			Pwd                string `json:"PWD"`
			Lang               string `json:"LANG"`
			XpcFlags           string `json:"XPC_FLAGS"`
			ItermProfile       string `json:"ITERM_PROFILE"`
			Colorfgbg          string `json:"COLORFGBG"`
			LcTerminalVersion  string `json:"LC_TERMINAL_VERSION"`
			SSHAuthSock        string `json:"SSH_AUTH_SOCK"`
			TermSessionID      string `json:"TERM_SESSION_ID"`
		} `json:"env"`
		MergeLogs     bool   `json:"merge_logs"`
		Vizion        bool   `json:"vizion"`
		Autorestart   bool   `json:"autorestart"`
		Watch         bool   `json:"watch"`
		InstanceVar   string `json:"instance_var"`
		Pmx           bool   `json:"pmx"`
		Automation    bool   `json:"automation"`
		Treekill      bool   `json:"treekill"`
		Username      string `json:"username"`
		WindowsHide   bool   `json:"windowsHide"`
		KillRetryTime int    `json:"kill_retry_time"`
	} `json:"pm2_env"`
	PmID  int `json:"pm_id"`
	Monit struct {
		Memory int `json:"memory"`
		CPU    int `json:"cpu"`
	} `json:"monit"`
}

func main() {
	//Receiving variables from the user
	set_user_process_names := flag.String("names", "socket", "Name of processes that must be online")
	set_user_bot_api_key := flag.String("tg", "123", "Key telegram bot API")
	set_user_bot_admin_chat_id := flag.String("chat", "123", "Admin Chat ID in Telegram")
	flag.Parse()

	//Conversion of chat_id to int
	admin_chat_id, err := strconv.Atoi(*set_user_bot_admin_chat_id)
	if err != nil {
		panic(err)
	}

	//Creating a slice from the names of user processes
	user_process_names := strings.Split(*set_user_process_names, ",")

	//Getting PM2 processes on the server
	cmd := exec.Command("pm2", "jlist")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	pm2_data := out.String()

	var pm2_data_json Pm2Data
	json.Unmarshal([]byte(pm2_data), &pm2_data_json)

	processes_online := []string{}

	for _, v := range pm2_data_json {
		if v.Pm2Env.Status == "online" {
			processes_online = append(processes_online, v.Name)
		}
	}

	//Comparison of user processes and current processes on the server
	for i := 0; i < len(user_process_names); i++ {
		result := Contains(processes_online, user_process_names[i])

		if !result {
			message := "PM2 процесс с именем " + user_process_names[i] + " не найден."
			SendMessageTelegram(message, *set_user_bot_api_key, admin_chat_id)
		}
	}

}

//Finding a value in a slice
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

//Send message to Telegram
func SendMessageTelegram(message string, api_key string, chat_id int) (bool, error) {
	bot, err := tgbotapi.NewBotAPI(api_key)
	if err != nil {
		log.Panic(err)
	}

	msg := tgbotapi.NewMessage(int64(chat_id), message)
	_, err = bot.Send(msg)

	if err != nil {
		return false, errors.New("message sending error")
	}

	return true, nil
}

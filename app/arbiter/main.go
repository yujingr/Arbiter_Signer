// Copyright (c) 2025 The bel2 developers

package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/arbiter"
	"github.com/BeL2Labs/Arbiter_Signer/app/arbiter/config"
	"golang.org/x/term"
)

type ConfigFile struct {
	Chain struct {
		Esc string `yaml:"esc"`
	} `yaml:"chain"`
	Arbiter struct {
		Listener                         bool   `yaml:"listener"`
		Signer                           bool   `yaml:"signer"`
		Network                          string `yaml:"network"`
		EscStartHeight                   uint64 `yaml:"escStartHeight"`
		EscArbiterContractAddress        string `yaml:"escArbiterContractAddress"`
		EscArbiterManagerContractAddress string `yaml:"escArbiterManagerContractAddress"`
		DataPath                         string `yaml:"dataPath"`
		KeyFilePath                      string `yaml:"keyFilePath"`
		EscArbiterAddress                string `yaml:"escArbiterAddress"`
		EscPrivateKey                    string `yaml:"escPrivateKey"`
		BtcPrivateKey                    string `yaml:"btcPrivateKey"`
	} `yaml:"arbiter"`
}

func getDefaultConfig() ConfigFile {
	var cfg ConfigFile
	cfg.Chain.Esc = "https://api.elastos.io/esc"
	cfg.Arbiter.Listener = true
	cfg.Arbiter.Signer = true
	cfg.Arbiter.Network = "mainnet"
	cfg.Arbiter.EscStartHeight = 28437808
	cfg.Arbiter.EscArbiterContractAddress = "0xA10b92006743Ef3B12077da67e465963743b03D3"
	cfg.Arbiter.EscArbiterManagerContractAddress = "0x9963b5214434776D043A4e98Bc7f33321F6aaCfc"
	
	// Get executable directory for portable paths
	execPath, err := os.Executable()
	if err != nil {
		execPath = "."
	}
	execDir := filepath.Dir(execPath)

	cfg.Arbiter.DataPath = filepath.Join(execDir, "app", "arbiter", "data")
	cfg.Arbiter.KeyFilePath = filepath.Join(execDir, "app", "arbiter", "data", "keys")
	cfg.Arbiter.EscArbiterAddress = ""
	cfg.Arbiter.EscPrivateKey = ""
	cfg.Arbiter.BtcPrivateKey = ""
	
	return cfg
}

func setupConfig() error {
	// Get executable directory
	execPath, err := os.Executable()
	if err != nil {
		execPath = "."
	}
	execDir := filepath.Dir(execPath)
	
	// Use absolute path for config
	configPath := filepath.Join(execDir, "config.yaml")
	
	fmt.Printf("Looking for config file at: %s\n", configPath)
	
	// Check if config already exists and warn user
	if _, err := os.Stat(configPath); err == nil {
		fmt.Println("Warning: Config file already exists and will be overwritten.")
		fmt.Println("Press Enter to continue or Ctrl+C to abort...")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	cfg := getDefaultConfig()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nPress Enter to keep the default value, or type a new value:")

	// Chain configuration
	fmt.Printf("ESC Chain URL [%s]: ", cfg.Chain.Esc)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Chain.Esc = strings.TrimSpace(input)
	}

	// Arbiter configuration
	fmt.Printf("Listener enabled [%v]: ", cfg.Arbiter.Listener)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.Listener = strings.ToLower(strings.TrimSpace(input)) == "true"
	}

	fmt.Printf("Signer enabled [%v]: ", cfg.Arbiter.Signer)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.Signer = strings.ToLower(strings.TrimSpace(input)) == "true"
	}

	fmt.Printf("Network [%s]: ", cfg.Arbiter.Network)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.Network = strings.TrimSpace(input)
	}

	fmt.Printf("ESC Start Height [%d]: ", cfg.Arbiter.EscStartHeight)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		if height, err := strconv.ParseUint(strings.TrimSpace(input), 10, 64); err == nil {
			cfg.Arbiter.EscStartHeight = height
		}
	}

	fmt.Printf("ESC Arbiter Contract Address [%s]: ", cfg.Arbiter.EscArbiterContractAddress)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.EscArbiterContractAddress = strings.TrimSpace(input)
	}

	fmt.Printf("ESC Arbiter Manager Contract Address [%s]: ", cfg.Arbiter.EscArbiterManagerContractAddress)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.EscArbiterManagerContractAddress = strings.TrimSpace(input)
	}

	fmt.Printf("Data Path [%s]: ", cfg.Arbiter.DataPath)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.DataPath = strings.TrimSpace(input)
	}

	fmt.Printf("Key File Path [%s]: ", cfg.Arbiter.KeyFilePath)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.KeyFilePath = strings.TrimSpace(input)
	}

	fmt.Printf("ESC Arbiter Address [%s]: ", cfg.Arbiter.EscArbiterAddress)
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		cfg.Arbiter.EscArbiterAddress = strings.TrimSpace(input)
	}

	// Create directories
	os.MkdirAll(cfg.Arbiter.DataPath, 0755)
	os.MkdirAll(cfg.Arbiter.KeyFilePath, 0755)

	// Write config file
	yamlData, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	err = os.WriteFile(configPath, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	// Always recreate key files
	return createKeyFiles(cfg)
}

func createKeyFiles(cfg ConfigFile) error {
	// Create ESC key file
	escKeyFile := filepath.Join(cfg.Arbiter.KeyFilePath, "escKey.json")
	if _, err := os.Stat(escKeyFile); err == nil {
		fmt.Println("\nWarning: ESC key file already exists and will be overwritten.")
	}

	// Loop until valid ESC key is provided
	var escKey string
	for {
		fmt.Print("\nEnter ESC private key (64 hex characters): ")
		// Read password without echo
		keyBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return fmt.Errorf("error reading private key: %v", err)
		}
		fmt.Println() // Add newline after hidden input
		
		escKey = strings.TrimSpace(string(keyBytes))
		
		// Validate hex format and length
		if len(escKey) != 64 {
			fmt.Println("Error: Private key must be exactly 64 hex characters")
			continue
		}
		if _, err := hex.DecodeString(escKey); err != nil {
			fmt.Println("Error: Private key must be in hex format")
			continue
		}
		break
	}
	
	// Correct key file format
	keyData := fmt.Sprintf(`{"privKey":"%s"}`, escKey)
	if err := os.WriteFile(escKeyFile, []byte(keyData), 0600); err != nil {
		return fmt.Errorf("failed to create ESC key file: %v", err)
	}

	// Create BTC key file
	btcKeyFile := filepath.Join(cfg.Arbiter.KeyFilePath, "btcKey.json")
	if _, err := os.Stat(btcKeyFile); err == nil {
		fmt.Println("\nWarning: BTC key file already exists and will be overwritten.")
	}

	// Loop until valid BTC key is provided
	var btcKey string
	for {
		fmt.Print("\nEnter BTC private key (64 hex characters): ")
		// Read password without echo
		keyBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return fmt.Errorf("error reading private key: %v", err)
		}
		fmt.Println() // Add newline after hidden input
		
		btcKey = strings.TrimSpace(string(keyBytes))
		
		// Validate hex format and length
		if len(btcKey) != 64 {
			fmt.Println("Error: Private key must be exactly 64 hex characters")
			continue
		}
		if _, err := hex.DecodeString(btcKey); err != nil {
			fmt.Println("Error: Private key must be in hex format")
			continue
		}
		break
	}
	
	// Correct key file format
	keyData = fmt.Sprintf(`{"privKey":"%s"}`, btcKey)
	if err := os.WriteFile(btcKeyFile, []byte(keyData), 0600); err != nil {
		return fmt.Errorf("failed to create BTC key file: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) > 1 {
		operation := os.Args[1]
		switch strings.ToLower(operation) {
		case "getpk", "gen":
			myPriKey := os.Args[2]
			pk, err := arbiter.GetPubKey(myPriKey)
			if err != nil {
				fmt.Println("need hex encoded format private key, err:", err)
				return
			}
			fmt.Println("publicKey:", pk)
			return
		}
	}

	// Run setup first
	if err := setupConfig(); err != nil {
		fmt.Printf("Setup failed: %v\n", err)
		os.Exit(1)
	}

	ctx := gctx.New()
	var wg sync.WaitGroup
	wg.Add(1)
	
	// Initialize config file path for gf framework
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath(".")
	
	// start arbiter
	g.Log().Info(ctx, "Starting arbiter...")
	arb := arbiter.NewArbiter(ctx, initConfig(ctx))
	arb.Start()
	wg.Wait()
}

func initConfig(ctx context.Context) *config.Config {
	network, err := g.Cfg().Get(ctx, "arbiter.network")
	if err != nil {
		g.Log().Error(ctx, "get network config err:", err)
		os.Exit(1)
	}
	signer, err := g.Cfg().Get(ctx, "arbiter.signer")
	if err != nil {
		g.Log().Error(ctx, "get signer err:", err)
		os.Exit(1)
	}
	listener, err := g.Cfg().Get(ctx, "arbiter.listener")
	if err != nil {
		g.Log().Error(ctx, "get listener config err:", err)
		os.Exit(1)
	}
	http, err := g.Cfg().Get(ctx, "chain.esc")
	if err != nil {
		g.Log().Error(ctx, "get http config err:", err)
		os.Exit(1)
	}
	escStartHeight, err := g.Cfg().Get(ctx, "arbiter.escStartHeight")
	if err != nil {
		g.Log().Error(ctx, "get escStartHeight config err:", err)
		os.Exit(1)
	}
	escArbiterContractAddress, err := g.Cfg().Get(ctx, "arbiter.escArbiterContractAddress")
	if err != nil {
		g.Log().Error(ctx, "get escArbiterAddress config err:", err)
		os.Exit(1)
	}
	escArbiterManagerAddress, err := g.Cfg().Get(ctx, "arbiter.escArbiterManagerContractAddress")
	if err != nil {
		g.Log().Error(ctx, "get escArbiterManagerAddress config err:", err)
		os.Exit(1)
	}
	escArbiterAddress, err := g.Cfg().Get(ctx, "arbiter.escArbiterAddress")
	if err != nil {
		g.Log().Error(ctx, "get escArbiterAddress config err:", err)
		os.Exit(1)
	}
	gDataPath, err := g.Cfg().Get(ctx, "arbiter.dataPath")
	if err != nil {
		g.Log().Error(ctx, "get dataPath config err:", err)
		os.Exit(1)
	}
	gKeyFilePath, err := g.Cfg().Get(ctx, "arbiter.keyFilePath")
	if err != nil {
		g.Log().Error(ctx, "get keyFilePath config err:", err)
		os.Exit(1)
	}
	dataPath := getExpandedPath(gDataPath.String())
	keyFilePath := getExpandedPath(gKeyFilePath.String())

	g.Log().Info(ctx, "btcCreator:", signer)
	g.Log().Info(ctx, "listener:", listener)
	g.Log().Info(ctx, "http:", http)
	g.Log().Info(ctx, "escStartHeight:", escStartHeight)
	g.Log().Info(ctx, "escArbiterContractAddress:", escArbiterContractAddress)
	g.Log().Info(ctx, "escArbiterManagerAddress:", escArbiterManagerAddress)
	g.Log().Info(ctx, "escArbiterAddress:", escArbiterAddress)
	g.Log().Info(ctx, "dataPath:", dataPath)
	g.Log().Info(ctx, "keyFilePath:", keyFilePath)

	// if want to submit to ESC contract successfully, need to use esc ela as gas.
	escKeyFilePath := gfile.Join(keyFilePath, "escKey.json")
	arbiterKeyFilePath := gfile.Join(keyFilePath, "btcKey.json")
	logPath := gfile.Join(dataPath, "logs/")
	loanPath := gfile.Join(dataPath, "loan/")
	loanNeedSignReqPath := gfile.Join(loanPath, "request/")
	loanNeedSignFailedPath := gfile.Join(loanPath, "failed/")
	loanNeedSignSignedPath := gfile.Join(loanPath, "signed/")
	LoanSignedEventPath := gfile.Join(dataPath, "loan_signed_event/")

	return &config.Config{
		Network:                          network.String(),
		Signer:                           signer.Bool(),
		Listener:                         listener.Bool(),
		Http:                             http.String(),
		ESCStartHeight:                   escStartHeight.Uint64(),
		ESCArbiterContractAddress:        escArbiterContractAddress.String(),
		ESCArbiterManagerContractAddress: escArbiterManagerAddress.String(),
		ESCArbiterAddress:                escArbiterAddress.String(),

		DataDir:            dataPath,
		EscKeyFilePath:     escKeyFilePath,
		ArbiterKeyFilePath: arbiterKeyFilePath,

		LoanNeedSignReqPath:    loanNeedSignReqPath,
		LoanNeedSignFailedPath: loanNeedSignFailedPath,
		LoanNeedSignSignedPath: loanNeedSignSignedPath,
		LoanSignedEventPath:    LoanSignedEventPath,
		LoanLogPath:            logPath,
	}
}

func getExpandedPath(path string) string {
	if len(path) > 0 && path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return path
		}
		path = filepath.Join(homeDir, path[2:])
	}
	return path
}

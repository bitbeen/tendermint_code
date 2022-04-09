package main

import (
	"context"
	abciclient "github.com/tendermint/tendermint/abci/client"
	tconfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	tos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	NodeTypeFull      = "full"
	NodeTypeValidator = "validator"
	NodeTypeSeed      = "seed"
)

func main() {
	logger, err := log.NewDefaultLogger(log.LogFormatPlain, log.LogLevelInfo, true)
	cfg, err := InitConfig("./cfg", "validator", "test")
	creator := abciclient.NewLocalCreator(App{})
	gene, _ := types.GenesisDocFromFile(cfg.GenesisFile())
	app, err := node.New(cfg, logger, creator, gene)
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		panic(err)
	}

	defer func() {
		app.Stop()
		app.Wait()
	}()

	Wait()
}

func Wait() {
	var stopChan = make(chan os.Signal, 0)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan
}

func InitConfig(rootDir string, nodeType string, chainId string) (*tconfig.Config, error) {

	var pv *privval.FilePV
	var err error
	config := tconfig.DefaultConfig()
	config.SetRoot(rootDir)
	config.Mode = nodeType
	tconfig.EnsureRoot(config.RootDir)
	if strings.EqualFold(nodeType, NodeTypeValidator) {
		//创建验证节点的密钥配置文件

		privvalKeyFile := config.PrivValidator.KeyFile()
		privvalStateFile := config.PrivValidator.StateFile()
		if tos.FileExists(privvalKeyFile) {
			pv, err = privval.LoadFilePV(privvalKeyFile, privvalStateFile)
			if err != nil {
				return nil, err
			}

		} else {
			pv, err = privval.GenFilePV(privvalKeyFile, privvalStateFile, "ed25519")
			if err != nil {
				return nil, err
			}
			pv.Save()
		}

		nodeKeyFile := config.NodeKeyFile()

		if tos.FileExists(nodeKeyFile) {

		} else {
			_, err := types.LoadOrGenNodeKey(nodeType)
			if err != nil {
				return nil, err
			}
		}

	}

	//genesis文件
	genesisFile := config.GenesisFile()
	if tos.FileExists(genesisFile) {

	} else {
		genDoc := types.GenesisDoc{
			GenesisTime:     time.Now(),
			ChainID:         chainId,
			InitialHeight:   0,
			ConsensusParams: types.DefaultConsensusParams(),
		}

		ctx, cancel := context.WithTimeout(context.TODO(), 10)
		defer cancel()
		pubKey, err := pv.GetPubKey(ctx)
		if err != nil {
			return nil, err
		}

		validator := types.GenesisValidator{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   10,
		}
		genDoc.Validators = []types.GenesisValidator{validator}
		if err := genDoc.SaveAs(genesisFile); err != nil {
			return nil, err
		}
	}

	if err := tconfig.WriteConfigFile(config.RootDir, config); err != nil {
		return nil, err
	}

	return config, nil
}

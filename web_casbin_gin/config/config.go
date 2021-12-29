package config

import (
	"log"
	"os"
	"path"
)

var (
	BaseDir    string
	CfgDir     string
	ModelPath  string
	PolicyPath string
	err        error
)

func init() {
	if BaseDir, err = os.Getwd(); err != nil {
		log.Fatalf("get pwd failed, err: %s\n", err.Error())
	}
	CfgDir = path.Join(BaseDir, "config")
	ModelPath = path.Join(CfgDir, "model.conf")
	PolicyPath = path.Join(CfgDir, "policy.csv")
	log.Printf("path -> base: %s cfg: %s model: %s policy: %s\n", BaseDir, CfgDir, ModelPath, PolicyPath)
}

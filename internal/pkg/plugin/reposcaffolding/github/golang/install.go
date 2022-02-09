package golang

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"

	"github.com/merico-dev/stream/internal/pkg/log"
	"github.com/merico-dev/stream/pkg/util/github"
	"github.com/merico-dev/stream/pkg/util/zip"
)

// Install installs github-repo-scaffolding-golang with provided options.
func Install(options *map[string]interface{}) (bool, error) {
	var param Param
	if err := mapstructure.Decode(*options, &param); err != nil {
		return false, err
	}

	if errs := validate(&param); len(errs) != 0 {
		for _, e := range errs {
			log.Errorf("Param error: %s", e)
		}
		return false, fmt.Errorf("params are illegal")
	}

	return install(&param)
}

func install(param *Param) (bool, error) {
	// Clear workpath before return
	defer func() {
		if err := os.RemoveAll(DefaultWorkPath); err != nil {
			log.Errorf("Failed to clear workpath: %s", err)
		}
	}()

	if err := download(); err != nil {
		return false, err
	}

	if err := zip.UnZip(filepath.Join(DefaultWorkPath, github.DefaultLatestCodeZipfileName), DefaultWorkPath); err != nil {
		return false, err
	}

	if err := push(param); err != nil {
		return false, err
	}

	return true, nil
}

func download() error {
	ghOption := &github.Option{
		Owner:    DefaultTemplateOwner,
		Repo:     DefaultTemplateRepo,
		NeedAuth: false,
		WorkPath: DefaultWorkPath,
	}
	ghClient, err := github.NewClient(ghOption)
	if err != nil {
		return err
	}

	if err = ghClient.DownloadLatestCodeAsZipFile(); err != nil {
		return err
	}

	return nil
}

func push(param *Param) error {
	ghOption := &github.Option{
		Owner:    param.Owner,
		Repo:     param.Repo,
		NeedAuth: true,
	}
	ghClient, err := github.NewClient(ghOption)
	if err != nil {
		return err
	}

	err = InitRepoLocalAndPushToRemote(filepath.Join(DefaultWorkPath, DefaultTemplateRepo+"-main"), param, ghClient)
	if err != nil {
		return err
	}

	return nil
}
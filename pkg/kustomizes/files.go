package kustomizes

import (
	"io/ioutil"
	"path/filepath"

	"github.com/jenkins-x/jx/v2/pkg/util"
	"github.com/pkg/errors"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// LazyCreate lazily creates the kustomization configuration
func LazyCreate(k *types.Kustomization) *types.Kustomization {
	if k == nil {
		k = &types.Kustomization{}
	}
	k.FixKustomizationPostUnmarshalling()
	return k
}

// LoadKustomization loads the kustomization yaml file from the given directory
func LoadKustomization(dir string) (*types.Kustomization, error) {
	fileName := filepath.Join(dir, "kustomization.yaml")
	exists, err := util.FileExists(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to check if file exists %s", fileName)
	}

	answer := &types.Kustomization{}
	answer.FixKustomizationPostUnmarshalling()

	if !exists {
		return answer, nil
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load file %s", fileName)
	}
	err = yaml.Unmarshal(data, answer)
	if err != nil {
		return nil, errors.Wrapf(err, "failed parse YAML file %s", fileName)
	}
	return answer, nil
}

// SaveKustomization saves the kustomisation file in the given directory
func SaveKustomization(kustomization *types.Kustomization, dir string) error {
	data, err := yaml.Marshal(kustomization)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal Kustomization")
	}
	fileName := filepath.Join(dir, "kustomization.yaml")
	err = ioutil.WriteFile(fileName, data, util.DefaultFileWritePermissions)
	if err != nil {
		return errors.Wrapf(err, "failed write file %s", fileName)
	}
	return nil
}

package secretmapping

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/jenkins-x/jx-gitops/pkg/apis/gitops/v1alpha1"
	"github.com/jenkins-x/jx/v2/pkg/util"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

// LoadSecretMapping loads the secret mapping from the given directory
func LoadSecretMapping(dir string, failIfMissing bool) (*v1alpha1.SecretMapping, string, error) {
	absolute, err := filepath.Abs(dir)
	if err != nil {
		return nil, "", errors.Wrap(err, "creating absolute path")
	}
	relPath := filepath.Join(".jx", "gitops", "secret-mappings.yaml")

	for absolute != "" && absolute != "." && absolute != "/" {
		fileName := filepath.Join(absolute, relPath)
		absolute = filepath.Dir(absolute)

		exists, err := util.FileExists(fileName)
		if err != nil {
			return nil, "", err
		}

		if !exists {
			continue
		}

		config, err := LoadSecretMappingFile(fileName)
		return config, fileName, err
	}
	if failIfMissing {
		return nil, "", errors.Errorf("%s file not found", relPath)
	}
	return nil, "", nil
}

// LoadSecretMappingFile loads a specific secret mapping YAML file
func LoadSecretMappingFile(fileName string) (*v1alpha1.SecretMapping, error) {
	config := &v1alpha1.SecretMapping{}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to load file %s due to %s", fileName, err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML file %s due to %s", fileName, err)
	}

	return config, nil
}

package cmd

import (
	"github.com/jenkins-x/jx-gitops/pkg/cmd/annotate"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/extsecret"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/helm"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/ingress"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/jx_apps"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/kpt"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/kustomize"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/label"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/namespace"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/repository"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/split"
	"github.com/jenkins-x/jx-gitops/pkg/cmd/version"
	"github.com/jenkins-x/jx-gitops/pkg/common"
	"github.com/jenkins-x/jx/v2/pkg/log"
	"github.com/spf13/cobra"
)

// Main creates the new command
func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   common.TopLevelCommand,
		Short: "GitOps utility commands",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	cmd.AddCommand(helm.NewCmdHelm())
	cmd.AddCommand(jx_apps.NewCmdJxApps())
	cmd.AddCommand(kpt.NewCmdKpt())
	cmd.AddCommand(common.SplitCommand(annotate.NewCmdUpdateAnnotate()))
	cmd.AddCommand(common.SplitCommand(extsecret.NewCmdExtSecrets()))
	cmd.AddCommand(common.SplitCommand(ingress.NewCmdUpdateIngress()))
	cmd.AddCommand(common.SplitCommand(kustomize.NewCmdKustomize()))
	cmd.AddCommand(common.SplitCommand(label.NewCmdUpdateLabel()))
	cmd.AddCommand(common.SplitCommand(namespace.NewCmdUpdateNamespace()))
	cmd.AddCommand(common.SplitCommand(repository.NewCmdUpdateRepository()))
	cmd.AddCommand(common.SplitCommand(split.NewCmdSplit()))
	cmd.AddCommand(common.SplitCommand(version.NewCmdVersion()))
	return cmd
}

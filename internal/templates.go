package internal

var RunScriptTemplate = `#!/usr/bin/env bash

echo "Executing DCI Job {{ .JobName }}"
echo

set -x

# Please be sure to place your kubeconfig in the jobs folder.
export KUBECONFIG={{ .KubeconfigPath }}
# Ansible inside dci-openshift-app-agent-ctl uses it for colorful output.
export ANSIBLE_FORCE_COLOR=True

CONFIG={{ .ConfigDir }}/settings.yml

if [ ! -f $KUBECONFIG ]; then
	set +x
	echo "Error: kubeconfig not found in directory."
	exit 1
fi

dci-openshift-app-agent-ctl --config "${CONFIG}" -s
`

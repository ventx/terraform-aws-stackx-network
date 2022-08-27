package tests

import (
	"testing"
)

func TestAwsNetwork(t *testing.T) {
	runAwsNetworkTest(t)
	runAwsNetworkAllTest(t)
	runAwsNetworkNonK8sTest(t)
}

package dyna_plan

import "testing"

func TestHouseRobber(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	maxRobber := houseRobberDpWithMinSpace(nums)
	t.Logf("max_robber: %d", maxRobber)
}

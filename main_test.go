package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGreetingOutput(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			"default",
			nil,
			[]string{
				"Hello, world, from fabrik-test-delta",
				"e2e-cross-repo-spawn",
				"e2e-cross-repo-spawn-31",
				"e2e-cross-repo-spawn-52",
				"e2e-cross-repo-spawn-3472",
				"e2e-cross-repo-spawn-3496",
				"e2e-cross-repo-spawn-3620",
				"e2e-cross-repo-spawn-3737",
				"e2e-cross-repo-spawn-3834",
				"e2e-cross-repo-spawn-3855",
				"e2e-cross-repo-spawn-3922",
				"e2e-cross-repo-spawn-20260801-125714-6717",
				"e2e-cross-repo-spawn-20260801-144217-6757",
				"e2e-cross-repo-spawn-20260803-043620-7283",
				"e2e-cross-repo-spawn-20260804-114237-6848",
				"e2e-cross-repo-spawn-20260810-015521-8654",
				"e2e-cross-repo-spawn-20260810-042952-3676",
				"e2e-cross-repo-spawn-20260813-045054-8185",
				"e2e-cross-repo-spawn-20260813-132237-0922",
				"e2e-cross-repo-spawn-20260813-200250-3163",
				"e2e-cross-repo-spawn-20260814-033648-2864",
				"e2e-cross-repo-spawn-20260814-052928-1443",
				"e2e-cross-repo-spawn-20260816-183605-6156",
				"e2e-cross-repo-spawn-20260816-202551-4426",
				"e2e-cross-repo-spawn-20260828-152756-3489",
				"e2e-cross-repo-spawn-20260828-203211-3677",
				"e2e-cross-repo-spawn-20260829-025800-6960",
				"e2e-cross-repo-spawn-20260904-233317-1389",
				"e2e-cross-repo-spawn-20260905-033525-5078",
				"e2e-cross-repo-spawn-20260905-045125-7449",
				"e2e-cross-repo-spawn-20260906-032338-3125",
				"e2e-cross-repo-spawn-20260906-051635-6981",
				"e2e-cross-repo-spawn-20260907-041222-5625",
				"e2e-cross-repo-spawn-20260907-061600-4362",
				"e2e-cross-repo-spawn-20260925-211421-4577",
				"e2e-cross-repo-spawn-20260925-211231-8640",
				"e2e-cross-repo-spawn-20260927-051241-6031",
				"e2e-cross-repo-spawn-20260927-100522-3438",
				"e2e-cross-repo-spawn-20260927-130424-9126",
				"e2e-cross-repo-spawn-20260927-154712-4779",
				"e2e-cross-repo-spawn-20260927-174515-5813",
				"e2e-cross-repo-spawn-20260927-200053-0854",
				"e2e-cross-repo-spawn-20260927-224553-5238",
				"e2e-cross-repo-spawn-20260928-053845-2570",
				"e2e-cross-repo-spawn-20260929-160333-5841",
				"e2e-cross-repo-spawn-20260930-040636-8317",
				"e2e-cross-repo-spawn-20260930-165742-0672",
				"e2e-cross-repo-spawn-20260930-201519-9599",
				"e2e-cross-repo-spawn-20260930-223204-9105",
				"e2e-cross-repo-spawn-20261001-011358-9982",
				"e2e-cross-repo-spawn-20261013-044742-9594",
				"e2e-cross-repo-spawn-20261013-182234-6893",
				"e2e-cross-repo-spawn-20261014-020802-7520",
				"e2e-cross-repo-spawn-20261014-063103-4181",
				"e2e-cross-repo-spawn-20261014-073218-1422",
				"e2e-cross-repo-spawn-20261014-080419-9005",
				"e2e-cross-repo-spawn-20261014-193357-3941",
				"e2e-cross-repo-spawn-20261014-225257-2433",
				"e2e-cross-repo-spawn-20261014-232701-1219",
				"e2e-cross-repo-spawn-20261015-065000-2855",
				"e2e-cross-repo-spawn-20261015-072002-0653",
				"e2e-cross-repo-spawn-20261017-141052-4438",
				"e2e-cross-repo-spawn-20261017-165734-4675",
				"e2e-cross-repo-spawn-20261019-055333-0025",
				"e2e-cross-repo-spawn-20261019-062236-1759",
				"e2e-cross-repo-spawn-20261019-070912-5993",
				"e2e-cross-repo-spawn-20261019-122156-0514",
				"e2e-cross-repo-spawn-20261019-130515-9466",
				"e2e-cross-repo-spawn-20261019-134215-7557",
				"e2e-cross-repo-spawn-20261019-154343-1133",
			},
		},
		{
			"with argument",
			[]string{"Alice"},
			[]string{
				"Hello, Alice, from fabrik-test-delta",
				"e2e-cross-repo-spawn",
				"e2e-cross-repo-spawn-31",
				"e2e-cross-repo-spawn-52",
				"e2e-cross-repo-spawn-3472",
				"e2e-cross-repo-spawn-3496",
				"e2e-cross-repo-spawn-3620",
				"e2e-cross-repo-spawn-3737",
				"e2e-cross-repo-spawn-3834",
				"e2e-cross-repo-spawn-3855",
				"e2e-cross-repo-spawn-3922",
				"e2e-cross-repo-spawn-20260801-125714-6717",
				"e2e-cross-repo-spawn-20260801-144217-6757",
				"e2e-cross-repo-spawn-20260803-043620-7283",
				"e2e-cross-repo-spawn-20260804-114237-6848",
				"e2e-cross-repo-spawn-20260810-015521-8654",
				"e2e-cross-repo-spawn-20260810-042952-3676",
				"e2e-cross-repo-spawn-20260813-045054-8185",
				"e2e-cross-repo-spawn-20260813-132237-0922",
				"e2e-cross-repo-spawn-20260813-200250-3163",
				"e2e-cross-repo-spawn-20260814-033648-2864",
				"e2e-cross-repo-spawn-20260814-052928-1443",
				"e2e-cross-repo-spawn-20260816-183605-6156",
				"e2e-cross-repo-spawn-20260816-202551-4426",
				"e2e-cross-repo-spawn-20260828-152756-3489",
				"e2e-cross-repo-spawn-20260828-203211-3677",
				"e2e-cross-repo-spawn-20260829-025800-6960",
				"e2e-cross-repo-spawn-20260904-233317-1389",
				"e2e-cross-repo-spawn-20260905-033525-5078",
				"e2e-cross-repo-spawn-20260905-045125-7449",
				"e2e-cross-repo-spawn-20260906-032338-3125",
				"e2e-cross-repo-spawn-20260906-051635-6981",
				"e2e-cross-repo-spawn-20260907-041222-5625",
				"e2e-cross-repo-spawn-20260907-061600-4362",
				"e2e-cross-repo-spawn-20260925-211421-4577",
				"e2e-cross-repo-spawn-20260925-211231-8640",
				"e2e-cross-repo-spawn-20260927-051241-6031",
				"e2e-cross-repo-spawn-20260927-100522-3438",
				"e2e-cross-repo-spawn-20260927-130424-9126",
				"e2e-cross-repo-spawn-20260927-154712-4779",
				"e2e-cross-repo-spawn-20260927-174515-5813",
				"e2e-cross-repo-spawn-20260927-200053-0854",
				"e2e-cross-repo-spawn-20260927-224553-5238",
				"e2e-cross-repo-spawn-20260928-053845-2570",
				"e2e-cross-repo-spawn-20260929-160333-5841",
				"e2e-cross-repo-spawn-20260930-040636-8317",
				"e2e-cross-repo-spawn-20260930-165742-0672",
				"e2e-cross-repo-spawn-20260930-201519-9599",
				"e2e-cross-repo-spawn-20260930-223204-9105",
				"e2e-cross-repo-spawn-20261001-011358-9982",
				"e2e-cross-repo-spawn-20261013-044742-9594",
				"e2e-cross-repo-spawn-20261013-182234-6893",
				"e2e-cross-repo-spawn-20261014-020802-7520",
				"e2e-cross-repo-spawn-20261014-063103-4181",
				"e2e-cross-repo-spawn-20261014-073218-1422",
				"e2e-cross-repo-spawn-20261014-080419-9005",
				"e2e-cross-repo-spawn-20261014-193357-3941",
				"e2e-cross-repo-spawn-20261014-225257-2433",
				"e2e-cross-repo-spawn-20261014-232701-1219",
				"e2e-cross-repo-spawn-20261015-065000-2855",
				"e2e-cross-repo-spawn-20261015-072002-0653",
				"e2e-cross-repo-spawn-20261017-141052-4438",
				"e2e-cross-repo-spawn-20261017-165734-4675",
				"e2e-cross-repo-spawn-20261019-055333-0025",
				"e2e-cross-repo-spawn-20261019-062236-1759",
				"e2e-cross-repo-spawn-20261019-070912-5993",
				"e2e-cross-repo-spawn-20261019-122156-0514",
				"e2e-cross-repo-spawn-20261019-130515-9466",
				"e2e-cross-repo-spawn-20261019-134215-7557",
				"e2e-cross-repo-spawn-20261019-154343-1133",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdArgs := append([]string{"run", "."}, tt.args...)
			out, err := exec.Command("go", cmdArgs...).CombinedOutput()
			if err != nil {
				t.Fatalf("go run failed: %v\nOutput: %s", err, out)
			}
			got := string(out)
			for _, s := range tt.expected {
				if !strings.Contains(got, s+"\n") {
					t.Errorf("expected output to contain line %q, got %q", s, got)
				}
			}
		})
	}
}

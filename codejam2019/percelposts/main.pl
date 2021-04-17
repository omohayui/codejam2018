use strict;
use warnings;
use utf8;

my $t = <STDIN>; # test case 数

for (my $i = 1; $i <= $t; $i++) {
    my $km = <STDIN>;             # 区画のKm
    my @ms = split(/ /, <STDIN>); # 各pointの標高
    my $mark = 0;                 # 追加できるmarkの数
    my $tmpM = 0;
    for (my $j = 0; $j <= $km; $j++) {
        if ($j <= 1 || $j >= $km-1) {
            next;
        }

        my $before = 0; # false
        my $after = 0; # false

        my $high = 0; #false
        my $low = 0; #false
        for (my $k = $tmpM; $k < $j; $k++) {
            if ($ms[$k] > $ms[$k+1]) {
                if ($low) {
                    $before = 1;
                }
                $high = 1;
            }
            if ($ms[$k] < $ms[$k+1]) {
                if ($high) {
                    $before = 1;
                }
                $low = 1;
            }
        }

        my $aHigh = 0; #false
        my $aLow = 0; #false
        for (my $k = $j; $k < $km; $k++) {
            if ($ms[$k] > $ms[$k+1]) {
                if ($aLow) {
                    $after = 1;
                }
                $aHigh = 1;
            }
            if ($ms[$k] < $ms[$k+1]) {
                if ($aHigh) {
                    $after = 1;
                }
                $aLow = 1;
            }
        }

        if ($before && $after) {
            $mark++;
            $tmpM = $j;
        }
    }

    printf("Case #%d: %d\n", $i, $mark);
}


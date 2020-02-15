use strict;
use warnings;
use utf8;

my $t = <STDIN>; # test case 数

for (my $i = 1; $i <= $t; $i++) {
    my $k = <STDIN>; # burger の中の要素数
    my @bs = split(/ /, <STDIN>); # 最適なバンズまでの距離の配列
    my @ds;                        # バンズまでの距離の配列
    for (my $j = 0; $j < $k; $j++) {
        push @ds, distance($k, $j); # バンズまでの距離
    }

    # 最小誤差を出すために並び替える
    @ds = sort { $a <=> $b } @ds;
    @bs = sort { $a <=> $b } @bs;

    my $es = 0; # 誤差合計
    for (my $j = 0; $j < $k; $j++) {
        my $d = ($ds[$j] - $bs[$j]) ** 2; # 誤差
        $es += $d;
    }
    printf("Case #%d: %d\n", $i, $es);
}

# バンズからの距離
sub distance {
    my ($k, $j) = @_;
    my $d = $k - 1;
    return $d - $j if ($d - $j) < $j;
    return $j;
}


<?php

$arr = [1, 4, 5];
$target = 6;

for ($i=0; $i < count($arr); $i++) { 
    for ($j=$i; $j < count($arr) ; $j++) { 
        if ($arr[$i] + $arr[$j] == $target) {
            echo "[".$i." ".$j."]";
        }
    }
}

?>
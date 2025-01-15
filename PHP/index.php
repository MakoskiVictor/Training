<?php

class main {
    static public function loop() {
        // Iniciar el contador de tiempo
        $start_time = microtime(true);

        for ($i = 0; $i < 100000; $i++) {
            $calculate = $i * 3;
            echo $calculate . "\n";
        }

        // Calcular el tiempo transcurrido
        $end_time = microtime(true);
        $execution_time = $end_time - $start_time;
        echo $execution_time;
    }
}

main::loop();
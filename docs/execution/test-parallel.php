<?php

$payload = file_get_contents("forxy-payload.json");
$json_obj = json_decode($payload, true);

$multi_handler = curl_multi_init();
$curl_handlers = [];

foreach ($json_obj['requests'] as $id => $req) {
    $ch = curl_init($req['url']);
    
    $json_body = json_encode($req['body']);
    
    curl_setopt($ch, CURLOPT_HTTPGET, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $json_body);
    curl_setopt($ch, CURLOPT_HTTPHEADER, [
        'Content-Type: application/json'
    ]);

    curl_multi_add_handle($multi_handler, $ch);
    $curl_handlers[$id] = $ch;
}

$running = null;
$start = microtime(true);
do {
    curl_multi_exec($multi_handler, $running);
} while ($running > 0);

foreach ($curl_handlers as $id => $ch) {
    $response = curl_multi_getcontent($ch);
    $http_status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    echo ("Request $id: Status code: $http_status \n");
    
    curl_multi_remove_handle($multi_handler, $ch);
    curl_close($ch);
}

$end = microtime(true);
curl_multi_close($multi_handler);

echo ("Execution time: " . ($end-$start));
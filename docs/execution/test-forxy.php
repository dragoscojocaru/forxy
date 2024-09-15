<?php

$payload = file_get_contents("forxy-payload.json");

$url = "http://localhost:8080/http/fork";

$ch = curl_init($url);

curl_setopt($ch, CURLOPT_HTTPGET, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $payload); 
curl_setopt($ch, CURLOPT_HTTPHEADER, array(
    'Content-Type: application/json'
));

$start = microtime(true);

$response = curl_exec($ch);
$http_status = curl_getinfo($ch, CURLINFO_HTTP_CODE);

echo "Status code: $http_status\n Response body: $response\n";

$end = microtime(true);
curl_close($ch);

echo ("Execution time: " . ($end-$start));
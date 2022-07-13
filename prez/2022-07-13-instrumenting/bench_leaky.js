
import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  http.get('http://localhost:8081/leaky');
  sleep(1);
}

//  k6 run --vus 10 --duration 5s bench_leaky.js
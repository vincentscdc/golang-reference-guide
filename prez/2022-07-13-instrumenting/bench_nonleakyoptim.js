
import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  http.get('http://localhost:8081/nonleakyoptim');
  sleep(1);
}

// k6 run --vus 100 --duration 2s bench_nonleakyoptim.js
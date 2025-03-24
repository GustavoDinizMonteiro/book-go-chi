import http from 'k6/http'
import { check, sleep } from 'k6'

const URL = 'http://localhost:5000/books'

export default function () {
  let res = http.get(URL)

  check(res, {
    'status é 200': (r) => r.status === 200,
  });

  sleep(1)
}

import http from 'k6/http';
import { check, sleep } from "k6";

export default function () {
  const response = http.get("http://172.31.86.233:4000/api/get-gas-price", {headers: {Accepts: "application/json"}});
  check(response, { "status is 200": (r) => r.status === 200 });
  sleep(.300);
};

import http from "k6/http";
import { check } from "k6";

export default function () {
  const response = http.get("http://proxy:8000/hi", {
    headers: { Authorization: "123456" },
  });
  check(response, {
    "it succeeded because it was from the proxy and included the auth header":
      response.status === 200,
  });

  const response1 = http.get("http://proxy:8000/hi");
  check(response1, {
    "we got 403 since we didn't include the auth header":
      response1.status === 403,
  });
  const response2 = http.get("http://target:8000/hi");
  check(response2, {
    "the target gives 403 if we interact directly": response2.status === 403,
  });
}

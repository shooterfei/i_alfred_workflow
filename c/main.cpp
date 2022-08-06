#include <boost/beast/http.hpp>
#include <cstring>
#include <iostream>

namespace http = boost::beast::http;

using namespace std;

int main() {
  auto verb = http::verb::get;
  string query = "";
  string data = "";
  string host = "www.baidu.com";

  http::request<http::string_body> req(verb, query + data, 11);
  req.set(http::field::host, host);
  // req.set("key", "I am a header");
  req.prepare_payload();

  cout << req << endl;
  return 0;
}

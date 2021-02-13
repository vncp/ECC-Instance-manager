#include <cpprest/http_client.h>
#include <cpprest/json.h>
#pragma comment(lib, "cpprest_2_10")

using namespace web;
using namespace web::http;
using namespace web::http::client;

#include <iostream>
using namespace std;

void display_json(const json::value &jvalue, const utility::string_t &prefix) {
    cout << prefix << jvalue.serialize() << endl;
}

pplx::task<http_response> make_task_request(http_client &client, method mtd, const json::value &jvalue) {
    return (mtd == methods::GET || mtd == methods::HEAD) ?
        client.request(mtd, "/api") : 
        client.request(mtd, "/api", jvalue);
}

void make_request(http_client &client, method mtd, const json::value &jvalue) {
    make_task_request(client, mtd, jvalue)
        .then([](http_response response) {
          if(response.status_code() == status_codes::OK) {
              return response.extract_json();
          }  
          return pplx::task_from_result(json::value());
        })
        .then([](pplx::task<json::value> previousTask) {
            try {
                display_json(previousTask.get(), "R: ");
            } catch (const http_exception &e) {
                cout << e.what() << endl;
            }
        })
        .wait();
}

int main() {
    http_client client(U("http://localhost:3001"));
    auto putvalue = json::value::object();
    putvalue["one"] = json::value::string("100");
    putvalue["two"] = json::value::string("200");
    cout << "\nPut (add values)";
    display_json(putvalue, "S: ");
    make_request(client, methods::PUT, putvalue);

    auto getvalue = json::value::array();
    getvalue[0] = json::value::string("one");
    getvalue[1] = json::value::string("two");
    getvalue[2] = json::value::string("three");

    cout << "Post (get Some values)\n";
    display_json(getvalue, "S: ");
    make_request(client, methods::POST, getvalue);

    return 0;
}
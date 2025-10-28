import csv
import os
import json
from datetime import datetime
from http.server import BaseHTTPRequestHandler, HTTPServer

LOG_FILE = 'heartbeat_log.csv'
LOG_FIELDNAMES = ['server_timestamp', 'visit_id', 'url', 'referrer', 'duration_on_page', 'client_timestamp']

def initialize_csv():
    file_exists = os.path.isfile(LOG_FILE)
    if not file_exists:
        with open(LOG_FILE, mode='w', newline='', encoding='utf-8') as f:
            writer = csv.writer(f)
            writer.writerow(LOG_FIELDNAMES)
        print(f"Created log file: {LOG_FILE}")

class HeartbeatHandler(BaseHTTPRequestHandler):
    def _send_cors_headers(self):
        self.send_header('Access-Control-Allow-Origin', '*')
        self.send_header('Access-Control-Allow-Methods', 'POST, OPTIONS')
        self.send_header('Access-Control-Allow-Headers', 'Content-Type')

    def do_OPTIONS(self):
        #Responds to pre-flight CORS OPTIONS requests. (whatever that means)
        self.send_response(204)  # 204 No Content
        self._send_cors_headers()
        self.end_headers()

    def do_POST(self):
        if self.path == '/heartbeat':
            try:
                #Get request body length
                content_length = int(self.headers['Content-Length'])
                #Read and parse the JSON body
                body = self.rfile.read(content_length)
                data = json.loads(body.decode('utf-8'))
                if not data:
                    #self._send_response(400, {"status": "error", "message": "No data provided"})
                    print("Error getting request or parsing request data to JSON.")
                    return

                #Extract and assign data
                visit_id = data.get('visitId')
                url = data.get('url')
                referrer = data.get('referrer')
                duration = data.get('duration')
                client_timestamp = data.get('timestamp')
                if not all([visit_id, url, duration is not None]):
                    #self._send_response(400, {"status": "error", "message": "Missing required data"})
                    print(f"Some data missing: {data}")
                    return

                #Log the data (same logic as before)
                server_timestamp = datetime.now().isoformat()
                print(f"Heartbeat received: {visit_id} | URL: {url} | Duration: {duration:.2f}s")

                row = [
                    server_timestamp, 
                    visit_id, 
                    url,
                    referrer,
                    f"{duration:.2f}",
                    client_timestamp
                ]
                with open(LOG_FILE, mode='a', newline='', encoding='utf-8') as f:
                    writer = csv.writer(f)
                    writer.writerow(row)

                #Send success
                self._send_response(200, {"status": "success"})

            except json.JSONDecodeError:
                self._send_response(400, {"status": "error", "message": "Invalid JSON"})
            except Exception as e:
                print(f"Error processing request: {e}")
                self._send_response(500, {"status": "error", "message": str(e)})
        else:
            self._send_response(404, {"status": "error", "message": "Not Found"})

    def do_GET(self):
        if self.path.startswith('/log') or self.path.startswith('/stat') or self.path.startswith('/data'): #haven't decided yet
            #just return the text from the entire log file for now
            #as an page of text (this will be a regular GET request, not POST)
            with open("heartbeat_log.csv", "r") as log:
                log_content = log.read()
                self.send_response(200)
                self.send_header("Content-type", "text/plain")
                self.end_headers()
                self.wfile.write(log_content.encode('utf-8'))

    def _send_response(self, status_code, message_dict):
        #Sends a JSON response
        self.send_response(status_code)
        self._send_cors_headers() # Include CORS headers in all responses
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        self.wfile.write(json.dumps(message_dict).encode('utf-8'))

if __name__ == '__main__':
    initialize_csv() # Create the CSV with headers if it doesn't exist
    server_address = ("", 8080)
    httpd = HTTPServer(server_address, HeartbeatHandler)
    print(f"Starting http.server on http://localhost:8080...")
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        print("\nStopping server...")
        httpd.server_close()
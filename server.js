var buffer = new Buffer(1024*1024);
for (var i = 0; i < buffer.length; i++)
  buffer[i] = 100;
 
var http = require('http');
http.createServer(function (req, res) {
  res.writeHead(200, {'Content-Type': 'text/plain'});

	for (var i = 0; i < 10; i++) {
		a = {
			"foo": "bar",
			"bar": "baz",
			"nest": {
				"a": "b",
			},
		}
		b = JSON.stringify(a)
		res.write(b)
	}
	res.end()
  //res.end(buffer);
  //res.end("Hello world!");
}).listen(8000);
 
console.log('Server running at http://127.0.0.1:8000/');

var express = require('express');
var http = require('http');
var request = require('request');
var promise = require('promise');
var router = express.Router();

//check http://localhost:1323/api/v1/api/json-test/single-test/10000
/* GET home page. */
router.get('/', function (req, res, next) {
    var curUrl = req.originalUrl.split('/');
    res.render('index', {title: 'Express'});
});

var testCnt;
var start;
var nCnt;

router.get('/api/json-test/single-test/:count', function (req, res, next) {
    var curUrl = req.originalUrl.split('/');
    var cnt = req.params.count;
    // console.log('들어왔다');
    console.log(curUrl);

    start = Date.now();
    testCnt = 0;
    nCnt = cnt;
    var outRes = res;
    for (var i = 0; i < cnt; i++) {
        request('http://localhost:3000/api/json-test/single-json/', null, function (err, res, body) {
            testCnt++;
            // console.log(testCnt);
            if (testCnt == nCnt) {
                var end = Date.now();
                console.log("뭐하냐");
                console.log({
                    avgResponseTime: (end - start) / nCnt + "ms",
                    count: cnt,
                    sumResponseTime: end - start + "ms"
                });
                //template {"avgResponseTime":"6196us","count":10000,"sumResponseTime":"61967306us"}
                outRes.send({
                    avgResponseTime: (end - start) / nCnt + "ms",
                    count: cnt,
                    sumResponseTime: end - start + "ms"
                });
            }
        })
    }
});


module.exports = router;
